package auth

import "github.com/gofiber/fiber/v2"

type AuthRouter struct {
	handler *AuthHandler
}

func InitializeAuthModule() *AuthRouter {
	repo := &AuthRepository{}
	service := &AuthService{repository: repo}
	handler := &AuthHandler{service: service}
	return &AuthRouter{handler: handler}
}

func (router *AuthRouter) RegisterRoutes(r fiber.Router) {

	g := r.Group("/auth")
	g.Post("/register-user", router.handler.RegisterUser)
	g.Post("/login", router.handler.LoginUser)
	g.Post("/refresh-token", router.handler.RefreshToken)
}