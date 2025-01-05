package controller_kota_api

import (
	"encoding/csv"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Desa struct {
	ID   string
	Nama string
	Zip  string
}

type Kota struct {
	ID   string
	Nama string
	Type string
}

func GetNewsDetail(context *fiber.Ctx) error {
	selectNews := context.Params("id")
	if selectNews == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID parameter is required",
		})
	}

	decodedSearch, err := url.QueryUnescape(selectNews)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid search parameter",
		})
	}

	desaFile, err := os.Open("data/desakel.csv")
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Internal server error"})
	}
	defer desaFile.Close()

	kotaFile, err := os.Open("data/kabkota.csv")
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Internal server error"})
	}
	defer kotaFile.Close()

	desaReader := csv.NewReader(desaFile)
	desaRecords, err := desaReader.ReadAll()
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Internal server error"})
	}

	kotaReader := csv.NewReader(kotaFile)
	kotaRecords, err := kotaReader.ReadAll()
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Internal server error"})
	}

	var desaData []Desa
	for _, record := range desaRecords[1:] {
		desaData = append(desaData, Desa{
			ID:   record[0],
			Nama: record[1],
			Zip:  record[2],
		})
	}

	var kotaData []Kota
	for _, record := range kotaRecords[1:] {
		kotaData = append(kotaData, Kota{
			ID:   record[0],
			Nama: record[1],
			Type: record[2],
		})
	}

	cityVillagesMap := make(map[string][]Desa)
	for _, desa := range desaData {
		if strings.Contains(strings.ToLower(desa.Nama), strings.ToLower(decodedSearch)) {
			cityID := desa.ID[:5] 
			cityVillagesMap[cityID] = append(cityVillagesMap[cityID], desa)
		}
	}

	var responseData []map[string]interface{}
	for cityID, villages := range cityVillagesMap {
		var cityData Kota
		for _, kota := range kotaData {
			if kota.ID == cityID {
				cityData = kota
				break
			}
		}

		if cityData.ID != "" {
			cityResponse := map[string]interface{}{
				"kota":     cityData.Nama,
				"type":     cityData.Type,
				"villages": villages,
			}
			responseData = append(responseData, cityResponse)
		}
	}

	if len(responseData) > 0 {
		return context.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "News detail found",
			"data":    responseData,
		})
	} else {
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "No matching news found",
		})
	}
}
