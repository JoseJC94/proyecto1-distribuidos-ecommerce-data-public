package main

import (
	"net/http"
)

//Entities handlers

func handlerCustomer(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataCustomer("data/output/customers.csv")
	switch request.Method {
	case "GET":
		err = handleGetCustomer(writer, request)
	case "POST":
		err = handlePostCustomer(writer, request)
	case "PUT":
		err = handlePutCustomer(writer, request)
	case "DELETE":
		err = handleDeleteCustomer(writer, request)
	}
	writeDataCustomer("data/output/customers.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerInvoice(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataInvoice("data/output/invoices.csv")
	switch request.Method {
	case "GET":
		err = handleGetInvoice(writer, request)
	case "POST":
		err = handlePostInvoice(writer, request)
	case "PUT":
		err = handlePutInvoice(writer, request)
	case "DELETE":
		err = handleDeleteInvoice(writer, request)
	}
	writeDataInvoice("data/output/invoices.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerInvoiceLines(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataInvoiceLines("data/output/invoiceLines.csv")
	switch request.Method {
	case "GET":
		err = handleGetInvoiceLines(writer, request)
	case "POST":
		err = handlePostInvoiceLines(writer, request)
	case "PUT":
		err = handlePutInvoiceLines(writer, request)
	case "DELETE":
		err = handleDeleteInvoiceLines(writer, request)
	}
	writeDataInvoiceLines("data/output/invoiceLines.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerItem(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataItem("data/output/items.csv")
	switch request.Method {
	case "GET":
		err = handleGetItem(writer, request)
	case "POST":
		err = handlePostItem(writer, request)
	case "PUT":
		err = handlePutItem(writer, request)
	case "DELETE":
		err = handleDeleteItem(writer, request)
	}
	writeDataItem("data/output/items.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Relations handlers
func handlerOrder(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataOrder("data/output/orders.csv")

	switch request.Method {
	case "GET":
		err = handleGetOrder(writer, request)
	case "POST":
		err = handlePostOrder(writer, request)
	case "DELETE":
		err = handleDeleteOrder(writer, request)
	}
	writeDataOrder("data/output/orders.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerOrderDetail(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataOrderDetail("data/output/orderDetails.csv")

	switch request.Method {
	case "GET":
		err = handleGetOrderDetail(writer, request)
	case "POST":
		err = handlePostOrderDetail(writer, request)
	case "DELETE":
		err = handleDeleteOrderDetail(writer, request)
	}
	writeDataOrderDetail("data/output/orderDetails.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerOrderDetailItems(writer http.ResponseWriter, request *http.Request) {
	var err error
	readDataOrderDetailItem("data/output/orderDetailItems.csv")

	switch request.Method {
	case "GET":
		err = handleGetOrderDetailItem(writer, request)
	case "POST":
		err = handlePostOrderDetailItem(writer, request)
	case "DELETE":
		err = handleDeleteOrderDetailItem(writer, request)
	}
	writeDataOrderDetailItem("data/output/orderDetailItems.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}


