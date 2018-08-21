package common

import (
	"reflect"
	"regexp"

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

// NewUUID new uuid
func NewUUID() string {
	uuid := uuid.Must(uuid.NewV4())
	return uuid.String()
}

var IsAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`) // 英文字3~15
var IsNumberic = regexp.MustCompile(`^[0-9]*$`)
var IsDate = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2} \d{2}:\d{2}:\d{2}$`) // yyyy-mm-dd hh:mm:ss
var ValidPassword = regexp.MustCompile(`^[a-zA-Z0-9]{5,}$`)                  // パスワード

func IsValidPassword(password string) bool {
	return password != "" && ValidPassword.MatchString(password)
}

func IsValidName(userName string) bool {
	return userName != "" && IsAlphanumeric.MatchString(userName)
}

func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
