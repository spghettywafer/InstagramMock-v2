package main

import (
	config "InstagramMock-v2/config"
	controller "InstagramMock-v2/controller"
	repo "InstagramMock-v2/repository"
	service "InstagramMock-v2/service"
	"fmt"

	// "net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	postRepository *repo.PostRepository       = repo.NewPostRepository(db)
	postService    *service.PostService       = service.NewPostService(postRepository)
	postController *controller.PostController = controller.NewPostController(postService)
)

func InitRoute() *gin.Engine {

	router := gin.Default()
	router.Static("/css", "./templates/css")
	router.Static("/upload", "./upload")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", postController.ShowAll)
	router.GET("/:id", postController.FindByID)
	router.POST("/create", postController.Create)
	router.GET("/create", postController.PageCreate)
	router.POST("/:id/edit", postController.Update)
	router.POST("/:id", postController.Delete)
	router.GET("/:id/edit", postController.PageUpdate)

	return router
}

func main() {
	defer config.CloseDatabaseConnection(db)
	fmt.Println("Instagram-Mock")

	router := InitRoute()

	router.Run()
}
