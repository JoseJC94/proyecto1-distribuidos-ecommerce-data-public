//Data IO

package main

import (
	"encoding/csv"
	"os"
)

type Customer struct {
	CustomerId        string `json:"customerid"`
	Customername     string `json:"customername"`
	Country   string `json:"country"`
}

var Customers []Customer

func readDataCustomer(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	Customers = []Customer{}

	for _, record := range records {
		customer := Customer{
			CustomerId:        record[0],
			Customername:     record[1],
			Country:   record[2],
		}
		Customers = append(Customers, customer)
	}
	file.Close()
}

func writeDataCustomer(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, customer := range Customers {
		record := []string{
			customer.CustomerId, customer.Customername, customer.Country,
		}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
