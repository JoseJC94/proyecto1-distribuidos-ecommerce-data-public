//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type InvoiceLine struct {
	InvoiceLineId        string `json:"invoiceLineId"`
	Quantity     		 string `json:"quantity"`
	UnitPrice  			 string `json:"unitPrice"`
}

var InvoiceLines []InvoiceLine

func readDataInvoiceLines(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	InvoiceLines = []InvoiceLine{}

	for _, record := range records {
		invoiceLines := InvoiceLine{
			InvoiceLineId:  record[0],
			Quantity:     	record[1],
			UnitPrice:      record[2]}
		InvoiceLines = append(InvoiceLines, invoiceLines)
	}
	file.Close()
}

func writeDataInvoiceLines(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, invoiceLines := range InvoiceLines {
		record := []string{invoiceLines.InvoiceLineId, invoiceLines.Quantity, invoiceLines.UnitPrice}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
