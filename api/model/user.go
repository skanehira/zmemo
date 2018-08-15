package model

import (
	"log"
	"time"

	"zmemo/api/common"

	"github.com/jinzhu/gorm"
)

// User user info
type User struct {
	ID        string     `gorm:"primary_key;not null" json:"id"`
	UserName  string     `gorm:"not null" json:"userName"`
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

// CreateUser ユーザ作成
func (d *UserDB) CreateUser(newUser User) (User, error) {
	// 初期値
	newUser.ID = common.NewUUID()
	newUser.CreatedAt = common.GetTime()
	newUser.UpdatedAt = common.GetTime()
	newUser.Folders = Folders{}
	newUser.Memos = Memos{}

	// ユーザ登録
	if err := d.DB.Create(&newUser).Error; err != nil {
		log.Println("error: " + err.Error())
		return newUser, err
	}

	newUser, err := d.GetUser(newUser.ID)
	if err != nil {
		log.Println("error: " + err.Error())
		return newUser, err
	}

	return newUser, nil
}

// GetUser ユーザ情報取得
func (d *UserDB) GetUser(id string) (User, error) {
	user := User{ID: id}
	memos := Memos{}
	folders := Folders{}

	// ユーザが存在しない
	if err := d.DB.Find(&user).Related(&memos).Related(&folders).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("error: " + common.ErrNotFoundUser.Error())
			return user, common.ErrNotFoundUser
		}

		log.Println("error: " + err.Error())
		return user, err
	}

	user.Memos = memos
	user.Folders = folders

	return user, nil
}

// UpdateUser ユーザ更新
func (d *UserDB) UpdateUser(user User) (User, error) {
	// ユーザが存在しない場合はエラーを返す
	oldUser, err := d.GetUser(user.ID)
	if err != nil {
		return user, err
	}

	// 更新日取得
	newTime := common.GetTime()
	user.UpdatedAt = newTime
	user.CreatedAt = oldUser.CreatedAt

	if err := d.DB.Model(&user).Updates(user).Error; err != nil {
		return user, err
	}

	// 更新後のデータを返却する
	newUser, err := d.GetUser(user.ID)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

// DeleteUser ユーザ削除
func (d *UserDB) DeleteUser(id string) error {

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
		log.Println("error: " + err.Error())
		if err := db.Rollback().Error; err != nil {
			return err
		}

		return err
	}

	// delete folder
	if err := f.DeleteAllFolder(id); err != nil {
		log.Println("error: " + err.Error())
		if err := db.Rollback().Error; err != nil {
			return err
		}

		return err
	}

	// delete user
	if err := db.Delete(&user).Error; err != nil {
		log.Println("error: " + err.Error())
		if err := db.Rollback().Error; err != nil {
			return err
		}

		return err
	}

	// db commit
	if err := db.Commit().Error; err != nil {
		return err
	}

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
		return err
	}

	return nil
}

// UserList ユーザ一覧取得
func (d *UserDB) UserList() (Users, error) {
	users := Users{}

	if err := d.DB.Preload("Folders").Preload("Memos").Find(&users).Error; err != nil {
		log.Println("error: " + err.Error())
		return users, err
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
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

//var IsAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`) // 英文字3~15
//var IsNumberic = regexp.MustCompile(`^[0-9]*$`)
//var IsDate = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2} \d{2}:\d{2}:\d{2}$`) // yyyy-mm-dd hh:mm:ss
//var ValidPassword = regexp.MustCompile(`^[a-zA-Z0-9]*$`)                     // パスワード
//var IsUUID = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
//
//// 入力ありの場合はチェックする
//func IsValidPassword(password string) bool {
//	return password != "" && ValidPassword.MatchString(password)
//}
//
//func IsValidUserName(userName string) bool {
//	return userName != "" && IsAlphanumeric.MatchString(userName)
//}
//
//func IsValidUUID(u string) bool {
//	_, err := uuid.FromString(u)
//	return err == nil
//}
//
//func (u *User) UserValidation(user User, mode int) error {
//	if !IsValidUserName(user.UserName) {
//		return common.ErrInvalidUserName
//	}
//
//	// 作成時チェック
//	if mode == 0 {
//		// パスワードチェック
//		if !IsValidPassword(user.Password) {
//			return common.ErrInvalidPassword
//		}
//	}
//
//	// 更新時チェック
//	if mode == 1 {
//		// 入力があればチェックする
//		if user.Password != "" {
//			if !IsValidPassword(user.Password) {
//				return common.ErrInvalidPassword
//			}
//		}
//	}
//
//	return nil
//}
//
