package x

import (
	"fmt"
	"net/http"
	"reflect"
)

// START OMIT
func bindParams(req *http.Request, dest interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(dest))
	m := DefaultMapper.FieldMap(v) // HL
	for name, values := range req.Form {
		f, ok := m[name]
		if !ok {
			continue
		}
		for _, value := range values {
			if err := setValue(f, value); err != nil { // HL
				return fmt.Errorf("%s: %v", name, err)
			}
		}
	}
	return nil
}

// END OMIT
