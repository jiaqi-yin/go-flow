package asyncservice

import (
	"errors"
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
func ServiceA() (*ResponseA, error) {
	time.Sleep(500 * time.Millisecond) // Simulate network delay
	if rand.Float32() < 0.3 {          // 30% chance of error
		return nil, errors.New("ServiceA error")
	}
	return &ResponseA{Data: "Data from ServiceA"}, nil
}

func ServiceB() ([]*ResponseB, error) {
	time.Sleep(500 * time.Millisecond) // Simulate network delay
	if rand.Float32() < 0.3 {          // 30% chance of error
		return nil, errors.New("ServiceB error")
	}
	return []*ResponseB{
		{Data: rand.Intn(100)},
		{Data: rand.Intn(100)},
	}, nil
}

func ServiceC() ([]*ResponseC, error) {
	time.Sleep(500 * time.Millisecond) // Simulate network delay
	if rand.Float32() < 0.3 {          // 30% chance of error
		return nil, errors.New("ServiceB error")
	}
	return []*ResponseC{
		{Data: rand.Float32()},
		{Data: rand.Float32()},
	}, nil
}
