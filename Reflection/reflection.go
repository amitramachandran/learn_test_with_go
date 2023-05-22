package reflection

import (
	"reflect"
)

func Walk(x interface{}, f func(input string)) string {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	return field.String()
}
