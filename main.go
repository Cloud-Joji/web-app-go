package main

import (
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main(){
	
	port := os.Getenv("PORT")
	app := fiber.New(type)
	app.Listen(port)
	fmt.Println("Server listening on: " + port)

}