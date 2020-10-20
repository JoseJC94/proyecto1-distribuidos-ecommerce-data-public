package main

/*Requests*/

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleGetOrder(w http.ResponseWriter, r *http.Request) (err error) {
	vars := mux.Vars(r)
	var customerId = vars["customerId"]
	//var orderDetailId =vars["orderDetailId"] // the page

	//id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findCustomerOrder(customerId)
	fmt.Println("GET Order for customer: ", customerId, ", location: ",i)
	if i == -1 {
		dataJson, _ := json.Marshal(Orders[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	var CustomerOrders []Order
	for _, order := range Orders {
		if customerId == order.CustomerId {
			CustomerOrders = append(CustomerOrders, order)
		}
	}

	dataJson, err := json.Marshal(CustomerOrders[1:])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostOrder(w http.ResponseWriter, r *http.Request) (err error) {
	//vars := mux.Vars(r)
	//var orderId = vars["orderId"]
	//var orderDetailId =vars["orderDetailId"] // the page

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	order := Order{}
	json.Unmarshal(body, &order)
	Orders = append(Orders, order)
	w.WriteHeader(200)
	return
}

func handleDeleteOrder(w http.ResponseWriter, r *http.Request) (err error) {
	vars := mux.Vars(r)
	var orderId = vars["orderId"]
	//var orderDetailId =vars["orderDetailId"] // the page

	//id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findOrder(orderId)
	if i == -1 {
		return
	} else {
		deleteOrder(orderId)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findOrder(x string) int {
	for i, order := range Orders {
		if x == order.InvoiceNo {
			return i
		}
	}
	return -1
}

func findCustomerOrder(x string) int {
	for i, order := range Orders {
		if x == order.CustomerId {
			return i
		}
	}
	return -1
}

func deleteOrder(id string){
	i := findOrder(id)
	Orders = append(Orders[:i], Orders[i+1:]...)
}