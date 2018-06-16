package main

import "github.com/gin-gonic/gin"
import "net/http"



type ToDoTask struct {
	Task     string `form:"task" json:"task" binding:"required"`
	Status string `form:"status" json:"status" binding:"required"`
}


func setupRest() *gin.Engine{
	r := gin.Default()
	r.GET("/rest/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Rest Service is Up & Running",
		})
	})
    

    taskRestRouter := r.Group("rest/task")
    {
        taskRestRouter.GET("/", func(c *gin.Context) {
            c.String(http.StatusOK, "API For All List: rest/task/")
        })
        taskRestRouter.GET("/:id", func(c *gin.Context) {
            id := c.Param("id")
            c.String(http.StatusOK, "API For Single Todo: rest/task/"+id)
        })
        taskRestRouter.POST("/", func(c *gin.Context) {
            var jsonData ToDoTask
            if err := c.ShouldBindJSON( &jsonData); err == nil {
                if jsonData.Task != "" && jsonData.Status != "" {
                    c.JSON(http.StatusOK, gin.H{"status": "Task logged in"})
                } else {
                    c.JSON(http.StatusUnauthorized, gin.H{"status": "Task failed to add in database"})
                }
            } else {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            }
            
        })
        taskRestRouter.PUT("/:id", func(c *gin.Context) {
            id := c.Param("id")
            c.String(http.StatusOK, "API For Update Single Todo: rest/task/"+id)
        })
        taskRestRouter.DELETE("/:id", func(c *gin.Context) {
            c.String(http.StatusOK, "API For DELETE Single Todo: rest/task/1")
        })
        
	}
    
   
    
    
	return r
}