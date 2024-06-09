package flow_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-flow/flow"
)

func TestNewFlow(t *testing.T) {
	flow := flow.NewFlow[int]()
	require.NotNil(t, flow)
}

func TestFlow(t *testing.T) {
	var result = []interface{}{}

	// Create a new Flow
	flow := flow.NewFlow[interface{}]()

	// Emit values in a separate goroutine
	go func() {
		flow.Emit("Hello, World!")
		for i := 1; i <= 5; i++ {
			flow.Emit(i)
		}
		flow.Emit(Sample{Name: "John"})
		flow.Close()
	}()

	// Collect the values
	flow.Collect(func(val interface{}) {
		result = append(result, val)
	})

	assert.Equal(t, "Hello, World!", result[0])
	assert.Equal(t, 1, result[1])
	assert.Equal(t, 2, result[2])
	assert.Equal(t, 3, result[3])
	assert.Equal(t, 4, result[4])
	assert.Equal(t, 5, result[5])
	assert.Equal(t, Sample{Name: "John"}, result[6])
}

type Sample struct {
	Name string
}
