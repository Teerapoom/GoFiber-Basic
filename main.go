package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	contorller "github.com/teerapoom/miniProjsct_Book/Contorller"
	middleware "github.com/teerapoom/miniProjsct_Book/Middleware"
	_ "github.com/teerapoom/miniProjsct_Book/docs" // load generated docs
	"github.com/teerapoom/miniProjsct_Book/model"
)

var Books []model.Book

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize standard Go html template engine
	engine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// User
	app.Post("/login", contorller.Login)

	contorller.SeedData()

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}))
	app.Use(middleware.CheckMiddleware) // ผ่าน Middleware ทุกครั้งs

	app.Get("/books", contorller.GetBooks)
	app.Get("/books/:id", contorller.GetBook)
	app.Post("/app/book", contorller.CreateBook)
	app.Put("/update/book/:id", contorller.UpdateBook)
	app.Delete("/remove/book/:id", contorller.DeleteBook)
	app.Post("/upload", contorller.UploadImage)
	// Setup route
	app.Get("/html", renderTemplate)
	// ENV
	app.Get("/env", getenv)
	log.Fatal(app.Listen(":8080"))
}

// ส่งค่าไปในไฟล์ PDF
func renderTemplate(c *fiber.Ctx) error {
	// Render the index with variable data index คือชื่อไฟล์ที่สร้าง
	return c.Render("index", fiber.Map{
		"Name":   "Hello World 🌏",
		"MyName": "Teerapoom",
	})
}

func getenv(c *fiber.Ctx) error {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "defaultSecret" // Default value if not specified
	}
	return c.JSON(
		fiber.Map{
			"SECRET_KEY": secretKey,
		})
}
