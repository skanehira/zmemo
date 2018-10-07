package model

import (
	"time"

	"zmemo/api/common"

	"github.com/jinzhu/gorm"
)

type Memo struct {
	ID        string     `gorm:"primary_key" json:"id"`
	UserID    string     `gorm:"primary_key" json:"userID"`
	FolderID  *string    `gorm:"unique;null" json:"folderID,omitempty"`
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

// Validation memo dta validate
func (m *Memo) Validation() {

}

// CreateMemo create new memo
func (d *MemoDB) CreateMemo(newMemo Memo) (Memo, error) {
	// 初期値
	newMemo.ID = common.NewUUID()
	newMemo.CreatedAt = time.Now()
	newMemo.UpdatedAt = time.Now()

	if err := d.DB.Create(&newMemo).Error; err != nil {
		return newMemo, common.Wrap(err)
	}

	newMemo, err := d.GetMemo(newMemo.UserID, newMemo.ID)
	if err != nil {
		return newMemo, err
	}

	return newMemo, nil
}

// GetMemo get memo
func (d *MemoDB) GetMemo(userID, memoID string) (Memo, error) {
	memo := Memo{ID: memoID, UserID: userID}

	if err := d.DB.Find(&memo).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.ErrNotFuondMemo
		}
		return memo, common.Wrap(err)
	}

	return memo, nil
}

// MemoList get memo list
func (d *MemoDB) MemoList(userID string) (Memos, error) {

	memos := Memos{}

	if err := d.DB.Model(Memo{}).Where("user_id = ?", userID).Scan(&memos).Error; err != nil {
		return memos, common.Wrap(err)
	}

	return memos, nil
}

// UpdateMemo update memo info
func (d *MemoDB) UpdateMemo(memo Memo) (Memo, error) {

	// メモがない場合はエラーを返す
	if _, err := d.GetMemo(memo.UserID, memo.ID); err != nil {
		return memo, err
	}

	// メモデータ更新
	newData := map[string]interface{}{"updated_at": time.Now()}

	if memo.Text != "" {
		newData["text"] = memo.Text
	}

	if memo.Title != "" {
		newData["title"] = memo.Title
	}

	if err := d.DB.Model(&memo).Where("id = ? and user_id = ?", memo.ID, memo.UserID).Updates(newData).Error; err != nil {
		return memo, common.Wrap(err)
	}

	// 更新後データを取得
	memo, err := d.GetMemo(memo.UserID, memo.ID)
	if err != nil {
		return memo, err
	}

	return memo, nil
}

// DeleteMemo delete memo
func (d *MemoDB) DeleteMemo(userID, memoID string) error {

	if _, err := d.GetMemo(userID, memoID); err != nil {
		return err
	}

	memo := Memo{ID: memoID, UserID: userID}

	if err := d.DB.Delete(&memo).Error; err != nil {
		return common.Wrap(err)
	}

	return nil
}

// DeleteAllMemo delete all memo
func (d *MemoDB) DeleteAllMemo(userID string) error {
	memo := Memo{UserID: userID}

	if err := d.DB.Delete(&memo).Error; err != nil {
		return common.Wrap(err)
	}

	return nil
}

// AddMemoToFolder add memo to target folder
func (d *MemoDB) AddMemoToFolder(m Memo) error {
	// フォルダ存在確認
	f := Folder{ID: *m.FolderID}
	if err := d.DB.Find(&f).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.ErrNotFoundFolder
		}
		return common.Wrap(err)
	}

	// update memo
	if err := d.DB.Model(&m).UpdateColumns(common.StructToMap(&m)).Error; err != nil {
		return common.Wrap(err)
	}

	return nil
}
