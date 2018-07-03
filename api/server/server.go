package server

import (
	"zmemo/api/handler"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Server サーバー構造体
type Server struct {
	Port string
	DB   *gorm.DB
}

// Start サーバ起動
func (s *Server) Start() {
	e := echo.New()

	e.Static("/", "public")

	// ハンドラー初期化
	m := handler.NewMemoHandler(s.DB)
	u := handler.NewUserHandler(s.DB)
	f := handler.NewFolderHandler(s.DB)

	// ルート登録
	// ユーザAPI
	e.POST("/users", u.CreateUser())
	e.GET("/users", u.GetUsers())
	e.GET("/users/:userName", u.GetUser())
	e.PUT("/users/:userName", u.UpdateUser())
	e.PUT("/users/:userName/password", u.UpdatePassword())
	e.DELETE("/users/:userName", u.DeleteUser())
	e.POST("/users/login", u.Login())

	// メモAPI
	e.POST("/memos", m.CreateMemo())
	e.GET("/memos/:userName/:memoId", m.GetMemo())
	e.GET("/memos/:userName", m.GetMemos())
	//	e.PUT("/memos/:userName/:foldername", m.AddMemoToFolder())
	e.DELETE("/memos/:userName/:memoId", m.DeleteMemo())

	// フォルダAPI
	e.POST("/folders", f.CreateFolder())
	e.GET("/folders/:userName/:folderName", f.GetFolder())
	e.GET("/folders/:userName", f.GetFolders())
	e.PUT("/folders/:userName/:folderName", f.UpdateFolderName())
	e.DELETE("/folders/:userName/:folderName", f.DeleteFolder())

	// e.HTTPErrorHandler = middleware.customHTTPErrorHandler

	// Start server
	e.Logger.Fatal(e.Start(":" + s.Port))
}
