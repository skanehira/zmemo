package common

import (
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

// StructToMap structをmap変換
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

// GetTime 現在の日付を取得
func GetTime() time.Time {
	return time.Now()
}

// NewUUID new uuid
func NewUUID() string {
	uuid := uuid.Must(uuid.NewV4())
	return uuid.String()
}
