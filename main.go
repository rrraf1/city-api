package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"
    middleware_kota_api "rafir.com/kota-api/middleware"
    routes_kota_api "rafir.com/kota-api/routes"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Printf("Warning: Error loading .env file: %v", err)
    }

    app := fiber.New()

    app.Use(cors.New(cors.Config{
        AllowCredentials: false,
        AllowOrigins:    "*",
        AllowHeaders:    "Origin, Content-Type, Accept",
    }))
    
    app.Use(middleware_kota_api.RecoveryMiddleware)

    r := routes_kota_api.NewRepository(nil)
    r.SetupRoutes(app)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Printf("Server listening on port %s", port)
    log.Fatal(app.Listen(":" + port))
}