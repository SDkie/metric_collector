package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	gin.SetMode(os.Getenv("MODE"))
	router := getRouter()
	router.Run(":" + os.Getenv("PORT"))
}
