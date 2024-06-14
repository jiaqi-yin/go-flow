package asyncservice

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Define the response structs
type ResponseA struct {
	Data string
}

type ResponseB struct {
	Data int
}

type ResponseC struct {
	Data float32
}

// Simulate the service calls that may return an error
func ServiceA(foo string) (*ResponseA, error) {
	time.Sleep(150 * time.Millisecond) // Simulate network delay
	if rand.Float32() < 0.3 {          // 30% chance of error
		return nil, errors.New("ServiceA error")
	}
	return &ResponseA{Data: fmt.Sprintf("Data from ServiceA: %s", foo)}, nil
}

func ServiceB() ([]*ResponseB, error) {
	time.Sleep(100 * time.Millisecond) // Simulate network delay
	if rand.Float32() < 0.3 {          // 30% chance of error
		return nil, errors.New("ServiceB error")
	}
	return []*ResponseB{
		{Data: rand.Intn(100)},
		{Data: rand.Intn(100)},
	}, nil
}

func ServiceC() ([]*ResponseC, error) {
	time.Sleep(50 * time.Millisecond) // Simulate network delay
	if rand.Float32() < 0.3 {         // 30% chance of error
		return nil, errors.New("ServiceC error")
	}
	return []*ResponseC{
		{Data: rand.Float32()},
		{Data: rand.Float32()},
	}, nil
}
