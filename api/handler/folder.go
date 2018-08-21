package handler

import (
	"net/http"

	"zmemo/api/common"

	"zmemo/api/model"

	logger "github.com/Sirupsen/logrus"
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
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(err), common.NewError(common.ErrInvalidPostData))
		}

		// フォルダ作成
		model := model.FolderDB{DB: s.DB}
		folder, err := model.CreateFolder(folder)
		if err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		return c.JSON(http.StatusOK, folders)
	}
}

func (s *folderHandler) UpdateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folder := model.Folder{}

		if err := c.Bind(&folder); err != nil {
			logger.Error(common.Wrap(err))
			return c.JSON(common.GetErrorCode(common.ErrInvalidPostData), common.NewError(common.ErrInvalidPostData))
		}

		folder.ID = c.Param("folderID")
		folder.UserID = c.Param("userID")

		model := model.FolderDB{DB: s.DB}

		if err := model.UpdateFolder(folder); err != nil {
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
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
			logger.Error(err)
			return c.JSON(common.GetErrorCode(err), common.NewError(err))
		}

		return c.NoContent(http.StatusOK)
	}
}
