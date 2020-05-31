package main

import (
	"fmt"
	"reflect"
)

// T struct
type T struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	t := T{Name: "ray", Age: 18}

	m, err := ToMap(&t, "json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

// ToMap struct to map
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got: %s", v.Kind())
	}

	t := v.Type()

	// 遍历结构体字段
	// 指定tagName值为map中的key，字段值为map中的value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)

		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}

	return out, nil
}
