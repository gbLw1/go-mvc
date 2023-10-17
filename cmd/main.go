package main

import (
	"fmt"
	"go-mvc/pkg/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("hello world")
}
