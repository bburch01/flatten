package flatten

import (
	"fmt"
	"reflect"
)

// Flatten flattens the supplied multidimensional array into a flat (one dimension) array or
// returns an error. Flatten currenlty only supports integer as the array element type.
func Flatten(s []interface{}) (r []int, err error) {
	for _, e := range s {
		switch v := e.(type) {
		case int:
			r = append(r, v)
		case []interface{}:
			var ts []int
			if ts, err = Flatten(v); err == nil {
				r = append(r, ts...)
			}
		default:
			err = fmt.Errorf("slice element (%v) type (%v) not supported", v, reflect.TypeOf(e).Kind())
		}
	}
	return
}

// Compose is a helper function that simplifies creating multidimensional arrays to be used for input
// to the Flatten function. Here is an example of how to use Abastract to build input for
// Flatten that corresponds to the multidimensional integer array [[1 2] [3, 4, 5 [6, 7]]] :
// s = Compose(1, 2, Compose(3, 4, 5, Compose(6, 7)))
//
func Compose(s ...interface{}) []interface{} {
	return s
}
