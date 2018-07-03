package common

import (
	"reflect"
	"time"
)

// structをmap変換
// ※structはポインタで渡すこと
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()

		if value != "" && value != nil {
			result[field] = value
		}
	}

	return result
}

// 現在の日付を取得
func GetTime() time.Time {
	return time.Now()
}
