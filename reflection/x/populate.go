package x

import (
	"fmt"
	"reflect"
	"strconv"
)

// START OMIT
func setValue(v reflect.Value, value string) error {
	switch v.Kind() {
	// Special handling of slices.
	case reflect.Slice:
		elem := reflect.New(v.Type().Elem()).Elem()
		if err := setValue(elem, value); err != nil {
			return err
		}
		v.Set(reflect.Append(v, elem))
	// Handling of selected types...
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

// END OMIT
