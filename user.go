package main

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	UserID    string  `gorm:"primary_key";json:"userId"`
	UserName  string  `gorm:"not null";json:"userName"`
	Password  string  `gorm:"not null";json:"password"`
	CreatedAt string  `gorm:"not null";json:"createAt"`
	UpdatedAt string  `gorm:"not null";json:"updateaAt"`
	DeletedAt *string `gorm:"null";json:"-"`
}

type Users []User

type UserDB struct {
	DB *gorm.DB
}

// ユーザ作成
func (d *UserDB) Create(user User) (User, error) {
	// ユーザID生成
	uuid, err := uuid.NewV4()
	if err != nil {
		return user, InvalidUserID
	}

	// 初期値
	user.UserID = uuid.String()
	user.CreatedAt = GetTime()
	user.UpdatedAt = GetTime()

	if err := d.DB.Create(&user).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// ユーザ更新
func (d *UserDB) Update(user User) (User, error) {
	// 更新日取得
	user.UpdatedAt = GetTime()

	// 構造体を使用すると値が0の場合はブランクで更新されてしまう
	// http://doc.gorm.io/crud.html#update
	// mapに変換してから更新することでそれを回避できる
	newUser := StructToMap(&user)

	// ユーザが存在しない場合 or 削除済みの場合はエラー
	db := d.DB.Find(&user)
	if err := db.Error; err != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return user, NotFoundUser
		}

		return user, db.Error
	}

	// updatesを使用するとcall_backでupdate_atが自動更新されるので
	// UpdateColumnsを使用することで回避できる
	if err := db.Model(user).UpdateColumns(newUser).Find(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

// ユーザ削除
func (d *UserDB) Delete(userId string) error {

	user := User{UserID: userId}

	// ユーザが存在しない or 削除済みの場合はエラー
	db := d.DB.Find(&user)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return NotFoundUser
		}
		return db.Error
	}

	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// ユーザ情報取得
func (d *UserDB) GetUser(userId string) (User, error) {

	user := User{UserID: userId}

	// ユーザが存在しない or 削除済みの場合はエラー
	if err := d.DB.Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return user, NotFoundUser
		}
		return user, err
	}

	return user, nil
}

// パスワード更新
func (d *UserDB) UpdatePassword(user User) error {
	// 更新日取得
	user.UpdatedAt = GetTime()
	newUser := StructToMap(&user)

	// ユーザが存在しない or 削除済みの場合はエラー
	db := d.DB.Find(&user)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return NotFoundUser
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
		return *users, err
	}

	return *users, nil
}
