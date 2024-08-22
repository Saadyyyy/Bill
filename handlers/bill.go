package handlers

import (
	"errors"
	"net/http"
	"print-bill/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func CreateBill(c echo.Context) error {
	var bill models.Bill
	if err := c.Bind(&bill); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.Create(&bill).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Preload related models
	var result models.Bill
	if err := db.Preload("Menu").Preload("Meja").First(&result, bill.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetBillByID(c echo.Context) error {
	id := c.Param("id")
	var bill models.Bill

	if err := db.Preload("Menu").Preload("Meja").First(&bill, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Bill not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, bill)
}

func GetMenuList(c echo.Context) error {
	var menus []models.Menu
	if err := db.Find(&menus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, menus)
}

func GetMeJaList(c echo.Context) error {
	var meja []models.Meja
	if err := db.Find(&meja).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, meja)
}

func CreateMenu(c echo.Context) error {
	var menu models.Menu
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if err := db.Create(&menu).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create menu"})
	}
	return c.JSON(http.StatusCreated, menu)
}

func CreateMeja(c echo.Context) error {
	var meja models.Meja
	if err := c.Bind(&meja); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := db.Create(&meja).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, meja)
}
func GetMenuByID(c echo.Context) error {
	id := c.Param("id")
	var menu models.Menu
	if err := db.First(&menu, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Menu not found")
	}
	return c.JSON(http.StatusOK, menu)
}
