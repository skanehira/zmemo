package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type date time.Time

type Memo struct {
	UserName   string  `gorm:"primary_key" json:"userName"`
	MemoID     string  `gorm:"primary_key" json:"memoId"`
	FolderName *string `gorm:"unique" json:"folderName"`
	Title      string  `gorm:"not null" json:"title"`
	Text       string  `gorm:"not null" json:"text"`
	CreatedAt  string  `gorm:"not null" json:"createAt"`
	UpdatedAt  string  `gorm:"not null" json:"updateAt"`
	DeletedAt  *string `gorm:"null" json:"-"`
}

type Memos []Memo

type MemoDB struct {
	DB *gorm.DB
}

func (m *MemoDB) CreateMemo(memo Memo) (Memo, error) {
	// メモID生成
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("error: " + err.Error())
		return memo, InvalidMemoID
	}

	// 初期値
	memo.MemoID = uuid.String()
	memo.CreatedAt = GetTime()
	memo.UpdatedAt = GetTime()

	if err := m.DB.Create(&memo).Find(&memo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = NotFuondMemo
			log.Println("error: " + err.Error())
		}
		log.Println("error: " + err.Error())
		return memo, err
	}

	return memo, nil
}

func (m *MemoDB) DeleteMemo(userName, memoId string) error {
	memo := Memo{UserName: userName, MemoID: memoId}

	db := m.DB.Find(&memo)

	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFuondMemo
		}
		log.Println("error: " + err.Error())

		return err
	}

	if err := db.Delete(&memo).Error; err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (m *MemoDB) GetMemos(userName string) (Memos, error) {

	memos := Memos{}

	if err := m.DB.Model(Memo{}).Where("user_name = ?", userName).Scan(&memos).Error; err != nil {
		log.Println("error: " + err.Error())
		return memos, err
	}

	return memos, nil
}

func (m *MemoDB) GetMemo(userName, memoId string) (Memo, error) {
	memo := Memo{UserName: userName, MemoID: memoId}

	if err := m.DB.Find(&memo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("error: " + err.Error())
			return memo, NotFuondMemo
		}
		log.Println("error: " + err.Error())
		return memo, err
	}

	return memo, nil
}

func (s *Server) AddMemoToFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		memo := struct {
			UserName   string
			FolderName string
			MemoID     string
		}{
			c.Param("userName"),
			c.Param("folderName"),
			c.QueryParam("memoId"),
		}

		if err := c.Bind(&memo); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}

		// バリデーション
		if err := folderValidation(memo.UserName, memo.FolderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		if !isValidMemoId(memo.MemoID) {
			log.Println("error: " + InvalidMemoID.Error())
			return c.JSON(GetErrorCode(InvalidMemoID), Message{InvalidMemoID.Error()})
		}

		db := FolderDB{s.DB}
		if err := db.AddMemoToFolder(memo.UserName, memo.MemoID, memo.FolderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)

	}
}
