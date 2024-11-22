package main

import (
	"example/web-service-gin/packages/controller"
	controller_auth "example/web-service-gin/packages/controller/auth"
	"example/web-service-gin/packages/database"
	env "example/web-service-gin/packages/env"
	"example/web-service-gin/packages/middleware"
	"example/web-service-gin/packages/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	env.LoadEnv()
	database.Connect()

	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Project{})
	database.Database.AutoMigrate(&model.Task{})
	database.Database.AutoMigrate(&model.Tag{})
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		currentTime := time.Now()
		ctx.String(http.StatusOK, "Hello World %v", currentTime)
	})
	router.POST("/register", controller_auth.Register)
	router.POST("/login", controller_auth.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/me", controller.Me)

	protectedRoutes.GET("/projects", controller.Projects)
	protectedRoutes.POST("/project", controller.CreateProject)

	protectedRoutes.POST("/tags", controller.CreateTag)
	protectedRoutes.POST("/tag", controller.CreateTag)

	protectedRoutes.GET("/tasks", controller.FindPaginateTask)
	protectedRoutes.POST("/task", controller.CreateTask)

	router.Run(":8080")
}
