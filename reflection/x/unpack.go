package x

import (
	"fmt"
	"net/http"
	"reflect"
)

// START OMIT
func unpack(req *http.Request, dest interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(dest))
	m := DefaultMapper.FieldMap(v) // HL
	for name, values := range req.Form {
		f, ok := m[name]
		if !ok {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else if err := populate(f, value); err != nil {
				return fmt.Errorf("%s: %v", name, err)
			}
		}
	}
	return nil
}

// END OMIT
