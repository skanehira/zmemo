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
	e.GET("/users/:userId", s.GetUser())
	e.PUT("/users/:userId", s.UpdateUser())
	e.PUT("/users/:userId/password", s.UpdatePassword())
	e.DELETE("/users/:userId", s.DeleteUser())

	// メモAPI
	e.POST("/memos", s.CreateMemo())
	e.GET("/memos/:userId/:memoid", s.GetMemo())
	e.GET("/memos/:userId", s.GetMemos())
	e.PUT("/memos/:userId/:foldername", s.AddMemoToFolder())
	e.DELETE("/memos/:userId/:memoid", s.DeleteMemo())

	// フォルダAPI
	e.POST("/folders", s.CreateFolder())
	e.GET("/folders/:userId/:folderName", s.GetFolder())
	e.GET("/folders/:userId", s.GetFolders())
	e.PUT("/folders/:userId/:folderName", s.UpdateFolderName())
	e.DELETE("/folders/:userId/:folderName", s.DeleteFolder())

	// Start server
	e.Logger.Fatal(e.Start(":" + s.Port))
}
