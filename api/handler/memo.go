package handler

import (
	"log"
	"net/http"

	"zmemo/api/common"
	"zmemo/api/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type memoHandler struct {
	DB *gorm.DB
}

func NewMemoHandler(db *gorm.DB) *memoHandler {
	return &memoHandler{db}
}

// メモ作成
func (s *memoHandler) CreateMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		memo := model.Memo{}

		if err := c.Bind(&memo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		db := model.MemoDB{s.DB}

		// メモ作成
		newMemo, err := db.CreateMemo(memo)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, newMemo)
	}
}

// メモ取得
func (s *memoHandler) GetMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		memoId := c.Param("memoId")

		db := model.MemoDB{s.DB}

		// メモ取得
		memo, err := db.GetMemo(userName, memoId)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, memo)
	}
}

// メモ一覧取得
func (s *memoHandler) GetMemos() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		db := model.MemoDB{s.DB}
		memos, err := db.GetMemos(userName)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, memos)
	}
}

// メモ削除
func (s *memoHandler) DeleteMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		memoId := c.Param("memoId")

		// パラメータバリデーション
		db := model.MemoDB{s.DB}
		if err := db.DeleteMemo(userName, memoId); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}
