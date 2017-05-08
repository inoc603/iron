package iron

import (
	"reflect"
	"strconv"
	"strings"
)

func join(prefix, name string) string {
	if prefix == "" {
		return name
	}

	if name == "" {
		return prefix
	}

	return prefix + "." + name
}

func getName(f reflect.StructField) string {
	name := strings.Split(f.Tag.Get("json"), ",")[0]
	if name == "" {
		return f.Name
	}
	return name
}

// Flatten convert a nested struct or map into a single layer map
func Flatten(obj interface{}) map[string]interface{} {
	return flatten(reflect.ValueOf(obj))
}

func flatten(ov reflect.Value) map[string]interface{} {
	res := make(map[string]interface{})

	for ov.Kind() == reflect.Ptr || ov.Kind() == reflect.Interface {
		if ov.IsNil() {
			res[""] = nil
			return res
		}
		ov = ov.Elem()
	}

	ot := ov.Type()

	switch ot.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		for i := 0; i < ov.Len(); i++ {
			for k, v := range flatten(ov.Index(i)) {
				res[join(strconv.Itoa(i), k)] = v
			}
		}
	case reflect.Struct:
		for i := 0; i < ot.NumField(); i++ {
			for k, v := range flatten(ov.Field(i)) {
				res[join(getName(ot.Field(i)), k)] = v
			}
		}
	case reflect.Map:
		for _, mapKey := range ov.MapKeys() {
			for k, v := range flatten(ov.MapIndex(mapKey)) {
				res[join(mapKey.String(), k)] = v
			}
		}
	default:
		res[""] = ov.Interface()
	}

	return res
}
