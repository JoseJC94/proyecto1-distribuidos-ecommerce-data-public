package main

/*Requests*/

import (
	"encoding/json"
	"net/http"
	"path"
)

func handleGetItem(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findItem(id)
	if i == -1 {
		dataJson, _ := json.Marshal(Items[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	dataJson, err := json.Marshal(Items[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostItem(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	item := Item{}
	json.Unmarshal(body, &item)
	Items = append(Items, item)
	w.WriteHeader(200)
	return
}

func handlePutItem(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findItem(id)
	var changed = false
	if i == -1 {
		return
	} else {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		reqItem := Item{}
		json.Unmarshal(body, &reqItem)
		if reqItem.ItemId != "" { Items[i].ItemId = reqItem.ItemId; changed=true}
		if reqItem.Description != "" { Items[i].Description = reqItem.Description; changed=true}
		if reqItem.UnitPrice    != "" { Items[i].UnitPrice = reqItem.UnitPrice; changed=true }
		if reqItem.PriceMultiplier != "" { Items[i].PriceMultiplier = reqItem.PriceMultiplier; changed=true }
	}
	if changed==true {
		w.WriteHeader(200)
	}
	return
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findItem(id)
	if i == -1 {
		return
	} else {
		deleteItem(id)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findItem(x string) int {
	for i, item := range Items {
		if x == item.ItemId {
			return i
		}
	}
	return -1
}

func deleteItem(id string){
	i := findItem(id)
	Items = append(Items[:i], Items[i+1:]...)
}