package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) CreateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folder := Folder{}

		if err := c.Bind(&folder); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}

		// バリデーション
		if err := folderValidation(folder.UserName, folder.FolderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// フォルダ作成
		db := FolderDB{s.DB}
		folder, err := db.CreateFolder(folder)
		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *Server) GetFolder() echo.HandlerFunc {
	return func(c echo.Context) error {

		userName := c.Param("userName")
		folderName := c.Param("folderName")

		// バリデーション
		if err := folderValidation(userName, folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := FolderDB{s.DB}
		folder, err := db.GetFolder(userName, folderName)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *Server) GetFolders() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		// バリデーション
		if !isValidUserName(userName) {
			log.Println("error: " + InvalidUserName.Error())
			return c.JSON(GetErrorCode(InvalidUserName), Message{InvalidUserName.Error()})
		}

		db := FolderDB{s.DB}
		folders, err := db.GetFolders(userName)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, folders)
	}
}

func (s *Server) UpdateFolderName() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		folderName := c.Param("folderName")

		// バリデーション
		if err := folderValidation(userName, folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := FolderDB{s.DB}

		if err := db.UpdateFolderName(userName, folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *Server) DeleteFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		folderName := c.Param("folderName")

		// バリデーション
		if err := folderValidation(userName, folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := FolderDB{s.DB}
		if err := db.DeleteFolder(userName, folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}
