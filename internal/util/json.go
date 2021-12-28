package util

import "reflect"

func StructToMapByJsonTag(s interface{}) map[string]interface{} {
	var res = map[string]interface{}{}

	if s == nil {
		return res
	}
	v := reflect.TypeOf(s)
	reflectValue := reflect.ValueOf(s)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			res[tag] = field
		}
	}
	return res
}
