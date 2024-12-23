package main

import (
	"go_docker_api/config"

	"github.com/gin-gonic/gin"
)

var logger *config.Logger

func main() {

	logger = config.GetLogger("main")

	// Initialze configs	
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// := is a short variable declaration operator. It declares a variable and assigns a value to it.
	// The type of the variable is inferred from the value on the right-hand side.
	// only works inside functions
	server := gin.Default()

	// like C, Go uses the & operator to get the address of a variable and the * operator to dereference a pointer.
	// * is a pointer to a type. It is used to store the memory address of a value.
	// & is the address of operator. It returns the memory address of a value.
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")

}