package main

import (
	"os"
	"fmt"
//	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
//	"github.com/joho/godotenv"
)

func main(){	
	/* Loading Environment Variables*/
	/* Deactive to deploy, active only local work */
	// getEnvs()

	/* Setting the port */
	port := os.Getenv("PORT")

	fmt.Println("port is :")
	fmt.Println(port)
	
	/* Import Fiber */
	app := fiber.New()

	/* Adding middleware to avoid CORS */
	app.Use(cors.New())

	/* Serve static files as routes */
	app.Static("/", "./client/dist")

	/* Route when get petitions */
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Users from backend",
		})
	})

	/* Serving the app */
	app.Listen(":"+port)
	fmt.Println("Server listening on: " + port)
	
}

/*
// Deactive to deploy, active only local work
func getEnvs(){
	// Loading Environment Variables with godotenv 
	env := godotenv.Load(".env")
	
	// If error, log! 
	if env != nil {
		log.Fatalf("Error loading .env file")
	}
	
}
*/