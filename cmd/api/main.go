package main

import (
	"log"

	"project-tracker-backend/internal/app"
	"project-tracker-backend/internal/cache"
	projectHandler "project-tracker-backend/internal/handlers/web/project"
	projectRepo "project-tracker-backend/internal/repos/project"
	webRoutes "project-tracker-backend/internal/routes/web"
	projectService "project-tracker-backend/internal/services/web/project"

	"github.com/gofiber/fiber/v2/middleware/cors"

	memberHandler "project-tracker-backend/internal/handlers/web/member"
	memberRepo "project-tracker-backend/internal/repos/member"
	memberService "project-tracker-backend/internal/services/web/member"

	taskHandler "project-tracker-backend/internal/handlers/web/task"
	taskRepo "project-tracker-backend/internal/repos/task"
	taskService "project-tracker-backend/internal/services/web/task"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := app.ConnectDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	redisClient, err := app.ConnectRedis()
	if err != nil {
		log.Fatal("Redis connection failed: ", err)
	}
	defer redisClient.Close()

	log.Println("Redis connected successfully")

	redisCache := cache.NewRedisCache(redisClient)

	projectRepository := projectRepo.NewRepository(db)
	memberRepository := memberRepo.NewRepository(db)
	taskRepository := taskRepo.NewRepository(db)

	projectSvc := projectService.NewService(
		projectRepository,
		taskRepository,
	)

	projectSvc.SetCache(
		redisCache,
	)

	projectH := projectHandler.NewHandler(
		projectSvc,
	)

	server := fiber.New()
	server.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
	}))
	memberSvc := memberService.NewService(
		memberRepository,
	)

	memberH := memberHandler.NewHandler(
		memberSvc,
	)

	taskSvc := taskService.NewService(
		taskRepository,
	)

	taskSvc.SetCache(
		redisCache,
	)

	taskH := taskHandler.NewHandler(
		taskSvc,
	)

	webRoutes.RegisterHealthRoute(server)

	webRoutes.RegisterProjectRoutes(
		server,
		projectH,
	)

	webRoutes.RegisterMemberRoutes(
		server,
		memberH,
	)

	webRoutes.RegisterTaskRoutes(
		server,
		taskH,
	)

	log.Println("Project Tracker API running on http://localhost:8080")

	log.Fatal(server.Listen(":8080"))

}
