package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"task_management_backend/controllers"
	"task_management_backend/dao"
	"task_management_backend/db"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	database, err := db.ConnectToDb()
	if err != nil {
		panic("failed to connect to database")
	}

	taskDAO := dao.NewTaskDAO(database)
	taskController := controllers.NewTaskController(taskDAO)

	router.POST("/tasks", taskController.CreateTask)
	router.GET("/tasks", taskController.GetAllTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	return router
}
