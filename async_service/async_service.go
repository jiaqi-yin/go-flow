package asyncservice

import (
	"sync"
)

type serviceCall struct {
	Name       string
	Call       func(...any) (any, error)
	ResponseCh chan any
	ErrorCh    chan error
}

func NewServiceCall(name string, call func(...any) (any, error)) *serviceCall {
	return &serviceCall{
		Name:       name,
		Call:       call,
		ResponseCh: make(chan any),
		ErrorCh:    make(chan error),
	}
}

type asyncServiceManager struct {
	Services []*serviceCall
}

func NewAsyncServiceManager(services ...*serviceCall) *asyncServiceManager {
	return &asyncServiceManager{
		Services: services,
	}
}

func (s *asyncServiceManager) Async() (map[string]any, []error) {
	numServices := len(s.Services)

	var wg sync.WaitGroup
	// Add the number of services to the WaitGroup
	wg.Add(numServices)

	// Call services concurrently
	for _, svc := range s.Services {
		go FetchService(&wg, svc.Call, svc.ResponseCh, svc.ErrorCh)
	}

	// Close channels once all goroutines are done
	go func() {
		wg.Wait()
		for _, svc := range s.Services {
			close(svc.ResponseCh)
			close(svc.ErrorCh)
		}
	}()

	// Collect responses and handle errors
	responses := make(map[string]any, numServices)
	errors := []error{}

	for _, service := range s.Services {
		select {
		case res := <-service.ResponseCh:
			responses[service.Name] = res
		case err := <-service.ErrorCh:
			errors = append(errors, err)
		}
	}

	return responses, errors
}

// Generic function to handle concurrent service calls
func FetchService[T any](wg *sync.WaitGroup, service func(...any) (T, error), responseCh chan<- T, errorCh chan<- error) {
	defer wg.Done()
	if res, err := service(); err != nil {
		errorCh <- err
	} else {
		responseCh <- res
	}
}
