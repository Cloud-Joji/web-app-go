package main

import (
	"context"
	"fmt"
//	"log"
	"os"

	"github.com/Cloud-Joji/web-app-go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
//	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	/* Loading Environment Variables*/
	/* Deactive to deploy, active only local work */
	// getEnvs()

	/* Setting the port */
	port := os.Getenv("PORT")

	if port == "" {
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
	fmt.Println("Database Connected")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	/* Adding middleware to avoid CORS */
	app.Use(cors.New())

	/* Serve static files as routes */
	app.Static("/", "./client/dist")

	/* Send data to database */
	app.Post("/api/certs", func(c *fiber.Ctx) error {
		var cert models.Cert

		c.BodyParser(&cert)

		coll := client.Database("certificados").Collection("Cursos")
		fmt.Println("---------------------")
		fmt.Println("New value added at: Name -- " + cert.Name)
		fmt.Println("New value added at: Platform -- " + cert.Platform)
		fmt.Println("---------------------")
		result, err := coll.InsertOne(context.TODO(), bson.D{
			{"name", cert.Name},
			{"platform", cert.Platform},
		})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	/* Get certs */
	app.Get("/api/certs", func(c *fiber.Ctx) error {
		var certs []models.Cert

		coll := client.Database("certificados").Collection("Cursos")
		results, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}

		for results.Next(context.TODO()) {
			var cert models.Cert
			results.Decode(&cert)
			certs = append(certs, cert)
		}

		return c.JSON(&fiber.Map{
			"certs": certs,
		})

	})

	/* Serving the app */
	app.Listen(":" + port)

}

/*
// Deactive to deploy, active only local work
func getEnvs() {
	// Loading Environment Variables with godotenv
	env := godotenv.Load("./client/.env")

	// If error, log!
	if env != nil {
		log.Fatalf("Error loading .env file")
	}

}
*/