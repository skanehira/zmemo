package model

import (
	"time"
	"zmemo/api/common"

	"github.com/jinzhu/gorm"
)

type Folder struct {
	ID         string     `gorm:"primary_key" json:"id"`
	UserID     string     `gorm:"primary_key" json:"userId"`
	FolderName string     `gorm:"primary_key" json:"FolderName"`
	CreatedAt  time.Time  `gorm:"null" json:"createAt"`
	UpdatedAt  time.Time  `gorm:"null" json:"updateAt"`
	DeletedAt  *time.Time `gorm:"null" json:"deletedAt"`
}

type FolderDB struct {
	DB *gorm.DB
}

// Folders type folder list
type Folders []Folder

// CreateFolder create new folder
func (d *FolderDB) CreateFolder(f Folder) (Folder, error) {
	// 初期値
	f.ID = common.NewUUID()
	f.CreatedAt = common.GetTime()
	f.UpdatedAt = common.GetTime()

	if err := d.DB.Create(&f).Error; err != nil {
		return f, err
	}

	if newFolder, err := d.GetFolder(f.UserID, f.ID); err != nil {
		return newFolder, err
	}

	return f, nil
}

// GetFolder get folder
func (d *FolderDB) GetFolder(userID, folderID string) (Folder, error) {

	f := Folder{
		ID:     folderID,
		UserID: userID,
	}

	if err := d.DB.Find(&f).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return f, common.ErrNotFoundFolder
		}
		return f, err
	}

	return f, nil
}

// FolderList get folder list
func (d *FolderDB) FolderList(userID string) (Folders, error) {
	folders := Folders{}

	if err := d.DB.Model(&Folder{}).Where("user_id = ?", userID).Scan(&folders).Error; err != nil {
		return folders, err
	}

	return folders, nil
}

// UpdateFolder update folder info
func (d *FolderDB) UpdateFolder(f Folder) error {

	// フォルダがない場合はエラーを返す
	if _, err := d.GetFolder(f.UserID, f.ID); err != nil {
		return err
	}

	// 更新データ
	newData := map[string]interface{}{"folder_name": f.FolderName, "updated_at": common.GetTime()}

	// フォルダを更新
	if err := d.DB.Model(&f).Updates(newData).Error; err != nil {
		return err
	}

	return nil
}

// DeleteFolder delete folder
func (d *FolderDB) DeleteFolder(userID, folderID string) error {
	folder := Folder{
		ID:     folderID,
		UserID: userID,
	}

	if _, err := d.GetFolder(userID, folderID); err != nil {
		return err
	}

	if err := d.DB.Delete(&folder).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAllFolder delete all folder
func (d *FolderDB) DeleteAllFolder(userID string) error {
	folder := Folder{UserID: userID}

	if err := d.DB.Delete(&folder).Error; err != nil {
		return err
	}

	return nil
}
