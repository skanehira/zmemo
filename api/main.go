package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"zmemo/api/config"
	"zmemo/api/logger"
	"zmemo/api/model"
	"zmemo/api/server"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	// get config
	config := config.New()

	// init logger
	logger.Init()

	// connect db
	// user:password@tcp(localhost:3306)/dbname?parseTime=true&charaset=utf8mb4,utf8
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(config.DBLog)

	// db migrate
	flag.Parse()
	if len(flag.Args()) > 0 {
		errors := []error{}

		if "create" == flag.Args()[0] {
			log.Print("info: start create tables...")

			if err := db.AutoMigrate(model.Memo{}).Error; err != nil {
				errors = append(errors, err)
			}
			if err := db.AutoMigrate(model.Folder{}).Error; err != nil {
				errors = append(errors, err)
			}
			if err := db.AutoMigrate(model.User{}).Error; err != nil {
				errors = append(errors, err)
			}

			// エラーがない場合
			if len(errors) < 1 {
				log.Printf("info: end create tables...")
			}

			// db.AutoMigrate(model.Memo{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
			// db.AutoMigrate(model.Folder{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
			// db.AutoMigrate(model.Users{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
		} else if "drop" == flag.Args()[0] {
			log.Print("info: start drop tables...")

			if err := db.DropTableIfExists(model.Memo{}).Error; err != nil {
				errors = append(errors, err)
			}
			if err := db.DropTableIfExists(model.Folder{}).Error; err != nil {
				errors = append(errors, err)
			}
			if err := db.DropTableIfExists(model.User{}).Error; err != nil {
				errors = append(errors, err)
			}

			// エラーがない場合
			if len(errors) < 1 {
				log.Printf("info: end drop tables...")
			}
		}

		for _, err := range errors {
			log.Printf("error: %s", err)
		}

		os.Exit(0)
	}

	// サーバ開始
	s := server.New(config.Port, db, echo.New())
	s.Start()

}
