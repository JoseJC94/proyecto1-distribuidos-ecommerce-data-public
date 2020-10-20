package main

/*Requests*/

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleGetOrderDetail(w http.ResponseWriter, r *http.Request) (err error) {
	vars := mux.Vars(r)
	var orderId = vars["orderId"]
	//var orderDetailDetailId =vars["orderDetailDetailId"] // the page

	//id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findDetailOrder(orderId)
	fmt.Println("GET OrderDetail for customer: ", orderId, ", location: ",i)
	if i == -1 {
		dataJson, _ := json.Marshal(OrderDetails[1:])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	var DetailOrders []OrderDetail
	for _, orderDetail := range OrderDetails {
		if orderId == orderDetail.InvoiceNo {
			DetailOrders = append(DetailOrders, orderDetail)
		}
	}

	dataJson, err := json.Marshal(DetailOrders[1:])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handlePostOrderDetail(w http.ResponseWriter, r *http.Request) (err error) {
	//vars := mux.Vars(r)
	//var orderDetailId = vars["orderDetailId"]
	//var orderDetailDetailId =vars["orderDetailDetailId"] // the page

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	orderDetail := OrderDetail{}
	json.Unmarshal(body, &orderDetail)
	OrderDetails = append(OrderDetails, orderDetail)
	w.WriteHeader(200)
	return
}

func handleDeleteOrderDetail(w http.ResponseWriter, r *http.Request) (err error) {
	vars := mux.Vars(r)
	var orderDetailId = vars["orderDetailId"]
	//var orderDetailDetailId =vars["orderDetailDetailId"] // the page

	//id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findOrderDetail(orderDetailId)
	if i == -1 {
		return
	} else {
		deleteOrderDetail(orderDetailId)
	}
	w.WriteHeader(200)
	return
}

/// Functions

func findDetailOrder(x string) int {
	for i, orderDetail := range OrderDetails {
		if x == orderDetail.InvoiceNo {
			return i
		}
	}
	return -1
}

func findOrderDetail(x string) int {
	for i, orderDetail := range OrderDetails {
		if x == orderDetail.InvoiceLineId {
			return i
		}
	}
	return -1
}

func deleteOrderDetail(id string){
	i := findOrderDetail(id)
	OrderDetails = append(OrderDetails[:i], OrderDetails[i+1:]...)
}