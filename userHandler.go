package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// ユーザ作成
func (s *Server) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}

		// データバリデーション
		if err := userValidation(user, isCreate); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// ユーザ作成
		db := UserDB{s.DB}
		newUser, err := db.Create(user)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// 登録完了
		return c.JSON(http.StatusOK, newUser)
	}
}

// ユーザ更新
func (s *Server) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}
		user.UserName = c.Param("userName")

		// データバリデーション
		if err := userValidation(user, isUpdate); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// ユーザ情報更新
		db := UserDB{s.DB}
		newUser, err := db.Update(user)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, newUser)
	}
}

// ユーザ削除
func (s *Server) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		// ユーザ名チェック
		if !isValidUserName(userName) {
			return c.JSON(GetErrorCode(InvalidUserName), Message{InvalidUserName.Error()})
		}

		// ユーザ削除
		db := UserDB{s.DB}
		if err := db.Delete(userName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}

// ユーザパスワード変更
func (s *Server) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}
		user.UserName = c.Param("userName")

		// データバリデーション
		if err := userValidation(user, isUpdate); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// ユーザパスワード変更
		db := UserDB{s.DB}
		if err := db.UpdatePassword(user); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}

// ユーザ情報取得
func (s *Server) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		// ユーザ名チェック
		if !isValidUserName(userName) {
			return c.JSON(GetErrorCode(InvalidUserName), Message{InvalidUserName.Error()})
		}

		// ユーザ情報取得
		db := UserDB{s.DB}
		user, err := db.GetUser(userName)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, user)
	}
}

// ユーザ一覧取得
func (s *Server) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO ロール処理を入れる

		db := UserDB{s.DB}
		users, err := db.GetUsers()

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, users)
	}
}

// ログイン
func (s *Server) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}

		// データバリデーション
		if err := userValidation(user, isCreate); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// ユーザ情報取得
		db := UserDB{s.DB}

		if err := db.GetLoginUser(user); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, user)
	}
}
