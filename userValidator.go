package main

import (
	"regexp"
)

var isAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`) // 英文字3~15
var isNumberic = regexp.MustCompile(`^[0-9]*$`)
var isDate = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2} \d{2}:\d{2}:\d{2}$`) // yyyy-mm-dd hh:mm:ss
var validPassword = regexp.MustCompile(`^[a-zA-Z0-9]*$`)                     // パスワード
var isUUID = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}`)

// 入力ありの場合はチェックする
func isValidPassword(password string) bool {

	return true
}

func isValidUserId(userId string) bool {
	return userId != "" || isUUID.MatchString(userId)
}

func userValidation(user User, mode int) error {

	// 更新時チェック
	if mode == isUpdate {
		if user.UserID == "" {
			return NotFoundUserID
		}
		if !isUUID.MatchString(user.UserID) {
			return InvalidUserID
		}

		// 入力があればチェックする
		if user.Password != "" && !validPassword.MatchString(user.Password) {
			return InvalidPassword
		}

		// 入力があればチェックする
		if user.UserName != "" && !isAlphanumeric.MatchString(user.UserName) {
			return InvalidUserName
		}
	}

	// 作成時チェック
	if mode == isCreate {
		// ユーザ名チェック
		if user.UserName == "" {
			return NotFoundUserName
		}
		if !isAlphanumeric.MatchString(user.UserName) {
			return InvalidUserName
		}

		// パスワードチェック
		if user.Password == "" {
			return NotFoundPassword
		}
		if !validPassword.MatchString(user.Password) {
			return InvalidPassword
		}
	}

	return nil
}
