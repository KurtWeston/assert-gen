package formatter

import (
	"fmt"
	"reflect"
	"strings"
)

func FormatValue(v interface{}) string {
	if v == nil {
		return "nil"
	}
	val := reflect.ValueOf(v)
	return formatReflectValue(val)
}

func formatReflectValue(val reflect.Value) string {
	if !val.IsValid() {
		return "nil"
	}

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", val.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", val.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", val.Bool())
	case reflect.Slice, reflect.Array:
		return formatSlice(val)
	case reflect.Map:
		return formatMap(val)
	case reflect.Struct:
		return formatStruct(val)
	case reflect.Ptr:
		if val.IsNil() {
			return "nil"
		}
		return "&" + formatReflectValue(val.Elem())
	case reflect.Interface:
		if val.IsNil() {
			return "nil"
		}
		return formatReflectValue(val.Elem())
	default:
		return fmt.Sprintf("%#v", val.Interface())
	}
}

func formatSlice(val reflect.Value) string {
	if val.Len() == 0 {
		return fmt.Sprintf("%s{}", val.Type())
	}
	parts := make([]string, val.Len())
	for i := 0; i < val.Len(); i++ {
		parts[i] = formatReflectValue(val.Index(i))
	}
	return fmt.Sprintf("%s{%s}", val.Type(), strings.Join(parts, ", "))
}

func formatMap(val reflect.Value) string {
	if val.Len() == 0 {
		return fmt.Sprintf("%s{}", val.Type())
	}
	parts := []string{}
	for _, key := range val.MapKeys() {
		mapVal := val.MapIndex(key)
		parts = append(parts, fmt.Sprintf("%s: %s", formatReflectValue(key), formatReflectValue(mapVal)))
	}
	return fmt.Sprintf("%s{%s}", val.Type(), strings.Join(parts, ", "))
}

func formatStruct(val reflect.Value) string {
	typ := val.Type()
	parts := []string{}
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldVal := val.Field(i)
		parts = append(parts, fmt.Sprintf("%s: %s", field.Name, formatReflectValue(fieldVal)))
	}
	return fmt.Sprintf("%s{%s}", typ.Name(), strings.Join(parts, ", "))
}
