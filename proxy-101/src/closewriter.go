func asCloseWriter(w io.Writer) (closeWriter, bool) {
	if cw, ok := w.(closeWriter); ok {
		return cw, true
	}

	v := reflect.Indirect(reflect.ValueOf(w))

	if v.Kind() != reflect.Struct {
		return nil, false
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.CanInterface() {
			if cw, ok := f.Interface().(closeWriter); ok {
				return cw, true
			}
		}
	}

	return nil, false
}
