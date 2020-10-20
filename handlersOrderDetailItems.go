package main

/*Requests*/

import (
"encoding/json"
"fmt"
"github.com/gorilla/mux"
"net/http"
)

func handleGetOrderDetailItem(w http.ResponseWriter, r *http.Request) (err error) {
	vars := mux.Vars(r)
	var itemId = vars["itemId"]
	//var orderDetailItemDetailId =vars["orderDetailItemDetailId"] // the page

	//id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findItemDetailItem(itemId)
	fmt.Println("GET OrderDetailItem for customer: ", itemId, ", location: ",i)
	if i == -1 {
		dataJson, _ := json.Marshal(OrderDetailItems[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	var CustomerOrderDetailItems []OrderDetailItem
	for _, orderDetailItem := range OrderDetailItems {
		if itemId == orderDetailItem.ItemId {
			CustomerOrderDetailItems = append(CustomerOrderDetailItems, orderDetailItem)
		}
	}

	dataJson, err := json.Marshal(CustomerOrderDetailItems[1:])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostOrderDetailItem(w http.ResponseWriter, r *http.Request) (err error) {
	//vars := mux.Vars(r)
	//var orderDetailItemId = vars["orderDetailItemId"]
	//var orderDetailItemDetailId =vars["orderDetailItemDetailId"] // the page

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	orderDetailItem := OrderDetailItem{}
	json.Unmarshal(body, &orderDetailItem)
	OrderDetailItems = append(OrderDetailItems, orderDetailItem)
	w.WriteHeader(200)
	return
}

func handleDeleteOrderDetailItem(w http.ResponseWriter, r *http.Request) (err error) {
	vars := mux.Vars(r)
	var orderDetailItemId = vars["orderDetailId"]
	//var orderDetailItemDetailId =vars["orderDetailItemDetailId"] // the page

	//id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findOrderDetailItem(orderDetailItemId)
	if i == -1 {
		return
	} else {
		deleteOrderDetailItem(orderDetailItemId)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findItemDetailItem(x string) int {
	for i, orderDetailItem := range OrderDetailItems {
		if x == orderDetailItem.ItemId {
			return i
		}
	}
	return -1
}

func findOrderDetailItem(x string) int {
	for i, orderDetailItem := range OrderDetailItems {
		if x == orderDetailItem.InvoiceLineId {
			return i
		}
	}
	return -1
}

func deleteOrderDetailItem(id string){
	i := findOrderDetailItem(id)
	OrderDetailItems = append(OrderDetailItems[:i], OrderDetailItems[i+1:]...)
}
