//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type OrderDetailItem struct {
	OrderDetailItemId        string `json:"orderDetailItemId"`
	InvoiceLineId     		 string `json:"invoiceLineId"`
	ItemId   				 string `json:"itemId"`
}

var OrderDetailItems []OrderDetailItem

func readDataOrderDetailItem(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Read()
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	OrderDetailItems = []OrderDetailItem{}

	for _, record := range records {
		orderDetailItem := OrderDetailItem{
			OrderDetailItemId:        record[0],
			InvoiceLineId:     record[1],
			ItemId:   record[2]}
		OrderDetailItems = append(OrderDetailItems, orderDetailItem)
	}
	file.Close()
}

func writeDataOrderDetailItem(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, orderDetailItem := range OrderDetailItems {
		record := []string{orderDetailItem.OrderDetailItemId, orderDetailItem.InvoiceLineId, orderDetailItem.ItemId}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
