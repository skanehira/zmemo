package main

import (
	"regexp"

	uuid "github.com/satori/go.uuid"
)

var isAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`) // 英文字3~15
var isNumberic = regexp.MustCompile(`^[0-9]*$`)
var isDate = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2} \d{2}:\d{2}:\d{2}$`) // yyyy-mm-dd hh:mm:ss
var validPassword = regexp.MustCompile(`^[a-zA-Z0-9]*$`)                     // パスワード
var isUUID = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

// 入力ありの場合はチェックする
func isValidPassword(password string) bool {
	return password != "" && validPassword.MatchString(password)
}

func isValidUserName(userName string) bool {
	return userName != "" && isAlphanumeric.MatchString(userName)
}

func isValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}

func userValidation(user User, mode int) error {
	if !isValidUserName(user.UserName) {
		return InvalidUserName
	}

	// 作成時チェック
	if mode == isCreate {
		// パスワードチェック
		if !isValidPassword(user.Password) {
			return InvalidPassword
		}
	}

	// 更新時チェック
	if mode == isUpdate {
		// 入力があればチェックする
		if user.Password != "" {
			if !isValidPassword(user.Password) {
				return InvalidPassword
			}
		}
	}

	return nil
}
