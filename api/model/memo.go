package model

import (
	"log"
	"time"

	"zmemo/api/common"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Memo struct {
	ID        string     `gorm:"primary_key" json:"id"`
	UserID    string     `gorm:"primary_key" json:"userID"`
	FolderID  string     `gorm:"unique" json:"folderID"`
	Title     string     `gorm:"not null" json:"title"`
	Text      string     `gorm:"not null" json:"text"`
	CreatedAt time.Time  `gorm:"null" json:"createAt"`
	UpdatedAt time.Time  `gorm:"null" json:"updateAt"`
	DeletedAt *time.Time `gorm:"null" json:"-"`
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
		return memo, common.ErrInvalidMemoID
	}

	// 初期値
	memo.ID = uuid.String()
	memo.CreatedAt = common.GetTime()
	memo.UpdatedAt = common.GetTime()

	if err := m.DB.Create(&memo).Find(&memo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.ErrNotFuondMemo
			log.Println("error: " + err.Error())
		}
		log.Println("error: " + err.Error())
		return memo, err
	}

	return memo, nil
}

func (m *MemoDB) DeleteMemo(userName, memoId string) error {
	memo := Memo{ID: memoId}

	db := m.DB.Find(&memo)

	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return common.ErrNotFuondMemo
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
	memo := Memo{ID: memoId}

	if err := m.DB.Find(&memo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("error: " + err.Error())
			return memo, common.ErrNotFuondMemo
		}
		log.Println("error: " + err.Error())
		return memo, err
	}

	return memo, nil
}

func (m *MemoDB) AddMemoToFolder(userName, memoId, folderName string) error {

	memo := Memo{ID: memoId}

	// ユーザ・フォルダ・メモがある場合、メモ情報を取得
	if err := m.DB.Model(&memo).Select("memos.*").Joins("inner join users on users.user_name = memos.user_name").
		Joins("inner join folders on users.user_name = folders.user_name").First(&memo).Error; err != nil {
		return err
	}

	// フォルダ名がプライマリキーなので、where句に自動的に追加される
	db := m.DB.Table("folders").Where("user_name = ? and memo_id = ?", userName, memoId).UpdateColumn("folder_name", folderName)

	if err := db.Error; err != nil {
		return err
	}

	// if err := db.Scan(&memo).Error; err != nil {
	// 	return err
	// }

	return nil
}

func (m *Memo) Validation() error {

	if m.Title == "" {
		return common.ErrNotFoundTitle
	}

	if m.Text == "" {
		return common.ErrInvalidMemo
	}

	return nil
}
