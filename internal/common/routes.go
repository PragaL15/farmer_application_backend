package common

import (
	"farmerapp/internal/common/auth"
	"farmerapp/internal/common/realtime"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router) {
	authRouter := auth.InitializeAuthModule()
	authRouter.RegisterRoutes(r)

	realtimeRouter := realtime.InitializeRealtimeModule()
	realtimeRouter.RegisterRoutes(r)
}
