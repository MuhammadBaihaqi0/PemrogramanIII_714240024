<<<<<<< HEAD
package main

import (
	"be_latihan/config"
	"be_latihan/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi database
	config.InitDB()

	//automigrate
	config.GetDB().AutoMigrate(&model.Mahasiswa{})

	app.Listen(":3000")
=======
package main

import (
	"be_latihan/config"
	"be_latihan/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi database
	config.InitDB()

	//automigrate
	config.GetDB().AutoMigrate(&model.Mahasiswa{})

	app.Listen(":3000")
>>>>>>> 018734980e81a8c102a73e54c523cc263a853aec
}