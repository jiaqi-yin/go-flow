package pipeline_test

import (
	"fmt"
	"testing"

	"go-flow/pipeline"
)

func TestPipeline(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	source := []*User{}
	for i := 0; i < 10; i++ {
		user := &User{
			Name: "John",
			Age:  i,
		}
		source = append(source, user)
	}
	userStream := pipeline.Generate(source)
	employeeStream := pipeline.Map(userStream, func(user *User) *Employee {
		return &Employee{
			User:     *user,
			Position: "Developer",
		}
	})
	employeeStream = pipeline.Filter(employeeStream, func(employee *Employee) bool { return employee.Age > 5 })

	// employeeList := pipeline.Collect(done, employeeStream)
	// for _, v := range employeeList {
	// 	fmt.Println(v.Name, v.Age, v.Position)
	// }

	for v := range employeeStream {
		fmt.Println(v.Name, v.Age, v.Position)
	}
}

type User struct {
	Name string
	Age  int
}

type Employee struct {
	User
	Position string
}
