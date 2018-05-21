package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// メモ作成
func (s *Server) CreateMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		memo := Memo{}

		if err := c.Bind(&memo); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// メモデータバリデーション
		if err := memoValidation(memo); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := MemoDB{s.DB}

		// メモ作成
		newMemo, err := db.CreateMemo(memo)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, newMemo)
	}
}

// メモ取得
func (s *Server) GetMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		memoId := c.Param("memoId")

		// パラメータバリデーション
		if !isValidUserId(userId) {
			return c.JSON(GetErrorCode(InvalidUserID), Message{InvalidUserID.Error()})
		}
		if !isValidMemoId(memoId) {
			return c.JSON(GetErrorCode(InvalidMemo), Message{InvalidMemoID.Error()})
		}

		db := MemoDB{s.DB}

		// メモ取得
		memo, err := db.GetMemo(userId, memoId)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, memo)
	}
}

// メモ一覧取得
func (s *Server) GetMemos() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")

		// パラメータバリデーション
		if !isValidUserId(userId) {
			return c.JSON(GetErrorCode(InvalidUserID), Message{InvalidUserID.Error()})
		}

		db := MemoDB{s.DB}
		memos, err := db.GetMemos(userId)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, memos)
	}
}

// メモ削除
func (s *Server) DeleteMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		memoId := c.Param("memoId")

		// パラメータバリデーション
		if !isValidUserId(userId) {
			return c.JSON(GetErrorCode(InvalidUserID), Message{InvalidUserID.Error()})
		}
		if !isValidMemoId(memoId) {
			return c.JSON(GetErrorCode(InvalidMemo), Message{InvalidMemoID.Error()})
		}

		db := MemoDB{s.DB}
		if err := db.DeleteMemo(userId, memoId); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}
