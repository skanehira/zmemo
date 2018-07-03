package handler

import (
	"net/http"
	"zmemo/api/common"
	"zmemo/api/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type userHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{db}
}

// ユーザ作成
func (s *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
		}

		// データバリデーション

		// ユーザ作成
		db := model.UserDB{s.DB}
		newUser, err := db.Create(user)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		// 登録完了
		return c.JSON(http.StatusOK, newUser)
	}
}

// ユーザ更新
func (s *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
		}
		user.UserName = c.Param("userName")

		// ユーザ情報更新
		db := model.UserDB{s.DB}
		newUser, err := db.Update(user)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, newUser)
	}
}

// ユーザ削除
func (s *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		// ユーザ削除
		db := model.UserDB{s.DB}
		if err := db.Delete(userName); err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}

// ユーザパスワード変更
func (s *userHandler) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
		}
		user.UserName = c.Param("userName")

		// ユーザパスワード変更
		db := model.UserDB{s.DB}
		if err := db.UpdatePassword(user); err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}

// ユーザ情報取得
func (s *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		// ユーザ情報取得
		db := model.UserDB{s.DB}
		user, err := db.GetUser(userName)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, user)
	}
}

// ユーザ一覧取得
func (s *userHandler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO ロール処理を入れる

		db := model.UserDB{s.DB}
		users, err := db.GetUsers()

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, users)
	}
}

// ログイン
func (s *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
		}

		// ユーザ情報取得
		db := model.UserDB{s.DB}

		if err := db.GetLoginUser(user); err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, user)
	}
}
