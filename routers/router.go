package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctrs "github.com/zachrey/blog/controllers"
	"github.com/zachrey/blog/models"
)

// LoadRouters 初始化router
func LoadRouters(router *gin.Engine) {
	loadRouters(router)
}

func loadRouters(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		post := models.GetPostByID(1)
		log.Println("@@ post", post)
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data":   post,
		})
	})

	router.POST("/upload", ctrs.UpLoadFile)

}