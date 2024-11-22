package main

import (
	"log"
	"os"

	"github.com/fbdaf/go-fiber-postgres/models"
	"github.com/fbdaf/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := storage.NewConnection(&storage.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		SSLMode:    os.Getenv("SSL_MODE"),
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		panic("failed to migrate books")
	}

	r := Repository{DB: db}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}

type Book struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"not null;unique" json:"title"`
	Author    string `gorm:"not null" json:"author"`
	Publisher string `gorm:"not null" json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/book/:id", r.GetBook)
	api.Get("/books", r.GetBooks)
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	err := r.DB.Create(&book).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book created successfully"})
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	book := Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	err := r.DB.Delete(&book).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}

func (r *Repository) GetBook(c *fiber.Ctx) error {
	book := Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	err := r.DB.Where("id = ?", book.ID).First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book fetched successfully", "data": book})
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	var books []Book
	err := r.DB.Find(&books).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Books fetched successfully", "data": books})
}


