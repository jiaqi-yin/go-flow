package asyncservice

import (
	"sync"
)

type ServiceCall struct {
	Call       func() (interface{}, error)
	ResponseCh chan interface{}
	ErrorCh    chan error
}

type AsyncServiceManager struct {
	Services []*ServiceCall
}

func NewAsyncServiceManager(services []*ServiceCall) *AsyncServiceManager {
	return &AsyncServiceManager{
		Services: services,
	}
}

func (s *AsyncServiceManager) Async() ([]interface{}, []error) {
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
	responses := make([]interface{}, numServices)
	errors := make([]error, numServices)

	for i, service := range s.Services {
		select {
		case res := <-service.ResponseCh:
			responses[i] = res
		case err := <-service.ErrorCh:
			errors[i] = err
		}
	}

	return responses, errors
}

// Generic function to handle concurrent service calls
func FetchService[T any](wg *sync.WaitGroup, service func() (T, error), responseCh chan<- T, errorCh chan<- error) {
	defer wg.Done()
	if res, err := service(); err != nil {
		errorCh <- err
	} else {
		responseCh <- res
	}
}
