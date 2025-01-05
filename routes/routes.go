package routes_kota_api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"gorm.io/gorm"
	controller_kota_api "rafir.com/kota-api/controller"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	apiLimiter := limiter.New(limiter.Config{
		Max:        15,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			token := c.Get("Authorization")
			if token != "" {
				return token
			}
			return c.IP()
		},
	})

	app.Get("/news/:id", apiLimiter, controller_kota_api.GetNewsDetail)
}
