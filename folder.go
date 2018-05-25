package main

import (
	"github.com/jinzhu/gorm"
)

type Folder struct {
	UserID     string  `gorm:"primary_key";json:"userId"`
	FolderName string  `gorm:"primary_key";json:"folderName"`
	CreatedAt  string  `gorm:"not null";json:"createAt"`
	UpdatedAt  string  `gorm:"not null";json:"updateAt"`
	DeletedAt  *string `gorm:"";json:"deletedAt"`
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

func (d *FolderDB) GetFolder(userId, folderName string) (Memos, error) {

	memos := new(Memos)
	if err := d.DB.Table("memos").Joins("inner join users on users.user_id = memos.user_id").
		Joins("inner join folders on folders.user_id = users.user_id").Where("memos.user_id = ? and memos.folder_name = ?", userId, folderName).Find(memos).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return *memos, NotFoundFolder
		}
		return *memos, err
	}

	return *memos, nil
}

func (d *FolderDB) GetFolders(userId string) (Folders, error) {
	folders := Folders{}

	if err := d.DB.Model(&Folder{}).Where("user_id = ?", userId).Scan(&folders).Error; err != nil {
		return folders, err
	}

	return folders, nil
}

func (d *FolderDB) UpdateFolderName(userId, folderName string) error {
	newFolder := Folder{
		UserID:     userId,
		FolderName: folderName,
		UpdatedAt:  GetTime(),
	}

	if err := d.DB.Model(&newFolder).UpdateColumns(StructToMap(&newFolder)).Error; err != nil {
		return err
	}

	return nil
}

func (d *FolderDB) DeleteFolder(userId, folderName string) error {
	folder := Folder{UserID: userId, FolderName: folderName}

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

func (d *FolderDB) AddMemoToFolder(userId, memoId, folderName string) error {

	// folder := Folder{UserID: userId, FolderName: folderName}
	memo := Memo{UserID: userId, MemoID: memoId}

	// ユーザ・フォルダ・メモがある場合、メモ情報を取得
	// select memos.* from memos inner join users on users.user_id = memos.user_id inner join folders on users.user_id = folders.user_id
	if err := d.DB.Model(&memo).Select("memos.*").Joins("inner join users on users.user_id = memos.user_id").
		Joins("inner join folders on users.user_id = folders.user_id").First(&memo).Error; err != nil {
		return err
	}

	// フォルダ名がプライマリキーなので、where句に自動的に追加される
	// UPDATE `memos` SET `folder_name` = 'テストフォルダ3'  WHERE `memos`.`deleted_at` IS NULL AND `memos`.`user_id` = '8b03695b-bf4d-4a94-802d-13057ecf6d62' AND `memos`.`memo_id` = '69958eeb-5f31-42c8-9290-9173f801ac2c' AND `memos`.`folder_name` = 'テストフォルダ3'
	db := d.DB.Table("folders").Where("user_id = ? and memo_id = ?", userId, memoId).UpdateColumn("folder_name", folderName)

	if err := db.Error; err != nil {
		return err
	}

	// if err := db.Scan(&memo).Error; err != nil {
	// 	return err
	// }

	return nil
}
