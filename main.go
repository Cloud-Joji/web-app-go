package main

import (
	"os"
	"fmt"
	"context"
//	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
//	"github.com/joho/godotenv"
	"github.com/Cloud-Joji/web-app-go/models"
)

func main(){	
	/* Loading Environment Variables*/
	/* Deactive to deploy, active only local work */
	// getEnvs()

	/* Setting the port */
	port := os.Getenv("PORT")
	
	if port == ""{
		port = "4000"
	}

	fmt.Println("Server listening on port: " + port)
	
	/* Import Fiber */
	app := fiber.New()

	/* Connecting to Database */
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	cluster := os.Getenv("DB_CLUSTER")

	uri := "mongodb+srv://" + username + ":" + password + "@" + cluster + "/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil{
		panic(err)
	}

	/* Adding middleware to avoid CORS */
	app.Use(cors.New())

	/* Serve static files as routes */
	app.Static("/", "./client/dist")

	/* Route when get petitions */
	app.Get("/certs", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Certifications from backend",
		})
	})

	app.Post("/certs", func(c *fiber.Ctx) error {
		var cert models.Cert

		c.BodyParser(&cert)

		coll := client.Database("go-cert-wapp").Collection("Platzi")
		coll.InsertOne(context.TODO(), bson.D{
			{"name", cert.Name},
		})

		return c.JSON(&fiber.Map{
			"data": "Adding certification...",
		})
	})

	/* Serving the app */
	app.Listen(":" + port)
	
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