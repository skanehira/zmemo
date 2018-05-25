package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// get config
	config := new(Config)
	if err := configor.Load(config, "./config.yaml"); err != nil {
		panic(err)
	}

	// connect db
	// user:password@tcp(localhost:3306)/dbname
	db, err := gorm.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.Host+":"+config.DB.Port+")/"+config.DB.Name)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	// db migrate
	flag.Parse()
	if len(flag.Args()) > 0 {
		if "migrate" == flag.Args()[0] {
			fmt.Println("db migreate…")
			if err := db.AutoMigrate(User{}).Error; err != nil {
				panic(err)
			}

			db.AutoMigrate(Memo{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
			db.AutoMigrate(Folder{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
		}
		os.Exit(0)
	}

	// サーバ開始
	server := Server{config.Port, db}
	server.Start()

}
