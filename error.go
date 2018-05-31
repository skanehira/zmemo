package main

import (
	"errors"
	"net/http"
)

var (
	// ユーザ系エラー
	NotFoundUserName  = errors.New("ユーザ名がありません")
	InvalidUserName   = errors.New("英数字のみを使用して下さい")
	NotFoundPassword  = errors.New("パスワードがありません")
	InvalidPassword   = errors.New("パスワードが正しくありません")
	NotFoundDate      = errors.New("日付がありません")
	InvalidDate       = errors.New("日付が正しくありません")
	NotFoundUser      = errors.New("ユーザがありません")
	NotFuondMemoID    = errors.New("メモIDがありません")
	InvalidMemoID     = errors.New("メモIDが正しくありません")
	InvalidMemo       = errors.New("メモ内がが正しくありません")
	NotFuondMemo      = errors.New("メモがありません")
	NotFoundFolder    = errors.New("フォルダがありません")
	InvalidFolderName = errors.New("フォルダ名が正しくありません")
	NotFoundTitle     = errors.New("タイトルがありません")

	InvalidPostData = errors.New("データ形式が正しくありません")
)

// エラー型からレスポンスコードを取得
func GetErrorCode(err error) int {
	switch err {
	case NotFoundUserName, InvalidUserName,
		NotFoundDate, InvalidDate, NotFoundPassword, InvalidPassword,
		InvalidPostData, NotFuondMemoID, InvalidMemoID, InvalidMemo, NotFuondMemo,
		InvalidFolderName, NotFoundTitle:
		return http.StatusBadRequest
	case NotFoundUser, NotFuondMemo, NotFoundFolder:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

type Message struct {
	Message string `json:"message"`
}
