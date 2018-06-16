package main

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"


type ToDoTask struct {
    gorm.Model
	Task     string `form:"task" json:"task" binding:"required"`
	Status string `form:"status" json:"status" binding:"required"`
}

func setupRest() *gin.Engine{
	r := gin.Default()
    db, err := gorm.Open("mysql", "tracebackerror:12345@/todo?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("failed to connect database")
    }
    //defer db.Close()
    db.AutoMigrate(&ToDoTask{})
    
    
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
            
            c.JSON(200, gin.H{
			"message": "Implementing",
            })
            c.String(http.StatusOK, "API For Single Todo: rest/task/"+id)
        })
        taskRestRouter.POST("/", func(c *gin.Context) {
            var jsonData ToDoTask
            if err := c.ShouldBindJSON( &jsonData); err == nil {
                //todotaskdata := ToDoTask{Task: string(jsonData.Task), Status: string(jsonData.Status)}
                
                db.Create(&jsonData)
                
                message := "Task Logged"
                c.JSON(http.StatusOK, gin.H{"message": message, ""})
                
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