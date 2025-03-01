package main

import (
	"log/slog"

	postgres_ "github.com/1SergPav1/notes_api/internal/adapter/postgres"
	"github.com/1SergPav1/notes_api/internal/entity"
	"github.com/1SergPav1/notes_api/internal/handlers"
	"github.com/1SergPav1/notes_api/internal/middleware"
	"github.com/1SergPav1/notes_api/internal/service"
	"github.com/1SergPav1/notes_api/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	utils.InitLogger()
	log := utils.Log

	dsn := "host=localhost user=admin password=12345678 dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Error("Ошибка подключения к БД", slog.String("error", err.Error()))
	}

	db.AutoMigrate(entity.User{}, entity.Note{})

	userRepo := postgres_.NewUserRepo(db)
	noteRepo := postgres_.NewNoteRepo(db)

	authService := service.NewAuthService(userRepo)
	noteService := service.NewNoteService(noteRepo)

	authHandler := handlers.NewAuthHandler(authService)
	noteHandler := handlers.NewNoteHandler(noteService)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddlware(log))

	authRouts := router.Group("/auth")
	{
		authRouts.POST("/register", authHandler.Register)
		authRouts.POST("/login", authHandler.Login)
	}

	noteRouts := router.Group("/notes").Use(middleware.AuthMiddleware())
	{
		noteRouts.POST("/", noteHandler.CreateNote)
		noteRouts.GET("/", noteHandler.GetNotes)
		noteRouts.PUT("/:id", noteHandler.UpdateNote)
		noteRouts.DELETE("/:id", noteHandler.DeleteNote)
	}

	log.Info("!! Сервер запущен на порту 8085")
	router.Run(":8085")
}
