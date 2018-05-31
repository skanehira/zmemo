package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Server struct {
	Port string
	DB   *gorm.DB
}

func (s *Server) Start() {
	e := echo.New()

	e.Static("/", "public")

	// ユーザAPI
	e.POST("/users", s.CreateUser())
	e.GET("/users", s.GetUsers())
	e.GET("/users/:userName", s.GetUser())
	e.PUT("/users/:userName", s.UpdateUser())
	e.PUT("/users/:userName/password", s.UpdatePassword())
	e.DELETE("/users/:userName", s.DeleteUser())
	e.POST("/users/login", s.Login())

	// メモAPI
	e.POST("/memos", s.CreateMemo())
	e.GET("/memos/:userName/:memoId", s.GetMemo())
	e.GET("/memos/:userName", s.GetMemos())
	e.PUT("/memos/:userName/:foldername", s.AddMemoToFolder())
	e.DELETE("/memos/:userName/:memoId", s.DeleteMemo())

	// フォルダAPI
	e.POST("/folders", s.CreateFolder())
	e.GET("/folders/:userName/:folderName", s.GetFolder())
	e.GET("/folders/:userName", s.GetFolders())
	e.PUT("/folders/:userName/:folderName", s.UpdateFolderName())
	e.DELETE("/folders/:userName/:folderName", s.DeleteFolder())

	// Start server
	e.Logger.Fatal(e.Start(":" + s.Port))
}
