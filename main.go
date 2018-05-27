package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/comail/colog"
)

func main() {

	// get config
	config := new(Config)
	if err := configor.Load(config, "./config.yaml"); err != nil {
		panic(err)
	}

	// connect db
	// user:password@tcp(localhost:3306)/dbname
	db, err := gorm.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.Host+":"+config.DB.Port+")/"+config.DB.Name+"?"+"charset=utf8mb4,utf8")

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
	server := Server{config.Port, db}
	server.Start()

}
