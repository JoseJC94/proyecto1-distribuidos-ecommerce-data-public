package main

/*Requests*/

import (
	"encoding/json"
	"net/http"
	"path"
)

func handleGetCustomer(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findCustomer(id)
	if i == -1 {
		dataJson, _ := json.Marshal(Customers[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	dataJson, err := json.Marshal(Customers[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostCustomer(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	customer := Customer{}
	json.Unmarshal(body, &customer)
	Customers = append(Customers, customer)
	w.WriteHeader(200)
	return
}

func handlePutCustomer(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findCustomer(id)
	var changed = false
	if i == -1 {
		return
	} else {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		reqCustomer := Customer{}
		json.Unmarshal(body, &reqCustomer)
		if reqCustomer.CustomerId != "" { Customers[i].CustomerId = reqCustomer.CustomerId; changed=true}
		if reqCustomer.Customername != "" { Customers[i].Customername = reqCustomer.Customername; changed=true}
		if reqCustomer.Country != "" { Customers[i].Country = reqCustomer.Country; changed=true}
	}
	if changed==true {
		w.WriteHeader(200)
	}
	return
}

func handleDeleteCustomer(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findCustomer(id)
	if i == -1 {
		return
	} else {
		deleteCustomer(id)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findCustomer(x string) int {
	for i, customer := range Customers {
		if x == customer.CustomerId {
			return i
		}
	}
	return -1
}

func deleteCustomer(id string){
	i := findCustomer(id)
	Customers = append(Customers[:i], Customers[i+1:]...)
}