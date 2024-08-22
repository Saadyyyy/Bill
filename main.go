package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"print-bill/handlers"
	"print-bill/models"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=oke dbname=billiard port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the Bill model
	if err := db.AutoMigrate(&models.Meja{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	if err := db.AutoMigrate(&models.Menu{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	if err := db.AutoMigrate(&models.Bill{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

func (tr *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return tr.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	initDB()

	// Set the db variable in handlers package
	handlers.SetDB(db)

	// Define routes
	e.POST("/bill", handlers.CreateBill)
	e.GET("/bill/create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "create_bill.html", nil)
	})
	e.GET("/bill/:id", handlers.GetBillByID)

	e.GET("/menu", handlers.GetMenuList)
	e.GET("/meja", handlers.GetMeJaList)

	// Set up template rendering
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.POST("/menu", handlers.CreateMenu)
	e.GET("/menu/create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "menu.html", nil)
	})

	e.POST("/meja", handlers.CreateMeja)
	e.GET("/meja/create", func(c echo.Context) error {
		return c.Render(http.StatusOK, "meja.html", nil)
	})
	e.GET("/menu/:id", handlers.GetMenuByID)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
