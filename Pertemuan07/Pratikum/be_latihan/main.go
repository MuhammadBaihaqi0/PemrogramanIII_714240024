<<<<<<< HEAD
package main

import (
	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Inisialisasi database
	config.InitDB()

	//automigrate
	config.GetDB().AutoMigrate(&model.Mahasiswa{})
	router.SetupRoutes(app)

	app.Listen(":3000")
=======
package main

import (
	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Inisialisasi database
	config.InitDB()

	//automigrate
	config.GetDB().AutoMigrate(&model.Mahasiswa{})
	router.SetupRoutes(app)

	app.Listen(":3000")
>>>>>>> 018734980e81a8c102a73e54c523cc263a853aec
}