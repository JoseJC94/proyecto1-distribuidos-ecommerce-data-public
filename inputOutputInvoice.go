//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type Invoice struct {
	InvoiceNo        string `json:"invoiceNo"`
	InvoiceDate     string `json:"invoicedate"`
}

var Invoices []Invoice

func readDataInvoice(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	Invoices = []Invoice{}

	for _, record := range records {
		invoice := Invoice{
			InvoiceNo:        record[0],
			InvoiceDate:     record[1],
		}
		Invoices = append(Invoices, invoice)
	}
	file.Close()
}

func writeDataInvoice(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, invoice := range Invoices {
		record := []string{invoice.InvoiceNo, invoice.InvoiceDate,
			}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
