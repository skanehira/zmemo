package handler

import (
	"log"
	"net/http"

	"zmemo/api/common"
	"zmemo/api/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type folderHandler struct {
	DB *gorm.DB
}

func NewFolderHandler(db *gorm.DB) *folderHandler {
	return &folderHandler{db}
}

func (s *folderHandler) CreateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folder := model.Folder{}

		if err := c.Bind(&folder); err != nil {
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
		}

		// フォルダ作成
		db := model.FolderDB{s.DB}
		folder, err := db.CreateFolder(folder)
		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *folderHandler) GetFolder() echo.HandlerFunc {
	return func(c echo.Context) error {

		folderName := c.Param("folderName")

		db := model.FolderDB{s.DB}
		folder, err := db.GetFolder(folderName)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *folderHandler) GetFolders() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")

		db := model.FolderDB{s.DB}
		folders, err := db.GetFolders(userName)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, folders)
	}
}

func (s *folderHandler) UpdateFolderName() echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.Param("userName")
		folderName := c.Param("folderName")

		db := model.FolderDB{s.DB}

		if err := db.UpdateFolderName(userName, folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *folderHandler) DeleteFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folderName := c.Param("folderName")

		db := model.FolderDB{s.DB}
		if err := db.DeleteFolder(folderName); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}
