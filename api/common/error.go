package common

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	ErrNotFoundUserName  = errors.New("ユーザ名がありません")
	ErrInvalidUserName   = errors.New("英数字のみを使用して下さい")
	ErrNotFoundPassword  = errors.New("パスワードがありません")
	ErrInvalidPassword   = errors.New("パスワードが正しくありません")
	ErrNotFoundDate      = errors.New("日付がありません")
	ErrInvalidDate       = errors.New("日付が正しくありません")
	ErrNotFoundUser      = errors.New("ユーザがありません")
	ErrNotFuondMemoID    = errors.New("メモIDがありません")
	ErrInvalidMemoID     = errors.New("メモIDが正しくありません")
	ErrInvalidMemo       = errors.New("メモ内がが正しくありません")
	ErrNotFuondMemo      = errors.New("メモがありません")
	ErrNotFoundFolder    = errors.New("フォルダがありません")
	ErrInvalidFolderName = errors.New("フォルダ名が正しくありません")
	ErrNotFoundTitle     = errors.New("タイトルがありません")
	ErrInvalidPostData   = errors.New("データ形式が正しくありません")
)

// ErrorMessage エラー構造体
type ErrorMessage struct {
	Message interface{} `json:"message"`
}

// NewError 新しいエラーを生成
func NewError(err error) ErrorMessage {
	return ErrorMessage{
		Message: err.Error(),
	}
}

// WrapError error Wrapper
func WrapError(err error) error {
	return errors.Wrap(err, "")
}

// Error エラーメッセージを出力
func (err ErrorMessage) Error() string {
	return fmt.Sprintf("error : %s", err.Message)
}

// エラー型からレスポンスコードを取得
func GetErrorCode(err error) int {
	switch err {
	case ErrInvalidUserName,
		ErrInvalidDate, ErrInvalidPassword,
		ErrInvalidPostData, ErrInvalidMemoID, ErrInvalidMemo,
		ErrInvalidFolderName:
		return http.StatusBadRequest
	case ErrNotFoundUserName, ErrNotFoundDate,
		ErrNotFoundPassword, ErrNotFuondMemoID, ErrNotFuondMemo, ErrNotFoundTitle:
		return http.StatusNotFound
	case ErrNotFoundUser, ErrNotFuondMemo, ErrNotFoundFolder:
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
