package model

import (
	"time"
	"zmemo/api/common"

	"github.com/jinzhu/gorm"
)

type Folder struct {
	ID         string     `gorm:"primary_key;not null" json:"id"`
	FolderName string     `gorm:"not null" json:"FolderName"`
	CreatedAt  time.Time  `gorm:"null" json:"createAt"`
	UpdatedAt  time.Time  `gorm:"null" json:"updateAt"`
	DeletedAt  *time.Time `gorm:"null" json:"deletedAt"`
}

type FolderDB struct {
	DB *gorm.DB
}

type Folders []Folder

func (d *FolderDB) CreateFolder(folder Folder) (Folder, error) {
	// 初期値
	folder.CreatedAt = common.GetTime()
	folder.UpdatedAt = common.GetTime()

	if err := d.DB.Create(&folder).Find(&folder).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return folder, common.ErrNotFoundFolder
		}
		return folder, err
	}

	return folder, nil
}

func (d *FolderDB) GetFolder(folderName string) (Folder, error) {

	folder := Folder{
		FolderName: folderName,
	}
	if err := d.DB.Find(&folder).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return folder, common.ErrNotFoundFolder
		}
		return folder, err
	}

	return folder, nil
}

func (d *FolderDB) GetFolders(userName string) (Folders, error) {
	folders := Folders{}

	if err := d.DB.Model(&Folder{}).Where("user_name = ?", userName).Scan(&folders).Error; err != nil {
		return folders, err
	}

	return folders, nil
}

func (d *FolderDB) UpdateFolderName(userName, folderName string) error {
	newFolder := Folder{
		FolderName: folderName,
		UpdatedAt:  common.GetTime(),
	}

	if err := d.DB.Model(&newFolder).UpdateColumns(common.StructToMap(&newFolder)).Error; err != nil {
		return err
	}

	return nil
}

func (d *FolderDB) DeleteFolder(folderName string) error {
	folder := Folder{FolderName: folderName}

	db := d.DB.Find(&folder)

	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return common.ErrNotFoundFolder
		}
	}

	if err := d.DB.Delete(&folder).Error; err != nil {
		return err
	}

	return nil
}

func (f *Folder) FolderValidation(folderName string) error {

	if folderName == "" {
		return common.ErrInvalidFolderName
	}

	return nil
}
