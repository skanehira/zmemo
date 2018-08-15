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
		model := model.FolderDB{DB: s.DB}
		folder, err := model.CreateFolder(folder)
		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *folderHandler) GetFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")
		folderID := c.Param("folderID")

		model := model.FolderDB{DB: s.DB}
		folder, err := model.GetFolder(userID, folderID)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, folder)

	}
}

func (s *folderHandler) FolderList() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")

		model := model.FolderDB{DB: s.DB}
		folders, err := model.FolderList(userID)

		if err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.JSON(http.StatusOK, folders)
	}
}

func (s *folderHandler) UpdateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folder := model.Folder{}

		if err := c.Bind(&folder); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData.Error()))
		}

		folder.ID = c.Param("folderID")
		folder.UserID = c.Param("userID")

		model := model.FolderDB{DB: s.DB}

		if err := model.UpdateFolder(folder); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *folderHandler) DeleteFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folderID := c.Param("folderID")
		userID := c.Param("userID")

		model := model.FolderDB{DB: s.DB}
		if err := model.DeleteFolder(userID, folderID); err != nil {
			log.Println("error: " + err.Error())
			return c.JSON(common.GetErrorCode(err), common.NewError(err.Error()))
		}

		return c.NoContent(http.StatusOK)
	}
}
