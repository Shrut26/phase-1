package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Shrut26/postgresApi/models"
	"github.com/Shrut26/postgresApi/storage"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book has been added"})
	return nil
}

func (r *Repository) Getbooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Books fetched successfully",
			"data": bookModels})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Book{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Please provide valid id"})
		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete book"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book deleted successfully"})
	return nil
}

func (r *Repository) GetBookbyID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := models.Book{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "please provide valid ID"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(bookModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "unable to find book"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book found", "data": bookModel})
	return nil
}
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book.:id", r.GetBookbyID)
	api.Get("/books", r.Getbooks)
}

func main() {
	err := godotenv.Load("env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBname:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate database")
	}
	r := Repository{
		DB: db,
	}
	app := fiber.New(&fiber.Settings{StrictRouting: true})
	r.SetupRoutes(app)
	app.Listen(":8000")
}
