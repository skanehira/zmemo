package model

import (
	"time"

	"zmemo/api/common"
	"zmemo/api/logger"

	"github.com/jinzhu/gorm"
)

// User user info
type User struct {
	ID        string     `gorm:"primary_key;not null" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Password  string     `gorm:"not null" json:"password"`
	Memos     []Memo     `gorm:"null" json:"memos"`
	Folders   []Folder   `gorm:"null" json:"folders"`
	CreatedAt time.Time  `gorm:"null" json:"createAt"`
	UpdatedAt time.Time  `gorm:"null" json:"updateaAt"`
	DeletedAt *time.Time `gorm:"null" json:"-"`
}

type Users []User

// UserDB db
type UserDB struct {
	DB *gorm.DB
}

// Validation user data validate
func (u *User) Validation() error {
	if u.Password != "" || common.ValidPassword.MatchString(u.Password) {
		return common.ErrInvalidPassword

	}

	if u.Name != "" || common.IsAlphanumeric.MatchString(u.Name) {
		return common.ErrInvalidUserName
	}

	return nil
}

// CreateUser ユーザ作成
func (d *UserDB) CreateUser(newUser User) (User, error) {
	logger.Info("CreateUser() is start")

	// 初期値
	newUser.ID = common.NewUUID()
	newUser.CreatedAt = common.GetTime()
	newUser.UpdatedAt = common.GetTime()
	newUser.Folders = Folders{}
	newUser.Memos = Memos{}

	// ユーザ登録
	if err := d.DB.Create(&newUser).Error; err != nil {
		return newUser, common.WrapError(err)
	}

	newUser, err := d.GetUser(newUser.ID)
	if err != nil {
		return newUser, common.WrapError(err)
	}

	logger.Info("CreateUser() is end")

	return newUser, nil
}

// GetUser ユーザ情報取得
func (d *UserDB) GetUser(id string) (User, error) {
	logger.Info("GetUSer() is start")
	user := User{ID: id}
	memos := Memos{}
	folders := Folders{}

	// ユーザが存在しない
	if err := d.DB.Find(&user).Related(&memos).Related(&folders).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.ErrNotFoundUser
		}

		return user, common.WrapError(err)
	}

	user.Memos = memos
	user.Folders = folders

	logger.Info("GetUSer() is end")
	return user, nil
}

// UpdateUser ユーザ更新
func (d *UserDB) UpdateUser(user User) (User, error) {
	logger.Info("UpdateUser() is start")

	// ユーザが存在しない場合はエラーを返す
	_, err := d.GetUser(user.ID)
	if err != nil {
		return user, err
	}

	if err := d.DB.Model(&user).Updates(map[string]interface{}{"name": user.Name, "updated_at": common.GetTime()}).Error; err != nil {
		return user, common.WrapError(err)
	}

	// 更新後のデータを返却する
	newUser, err := d.GetUser(user.ID)
	if err != nil {
		return user, err
	}

	logger.Info("UpdateUser() is end")
	return newUser, nil
}

// DeleteUser ユーザ削除
func (d *UserDB) DeleteUser(id string) error {
	logger.Info("DeleteUser() is start")

	// ユーザが存在しない
	user, err := d.GetUser(id)
	if err != nil {
		return err
	}

	// start transaction
	db := d.DB.Begin()

	m := MemoDB{DB: db}
	f := FolderDB{DB: db}

	// delete memo
	if err := m.DeleteAllMemo(id); err != nil {
		if err := db.Rollback().Error; err != nil {
			return common.WrapError(err)
		}
		return common.WrapError(err)
	}

	// delete folder
	if err := f.DeleteAllFolder(id); err != nil {
		if err := db.Rollback().Error; err != nil {
			return common.WrapError(err)
		}

		return common.WrapError(err)
	}

	// delete user
	if err := db.Delete(&user).Error; err != nil {
		if err := db.Rollback().Error; err != nil {
			return common.WrapError(err)
		}

		return common.WrapError(err)
	}

	// db commit
	if err := db.Commit().Error; err != nil {
		return common.WrapError(err)
	}

	logger.Info("DeleteUser() is end")
	return nil
}

// UpdatePassword パスワード更新
func (d *UserDB) UpdatePassword(user User) error {
	// ユーザが存在しない
	_, err := d.GetUser(user.ID)
	if err != nil {
		return err
	}

	newData := map[string]interface{}{"updated_at": common.GetTime(), "password": user.Password}

	if err := d.DB.Model(&user).Where("id = ?", user.ID).Updates(newData).Error; err != nil {
		return common.WrapError(err)
	}

	return nil
}

// UserList ユーザ一覧取得
func (d *UserDB) UserList() (Users, error) {
	users := Users{}

	if err := d.DB.Preload("Folders").Preload("Memos").Find(&users).Error; err != nil {
		return users, common.WrapError(err)
	}

	return users, nil
}

// UserLogin ログインユーザ情報取得
func (d *UserDB) UserLogin(user User) error {

	// ユーザIDとパスワード
	if err := d.DB.Model(&user).Where("id = ? and password = ?", user.ID, user.Password).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.ErrNotFoundUser
		}
		return common.WrapError(err)
	}

	return nil
}
