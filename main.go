//Service

package main

import (
    "github.com/go-kit/kit/log"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)
    r := mux.NewRouter()
    http.Handle("/", r)

    //Relations
    r.HandleFunc("/customers/{customerId}/orders", handlerOrder).Methods("GET","POST")
    r.HandleFunc("/customers/{customerId}/orders/{orderId}", handlerOrder).Methods("DELETE")

    r.HandleFunc("/orders/{orderId}/orderDetails", handlerOrderDetail).Methods("GET","POST")
    r.HandleFunc("/orders/{orderId}/orderDetails/{orderDetailId}", handlerOrderDetail).Methods("DELETE")

    r.HandleFunc("/items/{itemId}/orderDetails", handlerOrderDetailItems).Methods("GET","POST")
    r.HandleFunc("/items/{itemId}/orderDetails/{orderDetailId}", handlerOrderDetailItems).Methods("DELETE")

    //Entities endpoints
    r.HandleFunc("/customers", handlerCustomer).Methods("GET","POST")
    r.HandleFunc("/customers/{id}", handlerCustomer).Methods("GET","PUT","DELETE")

    r.HandleFunc("/invoices", handlerInvoice).Methods("GET","POST")
    r.HandleFunc("/invoices/{id}", handlerInvoice).Methods("GET","PUT","DELETE")

    r.HandleFunc("/invoiceLines", handlerInvoiceLines).Methods("GET","POST")
    r.HandleFunc("/invoiceLines/{id}", handlerInvoiceLines).Methods("GET","PUT","DELETE")

    r.HandleFunc("/items", handlerItem).Methods("GET","POST")
    r.HandleFunc("/items/{id}", handlerItem).Methods("GET","PUT","DELETE")

    //r.HandleFunc("/orders/{orderId}/orderDetails", handlerOrder).Methods("GET","POST")
    //r.HandleFunc("/orders/{orderId}/orderDetails/{orderDetailId}", handlerOrder).Methods("DELETE")

    /*
    Query examples:

    Entities:
    customers, invoices, invoiceLines, items

    CRUD Example:

    GET customers
    GET customers/id/
    POST customers
    PUT customer/update
    DELETE customer/id

    Relations:

    GET customers/id/orders
    GET orders/id/orderDetails
    GET items/id/orderDetails

    POST /customers/17850/orders/
         {"invoiceNo":"536365"}

    POST /order/536367/orderDetails/
         {"invoiceLineId":"10"}

    DELETE /customers/17850/orders/536365

    DELETE /order/536367/orderDetails/10
    */

    //http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

/*
http://localhost:8080/customers/12797
http://localhost:8080/customers/12797/orders

*/