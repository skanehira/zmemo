package handler

import (
	"net/http"
	"zmemo/api/common"
	"zmemo/api/logger"
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
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(err))
		}

		// データバリデーション
		if err := user.CreateValidation(); err != nil {
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		// ユーザ作成
		model := model.UserDB{DB: s.DB}
		newUser, err := model.CreateUser(user)

		if err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData))
		}
		user.ID = c.Param("userID")

		if err := user.UpdateValidation(); err != nil {
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		// ユーザ情報更新
		model := model.UserDB{DB: s.DB}
		newUser, err := model.UpdateUser(user)

		if err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		return c.JSON(http.StatusOK, newUser)
	}
}

// ユーザ削除
func (s *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")

		if err := model.UserIDValidation(userID); err != nil {
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		// ユーザ削除
		model := model.UserDB{DB: s.DB}

		if err := model.DeleteUser(userID); err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(err))
		}
		user.ID = c.Param("userID")

		if err := model.UserPasswordValidation(user); err != nil {
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		// ユーザパスワード変更
		model := model.UserDB{DB: s.DB}
		if err := model.UpdatePassword(user); err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// ユーザ情報取得
func (s *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")

		if err := model.UserIDValidation(userID); err != nil {
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		// ユーザ情報取得
		model := model.UserDB{DB: s.DB}
		user, err := model.GetUser(userID)

		if err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(err)
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData))
		}

		// ユーザ情報取得
		model := model.UserDB{DB: s.DB}

		if err := model.UserLogin(user); err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		return c.NoContent(http.StatusOK)
	}
}
