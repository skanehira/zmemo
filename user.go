package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	UserName  string  `gorm:"primary_key;not null" json:"userName"`
	Password  string  `gorm:"not null" json:"password"`
	CreatedAt string  `gorm:"not null" json:"createAt"`
	UpdatedAt string  `gorm:"not null" json:"updateaAt"`
	DeletedAt *string `gorm:"null" json:"-"`
}

type Users []User

type UserDB struct {
	DB *gorm.DB
}

// ユーザ作成
func (d *UserDB) Create(user User) (User, error) {
	// 初期値
	user.CreatedAt = GetTime()
	user.UpdatedAt = GetTime()

	if err := d.DB.Create(&user).Find(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Println("error: " + NotFoundUser.Error())
			return user, NotFoundUser
		}
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
			log.Println("error: " + NotFoundUser.Error())
			return user, NotFoundUser
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
			return NotFoundUser
		}
		return db.Error
	}

	// フォルダ・メモ・ユーザを削除
	sql := "update users u, folders f, memos m set u.deleted_at = ?, f.deleted_at = ?, m.deleted_at = ? " +
		"where u.user_name = f.user_name and u.user_name = m.user_name and u.user_name = ?"

	time := GetTime()

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
			log.Println("error: " + NotFoundUser.Error())
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
			log.Println("error: " + NotFoundUser.Error())
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
			err = NotFoundUser
		}
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}
