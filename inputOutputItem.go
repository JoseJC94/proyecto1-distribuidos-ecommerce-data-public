//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type Item struct {
	ItemId        string `json:"itemId"`
	Description     string `json:"description"`
	UnitPrice   string `json:"unitPrice"`
	PriceMultiplier string `json:"priceMultiplier"`
}

var Items []Item

func readDataItem(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	Items = []Item{}

	for _, record := range records {
		item := Item{
			ItemId:        record[0],
			Description:     record[1],
			UnitPrice:   record[2],
			PriceMultiplier: record[3]}
		Items = append(Items, item)
	}
	file.Close()
}

func writeDataItem(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, item := range Items {
		record := []string{item.ItemId, item.Description, item.UnitPrice,
			item.PriceMultiplier}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
