package main

import (
	"github.com/jinzhu/gorm"
)

type Folder struct {
	UserName   string  `gorm:"primary_key" json:"userName"`
	FolderName string  `gorm:"primary_key" json:"folderName"`
	CreatedAt  string  `gorm:"not null" json:"createAt"`
	UpdatedAt  string  `gorm:"not null" json:"updateAt"`
	DeletedAt  *string `gorm:"" json:"deletedAt"`
}

type FolderDB struct {
	DB *gorm.DB
}

type Folders []Folder

func (d *FolderDB) CreateFolder(folder Folder) (Folder, error) {
	// 初期値
	folder.CreatedAt = GetTime()
	folder.UpdatedAt = GetTime()

	if err := d.DB.Create(&folder).Find(&folder).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return folder, NotFoundFolder
		}
		return folder, err
	}

	return folder, nil
}

func (d *FolderDB) GetFolder(userName, folderName string) (Memos, error) {

	memos := new(Memos)
	if err := d.DB.Table("memos").Joins("inner join users on users.user_name = memos.user_name").
		Joins("inner join folders on folders.user_name = users.user_name").Where("memos.user_name = ? and memos.folder_name = ?", userName, folderName).Find(memos).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return *memos, NotFoundFolder
		}
		return *memos, err
	}

	return *memos, nil
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
		UserName:   userName,
		FolderName: folderName,
		UpdatedAt:  GetTime(),
	}

	if err := d.DB.Model(&newFolder).UpdateColumns(StructToMap(&newFolder)).Error; err != nil {
		return err
	}

	return nil
}

func (d *FolderDB) DeleteFolder(userName, folderName string) error {
	folder := Folder{UserName: userName, FolderName: folderName}

	db := d.DB.Find(&folder)

	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return NotFoundFolder
		}
	}

	if err := d.DB.Delete(&folder).Error; err != nil {
		return err
	}

	return nil
}

func (d *FolderDB) AddMemoToFolder(userName, memoId, folderName string) error {

	memo := Memo{UserName: userName, MemoID: memoId}

	// ユーザ・フォルダ・メモがある場合、メモ情報を取得
	if err := d.DB.Model(&memo).Select("memos.*").Joins("inner join users on users.user_name = memos.user_name").
		Joins("inner join folders on users.user_name = folders.user_name").First(&memo).Error; err != nil {
		return err
	}

	// フォルダ名がプライマリキーなので、where句に自動的に追加される
	db := d.DB.Table("folders").Where("user_name = ? and memo_id = ?", userName, memoId).UpdateColumn("folder_name", folderName)

	if err := db.Error; err != nil {
		return err
	}

	// if err := db.Scan(&memo).Error; err != nil {
	// 	return err
	// }

	return nil
}
