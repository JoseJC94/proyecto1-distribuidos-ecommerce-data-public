package main

/*Requests*/

import (
	"encoding/json"
	"net/http"
	"path"
)

func handleGetInvoice(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findInvoice(id)
	if i == -1 {
		dataJson, _ := json.Marshal(Invoices[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	dataJson, err := json.Marshal(Invoices[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostInvoice(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	invoice := Invoice{}
	json.Unmarshal(body, &invoice)
	Invoices = append(Invoices, invoice)
	w.WriteHeader(200)
	return
}

func handlePutInvoice(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findInvoice(id)
	var changed = false
	if i == -1 {
		return
	} else {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		reqInvoice := Invoice{}
		json.Unmarshal(body, &reqInvoice)
		if reqInvoice.InvoiceNo != "" { Invoices[i].InvoiceNo = reqInvoice.InvoiceNo; changed=true}
		if reqInvoice.InvoiceDate != "" { Invoices[i].InvoiceDate = reqInvoice.InvoiceDate; changed=true}
	}
	if changed==true {
		w.WriteHeader(200)
	}
	return
}

func handleDeleteInvoice(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findInvoice(id)
	if i == -1 {
		return
	} else {
		deleteInvoice(id)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findInvoice(x string) int {
	for i, invoice := range Invoices {
		if x == invoice.InvoiceNo {
			return i
		}
	}
	return -1
}

func deleteInvoice(id string){
	i := findInvoice(id)
	Invoices = append(Invoices[:i], Invoices[i+1:]...)
}