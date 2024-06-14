package asyncservice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	asyncservice "go-flow/async_service"
)

func TestAsyncService(t *testing.T) {
	var errors = []error{}

	responses, errs := asyncservice.NewAsyncServiceManager(
		asyncservice.NewServiceCall("ServiceA", func(...any) (any, error) { return asyncservice.ServiceA("foo bar") }),
		asyncservice.NewServiceCall("ServiceB", func(...any) (any, error) { return asyncservice.ServiceB() }),
		asyncservice.NewServiceCall("ServiceC", func(...any) (any, error) { return asyncservice.ServiceC() }),
	).Async()

	// Handle errors and process responses
	for _, err := range errs {
		if err != nil {
			errors = append(errors, err)
		}
	}

	assert.Equal(t, 3, len(responses)+len(errors))
}
