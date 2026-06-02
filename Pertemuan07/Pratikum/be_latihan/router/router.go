<<<<<<< HEAD
package router

import (
	"be_latihan/handler"
	"be_latihan/model"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	mahasiswa := app.Group("/api/mahasiswa")
	mahasiswa.Get("/", handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", handler.InsertMahasiswa)	
	mahasiswa.Put("/:npm", handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", handler.DeleteMahasiswa)
=======
package router

import (
	"be_latihan/handler"
	"be_latihan/model"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	mahasiswa := app.Group("/api/mahasiswa")
	mahasiswa.Get("/", handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", handler.InsertMahasiswa)	
	mahasiswa.Put("/:npm", handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", handler.DeleteMahasiswa)
>>>>>>> 018734980e81a8c102a73e54c523cc263a853aec
}