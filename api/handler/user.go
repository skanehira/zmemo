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

// NewUserHandler new handler
func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{db}
}

// ユーザ作成
func (s *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}

		// jsonデータ取得
		if err := c.Bind(&user); err != nil {
			// return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(err.Error()))
		}

		// データバリデーション

		// ユーザ作成
		model := model.UserDB{DB: s.DB}
		newUser, err := model.CreateUser(user)

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
		user.ID = c.Param("userID")

		// ユーザ情報更新
		model := model.UserDB{DB: s.DB}
		newUser, err := model.UpdateUser(user)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, newUser)
	}
}

// ユーザ削除
func (s *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")

		// ユーザ削除
		model := model.UserDB{DB: s.DB}

		if err := model.DeleteUser(userID); err != nil {
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
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(err.Error()))
		}
		user.ID = c.Param("userID")

		// ユーザパスワード変更
		model := model.UserDB{DB: s.DB}
		if err := model.UpdatePassword(user); err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}

// ユーザ情報取得
func (s *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")

		// ユーザ情報取得
		model := model.UserDB{DB: s.DB}
		user, err := model.GetUser(userID)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, user)
	}
}

// ユーザ一覧取得
func (s *userHandler) UserList() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO ロール処理を入れる

		model := model.UserDB{DB: s.DB}
		users, err := model.UserList()

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
		model := model.UserDB{DB: s.DB}

		if err := model.UserLogin(user); err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}
