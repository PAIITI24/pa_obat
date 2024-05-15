package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hakushigo/pa_c_obat/controller"
	"github.com/hakushigo/pa_c_obat/helper"
)

func main() {

	// Migrate table
	helper.Migrator()

	// declare server
	server := fiber.New(
		fiber.Config{
			Immutable: false,
			AppName:   "Produk_Apotek_APP",
		})

	// kategori obat
	server.Post("/obat/kategori", controller.AddKategori)
	server.Get("/obat/kategori", controller.ListKategori)
	server.Get("/obat/kategori/:id", controller.GetKategori)
	server.Put("/obat/kategori/:id", controller.UpdateKategori)
	server.Delete("/obat/kategori/:id", controller.DeleteKategori)

	// obat
	server.Post("/obat/", controller.AddObat)
	server.Get("/obat/", controller.ListObat)
	server.Get("/obat/:id", controller.GetObat)
	server.Put("/obat/:id", controller.UpdateObat)
	server.Delete("/obat/:id", controller.DeleteObat)

	server.Listen(":3001")
}
