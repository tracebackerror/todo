package main

import (
	_"github.com/gin-gonic/gin"
)

func main() {
	r := setupRest()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
