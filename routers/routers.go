package routers

import (
	"end-assignment/autenthikasi"
	"end-assignment/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	users := map[string]string{
		"admin":   "password",
		"manager": "manager",
		"staff":   "staff",
	}
	router := gin.Default()
	router.GET("/users", controllers.GetUsers)
	router.Use(autenthikasi.AuthMiddleware(users))
	router.POST("/users", controllers.InsertUsers)
	router.PUT("/users/:user_id", controllers.UpdateUsers)
	router.DELETE("/users/:user_id", controllers.DeleteUsers)

	router.GET("/maintask", controllers.GetMainTask)
	router.Use(autenthikasi.AuthMiddleware(users))
	router.POST("/maintask", controllers.InsertMainTask)
	router.PUT("/maintask/:taskId", controllers.UpdateMainTask)
	router.DELETE("/maintask/:taskId", controllers.DeleteMainTask)

	router.GET("/subtask", controllers.GetSubTask)
	router.Use(autenthikasi.AuthMiddleware(users))
	router.POST("/subtask", controllers.InsertSubTask)
	router.PUT("/subtask/:subtask_id", controllers.UpdateSubTask)
	router.DELETE("/subtask/:subtask_id", controllers.DeleteSubTask)

	router.GET("/assngmnt", controllers.GetAssgnmnt)
	router.Use(autenthikasi.AuthMiddleware(users))
	router.POST("/assngmnt", controllers.InsertSubTask)
	router.PUT("/assngmnt/:assignment_id ", controllers.UpdateSubTask)
	router.DELETE("/assngmnt/:assignment_id ", controllers.DeleteSubTask)
	return router
}
