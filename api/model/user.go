package model

import (
	"log"
	"regexp"
	"time"

	"zmemo/api/common"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string     `gorm:"primary_key;not null" json:"id"`
	UserName  string     `gorm:"primary_key;not null" json:"userName"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `gorm:"null" json:"createAt"`
	UpdatedAt time.Time  `gorm:"null" json:"updateaAt"`
	DeletedAt *time.Time `gorm:"null" json:"-"`
}

type Users []User

type UserDB struct {
	DB *gorm.DB
}

// ユーザ作成
func (d *UserDB) Create(user User) (User, error) {
	// メモID生成
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("error: " + err.Error())
		return user, common.ErrInvalidMemoID
	}

	// 初期値
	user.ID = uuid.String()
	user.CreatedAt = common.GetTime()
	user.UpdatedAt = common.GetTime()

	if err := d.DB.Create(&user).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("error: " + common.ErrNotFoundUser.Error())
			return user, common.ErrNotFoundUser
		}
		return user, err
	}

	return user, nil
}

// ユーザ更新
func (d *UserDB) Update(user User) (User, error) {
	// 更新日取得
	user.UpdatedAt = common.GetTime()

	// 構造体を使用すると値が0の場合はブランクで更新されてしまう
	// http://doc.gorm.io/crud.html#update
	// mapに変換してから更新することでそれを回避できる
	newUser := common.StructToMap(&user)

	// ユーザが存在しない場合 or 削除済みの場合はエラー
	db := d.DB.Find(&user)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			log.Println("error: " + common.ErrNotFoundUser.Error())
			return user, common.ErrNotFoundUser
		}

		return user, db.Error
	}

	// updatesを使用するとcall_backでupdate_atが自動更新される
	// http://doc.gorm.io/crud.html#update
	// UpdateColumnsを使用することで回避できる
	if err := db.Model(user).UpdateColumns(newUser).Find(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

// ユーザ削除
func (d *UserDB) Delete(userName string) error {

	user := User{UserName: userName}

	// ユーザが存在しない or 削除済みの場合はエラー
	db := d.DB.Find(&user)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return common.ErrNotFoundUser
		}
		return db.Error
	}

	// フォルダ・メモ・ユーザを削除
	sql := "update users u, folders f, memos m set u.deleted_at = ?, f.deleted_at = ?, m.deleted_at = ? " +
		"where u.user_name = f.user_name and u.user_name = m.user_name and u.user_name = ?"

	time := common.GetTime()

	if err := db.Exec(sql, time, time, time, userName).Error; err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

// ユーザ情報取得
func (d *UserDB) GetUser(userName string) (User, error) {

	user := User{UserName: userName}

	// ユーザが存在しない or 削除済みの場合はエラー
	if err := d.DB.Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("error: " + common.ErrNotFoundUser.Error())
			return user, common.ErrNotFoundUser
		}
		return user, err
	}

	return user, nil
}

// パスワード更新
func (d *UserDB) UpdatePassword(user User) error {
	// 更新日取得
	user.UpdatedAt = common.GetTime()
	newUser := common.StructToMap(&user)

	// ユーザが存在しない or 削除済みの場合はエラー
	db := d.DB.Find(&user)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			log.Println("error: " + common.ErrNotFoundUser.Error())
			return common.ErrNotFoundUser
		}
		return db.Error
	}

	if err := db.Model(&user).UpdateColumns(newUser).Error; err != nil {
		return err
	}

	return nil
}

// ユーザ一覧取得
func (d *UserDB) GetUsers() (Users, error) {
	users := new(Users)

	if err := d.DB.Find(users).Error; err != nil {
		log.Println("error: " + err.Error())
		return *users, err
	}

	return *users, nil
}

// ログインユーザ情報取得
func (d *UserDB) GetLoginUser(user User) error {

	// ユーザIDとパスワード
	if err := d.DB.Model(&user).Where("user_name = ? and password = ?", user.UserName, user.Password).Scan(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = common.ErrNotFoundUser
		}
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

var IsAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]*$`) // 英文字3~15
var IsNumberic = regexp.MustCompile(`^[0-9]*$`)
var IsDate = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2} \d{2}:\d{2}:\d{2}$`) // yyyy-mm-dd hh:mm:ss
var ValidPassword = regexp.MustCompile(`^[a-zA-Z0-9]*$`)                     // パスワード
var IsUUID = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

// 入力ありの場合はチェックする
func IsValidPassword(password string) bool {
	return password != "" && ValidPassword.MatchString(password)
}

func IsValidUserName(userName string) bool {
	return userName != "" && IsAlphanumeric.MatchString(userName)
}

func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}

func (u *User) UserValidation(user User, mode int) error {
	if !IsValidUserName(user.UserName) {
		return common.ErrInvalidUserName
	}

	// 作成時チェック
	if mode == 0 {
		// パスワードチェック
		if !IsValidPassword(user.Password) {
			return common.ErrInvalidPassword
		}
	}

	// 更新時チェック
	if mode == 1 {
		// 入力があればチェックする
		if user.Password != "" {
			if !IsValidPassword(user.Password) {
				return common.ErrInvalidPassword
			}
		}
	}

	return nil
}
