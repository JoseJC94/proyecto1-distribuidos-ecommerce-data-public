//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type OrderDetail struct {
	OrderDetailId        string `json:"orderDetailId"`
	InvoiceNo     string `json:"invoiceNo"`
	InvoiceLineId   string `json:"invoiceLineId"`
}

var OrderDetails []OrderDetail

func readDataOrderDetail(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Read()
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	OrderDetails = []OrderDetail{}

	for _, record := range records {
		orderDetail := OrderDetail{
			OrderDetailId:        record[0],
			InvoiceNo:     record[1],
			InvoiceLineId:   record[2]}
		OrderDetails = append(OrderDetails, orderDetail)
	}
	file.Close()
}

func writeDataOrderDetail(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, orderDetail := range OrderDetails {
		record := []string{orderDetail.OrderDetailId, orderDetail.InvoiceNo, orderDetail.InvoiceLineId}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
