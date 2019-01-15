package main

import "reflect"

func (p *pp) printArg(arg interface{}, verb rune) {
	...
	// Some types can be done without reflection. // HL
	switch f := arg.(type) {
	case bool:
		p.fmtBool(f, verb)
	case float32:
		p.fmtFloat(float64(f), 32, verb)
	case int:
		p.fmtInteger(uint64(f), signed, verb)
	case int8:
		p.fmtInteger(uint64(f), signed, verb)
	...
	default:
		// If the type is not simple, it might have methods.
		if !p.handleMethods(verb) {
			// Need to use reflection, since the type had no
			// interface methods that could be used for formatting.
			p.printValue(reflect.ValueOf(f), verb, 0) // HL
		}
	}
}

func (p *pp) printValue(value reflect.Value, verb rune, depth int) {
	...
	switch f := value; value.Kind() {
	// reflect.Int, reflect.Int8 etc...
	case reflect.Struct:
		p.buf.WriteByte('{')
		for i := 0; i < f.NumField(); i++ { // HL
			...
			if p.fmt.plusV || p.fmt.sharpV {
				if name := f.Type().Field(i).Name; name != "" {
					p.buf.WriteString(name)
					p.buf.WriteByte(':')
				}
			}
			p.printValue(getField(f, i), verb, depth+1) // HL
		}
		p.buf.WriteByte('}')
	}
	// reflect.Slice, reflect.Map, reflect.Interface etc...
}