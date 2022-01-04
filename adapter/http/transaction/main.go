package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/guifgr/dio-expert-session-go/models/transaction"
	"github.com/guifgr/dio-expert-session-go/util"
)

//GetTransactions documentation
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transaction = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salario",
			Amount:    1200,
			Type:      0,
			CreatedAt: util.StringToTime("2020-05-04T15:04:05"),
		},
	}
	_ = json.NewEncoder(w).Encode(transaction)
}

//CreateTransactions documentation
func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transaction{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
