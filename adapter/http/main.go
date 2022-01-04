package http

import (
	"net/http"

	"github.com/guifgr/dio-expert-session-go/adapter/http/actuator"
	"github.com/guifgr/dio-expert-session-go/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Init documentation
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/health", actuator.Heath)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
