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

		model := model.MemoDB{DB: s.DB}

		// メモ作成
		newMemo, err := model.CreateMemo(memo)

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
		userID := c.Param("userID")
		memoID := c.Param("memoID")

		model := model.MemoDB{DB: s.DB}

		// メモ取得
		memo, err := model.GetMemo(userID, memoID)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, memo)
	}
}

// メモ一覧取得
func (s *memoHandler) MemoList() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")

		model := model.MemoDB{DB: s.DB}
		memos, err := model.MemoList(userID)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, memos)
	}
}

// メモ更新
func (s *memoHandler) UpdateMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		newMemo := model.Memo{}

		if err := c.Bind(&newMemo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		newMemo.ID = c.Param("memoID")
		newMemo.UserID = c.Param("userID")

		model := model.MemoDB{DB: s.DB}
		newMemo, err := model.UpdateMemo(newMemo)

		if err != nil {
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, newMemo)
	}
}

// メモ削除
func (s *memoHandler) DeleteMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")
		memoID := c.Param("memoID")

		model := model.MemoDB{DB: s.DB}
		if err := model.DeleteMemo(userID, memoID); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}

// フォルダ追加
func (s *memoHandler) AddMemoToFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		memo := model.Memo{}

		if err := c.Bind(&memo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		memo.ID = c.Param("memoID")
		memo.UserID = c.Param("userID")

		model := model.MemoDB{DB: s.DB}

		if err := model.AddMemoToFolder(memo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}
