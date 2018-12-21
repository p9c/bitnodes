package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parallelcointeam/bitnodes/mod"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	// vars :{{define "base"}}{{end}}= mux.Vars(r)
	// lang := vars["lang"]
	// getLang := mod.Home{}
	// if err := JDB.Read("lang", lang, &getLang); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// jsonLang, err := json.Marshal(getLang)
	// if err != nil {
	// 	fmt.Println("Error encoding JSON")
	// 	return
	// }
	// w.Write([]byte(jsonLang))
}

func ApiOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txnid := vars["txnid"]
	getOrder := mod.Order{}
	if err := JDB.Read("orders", txnid, &getOrder); err != nil {
		fmt.Println("Error", err)
	}
	jsonLang, err := json.Marshal(getOrder)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(jsonLang))
}
