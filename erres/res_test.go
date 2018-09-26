package erres_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/gemcook/restres-go/erres"
	"github.com/stretchr/testify/assert"
)

type errorObject struct {
	value int
}

func (e errorObject) String() string {
	return fmt.Sprintf("error value is %v", e.value)
}

func TestErrorResponse(t *testing.T) {

	res := erres.New("Unexpected", erres.ErrorUnknown, "something wrong", errors.New("segmentation fault"), errorObject{123})

	assert := assert.New(t)

	if len(res.Errors) == 0 {
		t.Fatalf("res.Errors() has wrong number of errors. got=%d, want=%d", len(res.Errors), 0)
	}
	assert.Equal(erres.ErrorUnknown, res.Errors[0].Type)
	assert.Equal("Unexpected", res.Errors[0].Code)

	messages := res.Errors[0].Messages
	if len(messages) != 3 {
		t.Fatalf("res.Errors[0].Messages has wrong number of messages. got=%d, want=%d", len(messages), 3)
	}
	assert.Equal("something wrong", messages[0])
	assert.Equal("segmentation fault", messages[1])
	assert.Equal("error value is 123", messages[2])

	if len(res.String()) == 0 {
		t.Fatalf("res.String() should return something.")
	}

	res.Append("NetworkChanged", erres.ErrorUnknown, "network change detected", 999)
	if len(res.Errors) != 2 {
		t.Fatalf("res.Errors length should be 2. got=%d", len(res.Errors))
	}

	assert.Equal(erres.ErrorUnknown, res.Errors[1].Type)
	assert.Equal("NetworkChanged", res.Errors[1].Code)

	messages = res.Errors[1].Messages
	if len(messages) != 2 {
		t.Fatalf("res.Errors[1].Messages has wrong number of messages. got=%d, want=%d", len(messages), 2)
	}

	assert.Equal("network change detected", messages[0])
	assert.Equal("999", messages[1])
}

func ExampleNew_with3Messages() {
	res := erres.New("Unexpected", erres.ErrorUnknown,
		"something wrong",
		errors.New("segmentation fault"),
		errorObject{123},
	)

	b, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(b))
	// Output:
	// {
	//   "errors": [
	//     {
	//       "code": "Unexpected",
	//       "type": "UnknownError",
	//       "messages": [
	//         "something wrong",
	//         "segmentation fault",
	//         "error value is 123"
	//       ]
	//     }
	//   ]
	// }
}
