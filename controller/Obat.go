package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/hakushigo/pa_c_obat/model"
	"strconv"
)

func AddObat(ctx fiber.Ctx) error {
	var data struct {
		KategoriObat []int      `json:"kategori_obat"`
		DataObat     model.Obat `json:"data_obat"`
	}

	err := json.Unmarshal(ctx.Body(), &data)

	// object model for obats
	obat := data.DataObat

	// array of object model of kategoriObat
	var KategoriObats []model.KategoriObat

	// fetch them
	for _, v := range data.KategoriObat {
		var tempKategoriObats model.KategoriObat
		find := db.Find(&tempKategoriObats, v)

		if find.RowsAffected == 0 {
			return ctx.Status(fiber.StatusRequestedRangeNotSatisfiable).JSON(fiber.Map{
				"status":  fiber.StatusRequestedRangeNotSatisfiable,
				"message": "the kategori with id " + strconv.Itoa(v) + " you find can't be found",
			})
		}

		KategoriObats = append(KategoriObats, tempKategoriObats)
	}

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	obat.KategoriObat = KategoriObats
	db.Create(&obat) // save obat data

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   obat,
	})
}

func ListObat(ctx fiber.Ctx) error {
	var daftarObat []model.Obat
	db.Preload("KategoriObat").Find(&daftarObat)

	return ctx.Status(200).JSON(daftarObat)
}

func GetObat(ctx fiber.Ctx) error {

	var dataObat []model.Obat

	id, _ := strconv.Atoi(ctx.Params("id"))
	result := db.Preload("KategoriObat").Find(&dataObat, id)
	if result.RowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"status": 404,
		})
	}

	return ctx.Status(200).JSON(dataObat)
}

func UpdateObat(ctx fiber.Ctx) error {
	var data model.Obat

	err := json.Unmarshal(ctx.Body(), &data)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	status := db.Find(&model.Obat{}, id).Updates(&data)

	if status.Error != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  status.Error,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":        200,
		"rows_affected": status.RowsAffected,
	})
}

func DeleteObat(ctx fiber.Ctx) error {
	var obat model.Obat

	// first fetch the id
	id, _ := strconv.Atoi(ctx.Params("id"))

	// fetch the Obat object
	err := db.First(&obat, id).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// delete the many-to-many association with KategoriObat
	err = db.Model(&obat).Association("KategoriObat").Clear()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// delete the Obat
	status := db.Delete(&obat)
	if status.Error != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":        200,
		"rows_affected": status.RowsAffected,
	})
}
