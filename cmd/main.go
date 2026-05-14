package main

import (
	_ "github.com/tech160/sbb-backend/docs"
	"github.com/tech160/sbb-backend/internal/router"
)

// @title           SBB Backend API
// @version         1.0
// @description     SBB Backend Service — health check & demo endpoints
// @contact.name    SBB Tech
// @contact.email   r.nuttapon@steelbestbuy.com
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	r := router.New()
	r.Run(":8080")
}
