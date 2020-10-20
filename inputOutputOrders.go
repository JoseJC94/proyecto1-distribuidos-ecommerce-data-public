//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type Order struct {
	OrderId        string `json:"orderId"`
	InvoiceNo     string `json:"invoiceNo"`
	CustomerId   string `json:"customerId"`
}

var Orders []Order

func readDataOrder(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Read()
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	Orders = []Order{}

	for _, record := range records {
		order := Order{
			OrderId:        record[0],
			InvoiceNo:     record[1],
			CustomerId:   record[2]}
		Orders = append(Orders, order)
	}
	file.Close()
}

func writeDataOrder(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, order := range Orders {
		record := []string{order.OrderId, order.InvoiceNo, order.CustomerId}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
