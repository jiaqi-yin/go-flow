package main

import (
	"fmt"
	asyncservice "go-flow/async_service"
	"time"
)

func main() {
	startTIme := time.Now()
	// Define the services to be called
	services := []*asyncservice.ServiceCall{
		{
			Call:       func() (interface{}, error) { return asyncservice.ServiceA() },
			ResponseCh: make(chan interface{}),
			ErrorCh:    make(chan error),
		},
		{
			Call:       func() (interface{}, error) { return asyncservice.ServiceB() },
			ResponseCh: make(chan interface{}),
			ErrorCh:    make(chan error),
		},
		{
			Call:       func() (interface{}, error) { return asyncservice.ServiceC() },
			ResponseCh: make(chan interface{}),
			ErrorCh:    make(chan error),
		},
	}
	responses, errors := asyncservice.NewAsyncServiceManager(services).Async()

	// Handle errors and process responses
	for i, err := range errors {
		if err != nil {
			fmt.Printf("Error calling Service%d: %v\n", i+1, err)
		} else {
			fmt.Printf("Received Response from Service%d: %v\n", i+1, responses[i])
		}
	}

	// Further processing of the responses if no errors
	allSuccessful := true
	for _, err := range errors {
		if err != nil {
			allSuccessful = false
			break
		}
	}

	if allSuccessful {
		processResponses(responses)
	}

	duration := time.Since(startTIme)
	fmt.Println("Total time taken 1:", duration)

	startTime2 := time.Now()
	asyncservice.ServiceA()
	asyncservice.ServiceB()
	asyncservice.ServiceC()
	duration2 := time.Since(startTime2)
	fmt.Println("Total time taken 2:", duration2)
}

// Example function to process the responses
func processResponses(responses []interface{}) {
	fmt.Println("Processing responses:")
	for _, res := range responses {
		switch r := res.(type) {
		case *asyncservice.ResponseA:
			fmt.Printf("ResponseA Data: %s\n", r.Data)
		case []*asyncservice.ResponseB:
			for _, b := range r {
				fmt.Printf("ResponseB Data: %d\n", b.Data)
			}
		case []*asyncservice.ResponseC:
			for _, c := range r {
				fmt.Printf("ResponseC Data: %f\n", c.Data)
			}
		default:
			fmt.Printf("Unknown response type: %v\n", r)
		}
	}
}
