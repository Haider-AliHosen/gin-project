package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// var (
// 	videoService    service.VideoService       = service.New()
// 	videoController controller.VideoController = controller.New(videoService)
// )

func main() {
	fmt.Println("Hello Gin, Welcome!")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok!!",
		})
	})
	// r.GET("/videos", func(c *gin.Context) {
	// 	c.JSON(200, videoController.FindAll())
	// })
	// r.POST("/videos", func(c *gin.Context) {
	// 	c.JSON(200, videoController.Save(c))
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
