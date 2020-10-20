package main

/*Requests*/

import (
	"encoding/json"
	"net/http"
	"path"
)

func handleGetInvoiceLines(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findInvoiceLines(id)
	if i == -1 {
		dataJson, _ := json.Marshal(InvoiceLines[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	dataJson, err := json.Marshal(InvoiceLines[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostInvoiceLines(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	invoiceLines := InvoiceLine{}
	json.Unmarshal(body, &invoiceLines)
	InvoiceLines = append(InvoiceLines, invoiceLines)
	w.WriteHeader(200)
	return
}

func handlePutInvoiceLines(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findInvoiceLines(id)
	var changed = false
	if i == -1 {
		return
	} else {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		reqInvoiceLines := InvoiceLine{}
		json.Unmarshal(body, &reqInvoiceLines)
		if reqInvoiceLines.InvoiceLineId != "" { InvoiceLines[i].InvoiceLineId = reqInvoiceLines.InvoiceLineId; changed=true}
		if reqInvoiceLines.Quantity != "" { InvoiceLines[i].Quantity = reqInvoiceLines.Quantity; changed=true}
		if reqInvoiceLines.UnitPrice    != "" { InvoiceLines[i].UnitPrice = reqInvoiceLines.UnitPrice; changed=true }
	}
	if changed==true {
		w.WriteHeader(200)
	}
	return
}

func handleDeleteInvoiceLines(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findInvoiceLines(id)
	if i == -1 {
		return
	} else {
		deleteInvoiceLines(id)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findInvoiceLines(x string) int {
	for i, invoiceLines := range InvoiceLines {
		if x == invoiceLines.InvoiceLineId {
			return i
		}
	}
	return -1
}

func deleteInvoiceLines(id string){
	i := findInvoiceLines(id)
	InvoiceLines = append(InvoiceLines[:i], InvoiceLines[i+1:]...)
}