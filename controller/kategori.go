package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/hakushigo/pa_c_obat/model"
	"strconv"
)

func AddKategori(ctx fiber.Ctx) error {
	reqbody := ctx.Body()

	var input model.KategoriObat

	err := json.Unmarshal(reqbody, &input)
	if err != nil {
		return err
	}

	result := db.Create(&input)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":   200,
		"inserted": input,
	})
}

func ListKategori(ctx fiber.Ctx) error {
	var QueryResult []model.KategoriObat

	db.Preload("Obat").Find(&QueryResult)

	return ctx.Status(200).JSON(QueryResult)
}

func GetKategori(ctx fiber.Ctx) error {
	var FoundKategori model.KategoriObat
	id, _ := strconv.Atoi(ctx.Params("id"))

	search := db.Preload("Obat").First(&FoundKategori, id)

	if search.RowsAffected <= 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"status": 404,
		})
	}

	return ctx.Status(200).JSON(FoundKategori)
}

func UpdateKategori(ctx fiber.Ctx) error {
	var update model.KategoriObat           // where the data would be put
	id, _ := strconv.Atoi(ctx.Params("id")) // take the id from the param

	ttfrtu := db.Find(&update, id) // find the corresponding row by the id; ttfrtu : Try To Find Row To Update
	if ttfrtu.RowsAffected <= 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"status": 404,
		})
	}

	dec := json.NewDecoder(bytes.NewReader(ctx.Body()))
	err := dec.Decode(&update)

	if err != nil {
		return err
	}

	store := db.Save(&update)
	if store.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  store.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   update,
	})
}

func DeleteKategori(ctx fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var DeletedKategori model.KategoriObat
	db.Find(&DeletedKategori, id)

	err := db.Model(&DeletedKategori).Association("Obat").Clear()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	result := db.Delete(&DeletedKategori)
	if result.Error != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
	})
}
