// Package test ut test utility
package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http/httptest"
	"zmemo/api/config"
	"zmemo/api/model"
	"zmemo/api/server"

	"github.com/comail/colog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// InitServer initialize echo server
func InitServer() *echo.Echo {
	// get config
	config := config.New("../config/config_test.yaml")

	// connect db
	// user:password@tcp(localhost:3306)/dbname?parseTime=true&charaset=utf8mb4,utf8
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	db.LogMode(config.DBLog)
	db.DropTable(model.Memo{})
	db.DropTable(model.Folder{})
	db.DropTable(model.Users{})

	db.AutoMigrate(model.Memo{})
	db.AutoMigrate(model.Folder{})
	db.AutoMigrate(model.Users{})

	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	e := echo.New()
	s := server.New(config.Port, db, e)

	s.InitHandler()

	return e
}

// POST http post
func POST(e *echo.Echo, path string, body []byte) (int, string) {
	req := httptest.NewRequest("POST", path, bytes.NewBuffer(body))
	req.Header.Set("Content-type", "application/json")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	return rec.Code, rec.Body.String()
}

// GET http get
func GET(e *echo.Echo, path string, params ...string) (int, string) {
	req := httptest.NewRequest("GET", path, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	return rec.Code, rec.Body.String()
}
