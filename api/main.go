package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"zmemo/api/config"
	"zmemo/api/model"
	"zmemo/api/server"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/comail/colog"
)

func main() {

	// get config
	config := config.NewConfig()

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
		if "migrate" == flag.Args()[0] {
			fmt.Println("db migreate…")
			if err := db.AutoMigrate(user.User{}).Error; err != nil {
				panic(err)
			}
			db.AutoMigrate(model.Memo{})
			db.AutoMigrate(model.Folder{})
			db.AutoMigrate(model.Users{})
			// db.AutoMigrate(model.Memo{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
			// db.AutoMigrate(model.Folder{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
			// db.AutoMigrate(model.Users{}).AddForeignKey("user_name", "users(user_name)", "RESTRICT", "RESTRICT")
		}
		os.Exit(0)
	}

	// ログ設定
	// colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	// ログの設定方法
	// log.Printf("trace: this is a trace log.")
	// log.Printf("debug: this is a debug log.")
	// log.Printf("info: this is an info log.")
	// log.Printf("warn: this is a warning log.")
	// log.Printf("error: this is an error log.")
	// log.Printf("alert: this is an alert log.")
	// log.Printf("this is a default level log.")

	// サーバ開始
	server := server.Server{config.Port, db}
	server.Start()

}
