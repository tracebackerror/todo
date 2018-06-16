package main

import "github.com/gin-gonic/gin"
import "net/http"

func setupRest() *gin.Engine{
	router := gin.Default()
	router.GET("/rest/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Rest Service is Up & Running",
		})
	})
    
    
    taskRestRouter := router.Group("rest/task")
	{
        taskRestRouter.GET("/", func(c *gin.Context) {
            c.String(http.StatusOK, "API For All List: rest/task/")
        })
        taskRestRouter.GET("/:id", func(c *gin.Context) {
            c.String(http.StatusOK, "API For Single Todo: rest/task/1")
        })
        taskRestRouter.POST("/", func(c *gin.Context) {
            c.String(http.StatusOK, "API For Create Single Todo: rest/task/")
        })
        taskRestRouter.PUT("/:id", func(c *gin.Context) {
            c.String(http.StatusOK, "API For Update Single Todo: rest/task/1")
        })
        taskRestRouter.DELETE("/:id", func(c *gin.Context) {
            c.String(http.StatusOK, "API For DELETE Single Todo: rest/task/1")
        })

     
	}
    
    
	return router
}