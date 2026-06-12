// package main

// import ("fmt")

// func main() {
// 	fmt.Println("API Observatory Started")
// }

package main

import (
	"fmt"

	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

	_ = db

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Observatory",
		})
	})

	r.Run(":8080")

}
