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
	Echo *echo.Echo
}

// New new server
func New(port string, db *gorm.DB, e *echo.Echo) *Server {
	return &Server{
		Port: port,
		DB:   db,
		Echo: e,
	}
}

// Start サーバ起動
func (s *Server) Start() {
	e := s.Echo

	s.InitHandler()

	// Start server
	e.Logger.Fatal(e.Start(":" + s.Port))
}

// InitHandler add handler
func (s *Server) InitHandler() {
	e := s.Echo

	e.Static("/", "../public")

	// ハンドラー初期化
	m := handler.NewMemoHandler(s.DB)
	u := handler.NewUserHandler(s.DB)
	f := handler.NewFolderHandler(s.DB)

	// ルート登録
	// ユーザAPI
	e.POST("/users", u.CreateUser())
	e.GET("/users", u.UserList())
	e.GET("/users/:userID", u.GetUser())
	e.PUT("/users/:userID", u.UpdateUser())
	e.PUT("/users/:userID/password", u.UpdatePassword())
	e.DELETE("/users/:userID", u.DeleteUser())
	e.POST("/users/login", u.Login())

	// メモAPI
	e.POST("/memos", m.CreateMemo())
	e.GET("/memos/:userID/:memoID", m.GetMemo())
	e.GET("/memos/:userID", m.MemoList())
	e.PUT("/memos/:userID/:memoID", m.UpdateMemo())
	e.DELETE("/memos/:userID/:memoID", m.DeleteMemo())

	// フォルダAPI
	e.POST("/folders", f.CreateFolder())
	e.GET("/folders/:userID/:folderID", f.GetFolder())
	e.GET("/folders/:userID", f.FolderList())
	e.PUT("/folders/:userID/:folderID", f.UpdateFolder())
	e.DELETE("/folders/:userID/:folderID", f.DeleteFolder())

}
