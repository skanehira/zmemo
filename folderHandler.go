package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) CreateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folder := Folder{}

		if err := c.Bind(&folder); err != nil {
			return c.JSON(GetErrorCode(InvalidPostData), Message{InvalidPostData.Error()})
		}

		// バリデーション
		if err := folderValidation(folder.UserID, folder.FolderName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		// フォルダ作成
		db := FolderDB{s.DB}
		folder, err := db.CreateFolder(folder)
		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *Server) GetFolder() echo.HandlerFunc {
	return func(c echo.Context) error {

		userId := c.Param("userId")
		folderName := c.Param("folderName")

		// バリデーション
		if err := folderValidation(userId, folderName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := FolderDB{s.DB}
		folder, err := db.GetFolder(userId, folderName)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *Server) GetFolders() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")

		// バリデーション
		if !isValidUserId(userId) {
			return c.JSON(GetErrorCode(InvalidUserID), Message{InvalidUserID.Error()})
		}

		db := FolderDB{s.DB}
		folders, err := db.GetFolders(userId)

		if err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.JSON(http.StatusOK, folders)
	}
}

func (s *Server) UpdateFolderName() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		folderName := c.Param("folderName")

		// バリデーション
		if err := folderValidation(userId, folderName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := FolderDB{s.DB}

		if err := db.UpdateFolderName(userId, folderName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *Server) DeleteFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("userId")
		folderName := c.QueryParam("folderName")

		// バリデーション
		if err := folderValidation(userId, folderName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		db := FolderDB{s.DB}
		if err := db.DeleteFolder(userId, folderName); err != nil {
			return c.JSON(GetErrorCode(err), Message{err.Error()})
		}

		return c.NoContent(http.StatusOK)
	}
}
