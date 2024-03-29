package flatten

import (
	"fmt"
	"reflect"
	"testing"
)

type validInputCase struct {
	in  []interface{}
	out []int
}

type invalidInputCase struct {
	in []interface{}
}

type errTypeB struct {
	code    string
	message string
}

var validInputCases = []validInputCase{
	{in: Compose(Compose(), 1, 2, 3), out: []int{1, 2, 3}},
	{in: Compose(1, 2, 3, Compose()), out: []int{1, 2, 3}},
	{in: Compose(1, 2, Compose(), 3), out: []int{1, 2, 3}},
	{in: Compose(Compose(1, 2, 3), 4, 5, 6), out: []int{1, 2, 3, 4, 5, 6}},
	{in: Compose(Compose(1, 2, 3), 4, 5, 6), out: []int{1, 2, 3, 4, 5, 6}},
	{in: Compose(Compose(1, Compose(2, Compose(3, 4)), 5), 6, Compose(Compose(), 7), 8, 9),
		out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

var invalidInputCases = []invalidInputCase{
	{in: Compose(1, 2, "three")},
	{in: Compose(Compose(), 1.0, 2, 3)},
	{in: Compose(Compose(), 1, errTypeB{code: "error", message: "kaboom!"}, 3)},
	{in: Compose(1, 2, true)},
}

func TestValidInputCases(t *testing.T) {

	var out []int
	var err error

	for _, tc := range validInputCases {
		if out, err = Flatten(tc.in); err != nil {
			t.Error("valid input test case failed with error: ", err)
			t.FailNow()
		}
		if !reflect.DeepEqual(out, tc.out) {
			t.Error(fmt.Sprintf("valid input test case failed, expected %v got %v: ", tc.out, out))
			t.FailNow()
		}
	}
}

func TestInvalidInputCases(t *testing.T) {
	for _, tc := range invalidInputCases {
		if _, err := Flatten(tc.in); err == nil {
			t.Error("failed to detect invalid (non-integer) type in the input data")
			t.FailNow()
		}
	}
}
