package main

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type date time.Time

type Memo struct {
	UserID    string  `gorm:"primary_key";json:"userId"`
	MemoID    string  `gorm:"primary_key";json:"memoId"`
	Text      string  `gorm:"not null";json:"text"`
	CreatedAt string  `gorm:"not null";json:"createAt"`
	UpdatedAt string  `gorm:"not null";json:"updateAt"`
	DeletedAt *string `gorm:"null";json:"-"`
}

type Memos []Memo

type MemoDB struct {
	DB *gorm.DB
}

func (m *MemoDB) CreateMemo(memo Memo) (Memo, error) {
	// メモID生成
	uuid, err := uuid.NewV4()
	if err != nil {
		return memo, InvalidMemoID
	}

	// 初期値
	memo.MemoID = uuid.String()
	memo.CreatedAt = GetTime()
	memo.UpdatedAt = GetTime()

	if err := m.DB.Create(&memo).Find(&memo).Error; err != nil {
		return memo, err
	}

	return memo, nil
}

func (m *MemoDB) DeleteMemo(userId, memoId string) error {
	memo := Memo{UserID: userId, MemoID: memoId}

	db := m.DB.Find(&memo)

	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFuondMemo
		}

		return err
	}

	if err := db.Delete(&memo).Error; err != nil {
		return err
	}

	return nil
}

func (m *MemoDB) GetMemos(userId string) (Memos, error) {

	memos := Memos{}

	if err := m.DB.Model(Memo{}).Where("user_id = ?", userId).Scan(&memos).Error; err != nil {
		return memos, err
	}

	return memos, nil
}

func (m *MemoDB) GetMemo(userId, memoId string) (Memo, error) {
	memo := Memo{UserID: userId, MemoID: memoId}

	if err := m.DB.Find(&memo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return memo, NotFuondMemo
		}
		return memo, err
	}

	return memo, nil
}
