package model

import (
	"testing"

	"zmemo/api/common"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// テスト準備
func initTest() {
	// get config
	config := new(common.Config)
	err := configor.Load(config, "./config.yaml")
	if err != nil {
		panic(err)
	}

	// connect db
	// user:password@tcp(localhost:3306)/dbname
	db, err = gorm.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.Host+":"+config.DB.Port+")/"+config.DB.Name+"?"+"charset=utf8mb4,utf8")

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
}

func TemplateMemoStruct() Memo {
	return Memo{
		UserName: "kanehira",
		Title:    "aaaa",
		Text:     "aaaa",
	}
}

// TestCreateMemo メモ作成テスト
func TestCreateMemo(t *testing.T) {
	initTest()
	newMemo := TemplateMemoStruct()
	db := MemoDB{
		DB: db,
	}
	_, err := db.CreateMemo(newMemo)

	if err != nil {
		t.Errorf("error:%s", err.Error())
	}
}
