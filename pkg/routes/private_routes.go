package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhatpankaj/go-fiber-rest-api/app/controllers"
	"github.com/prabhatpankaj/go-fiber-rest-api/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens

	route.Get("/roles", middleware.JWTProtected(), controllers.GetRoles) // get roles
}
