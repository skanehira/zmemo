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

	// ユーザAPI
	e.GET("/users", s.GetUsers())
	e.GET("/users/:userId", s.GetUser())
	e.POST("/users", s.CreateUser())
	e.PUT("/users", s.UpdateUser())
	e.PUT("/users/password", s.UpdatePassword())
	e.DELETE("/users/:userId", s.DeleteUser())

	// メモAPI
	e.POST("/memos", s.CreateMemo())
	e.GET("/memos/:userId/:memoId", s.GetMemo())
	e.GET("/memos/:userId", s.GetMemos())
	e.DELETE("/memos/:userId/:memoId", s.DeleteMemo())

	// Start server
	e.Logger.Fatal(e.Start(":" + s.Port))
}
