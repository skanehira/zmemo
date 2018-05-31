package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// メモ作成
func (s *Server) CreateMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		memo := Memo{}

		if err := c.Bind(&memo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// メモデータバリデーション
		if err := memoValidation(memo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := MemoDB{s.DB}

		// メモ作成
		newMemo, err := db.CreateMemo(memo)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, newMemo)
	}
}

// メモ取得
func (s *Server) GetMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		memoId := c.Param("memoId")

		// パラメータバリデーション
		if !isValidUserName(userName) {
			log.Println("error: " + InvalidUserName.Error())
			return c.JSON(GetErrorCode(InvalidUserName), Message{InvalidUserName.Error()})
		}
		if !isValidMemoId(memoId) {
			log.Println("error: " + InvalidUserName.Error())
			return c.JSON(GetErrorCode(InvalidMemo), Message{InvalidMemoID.Error()})
		}

		db := MemoDB{s.DB}

		// メモ取得
		memo, err := db.GetMemo(userName, memoId)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, memo)
	}
}

// メモ一覧取得
func (s *Server) GetMemos() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		// パラメータバリデーション
		if !isValidUserName(userName) {
			log.Println("error: " + InvalidUserName.Error())
			return c.JSON(GetErrorCode(InvalidUserName), Message{InvalidUserName.Error()})
		}

		db := MemoDB{s.DB}
		memos, err := db.GetMemos(userName)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, memos)
	}
}

// メモ削除
func (s *Server) DeleteMemo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		memoId := c.Param("memoId")

		// パラメータバリデーション
		if !isValidUserName(userName) {
			log.Println("error: " + InvalidUserName.Error())
			return c.JSON(GetErrorCode(InvalidUserName), Message{InvalidUserName.Error()})
		}
		if !isValidMemoId(memoId) {
			log.Println("error: " + InvalidMemo.Error())
			return c.JSON(GetErrorCode(InvalidMemo), Message{InvalidMemoID.Error()})
		}

		db := MemoDB{s.DB}
		if err := db.DeleteMemo(userName, memoId); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}
