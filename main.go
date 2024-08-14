package main

import "github.com/gofiber/fiber/v2"
import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
  	port := os.Getenv("PORT")
        if port == "" {  
                log.Printf("defaulting to port %s", port)
		app.Listen(":"+port)
        }else{	
		app.Listen(":8080")
	}
	
}
