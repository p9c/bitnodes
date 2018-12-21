package rts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alecthomas/template"
	"github.com/zeus16/coinpayments"
)

func AmpCheckoutGetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/amp/lyt/base.gohtml",
		"./tpl/amp/checkout.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/header.gohtml",
		"./tpl/amp/inc/nav.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
	)

	// w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:8999")
	w.Header().Set("Access-Control-Allow-Origin", "localhost:8999")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "localhost:8999")

	r.ParseForm()
	amount := r.FormValue("amount")
	email := r.FormValue("email")

	resp, err := Client.CallCreateTransaction(&coinpayments.TransactionRequest{
		Amount: amount, Currency1: "USD", Currency2: "LTCT", BuyerEmail: email})
	if err != nil {
		fmt.Println("Could not call create transaction: %s", err.Error())
	}

	tmpl.ExecuteTemplate(w, "base", resp)
}

func AmpCheckoutPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8999")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:8999")
	// r.ParseForm()

	// name := r.FormValue("name")
	// pass := r.FormValue("password")
	fmt.Println("ssssssssssssssssadasdasdasdasfdsgfdhgfjghkghkfgh")

	http.Redirect(w, r, "/amp/checkout", 302)
}

func CPay(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	amount := r.FormValue("amount")
	email := r.FormValue("email")
	coin := r.FormValue("coin")

	// nick := r.FormValue("nick")
	// source := r.FormValue("source")

	if email != "" {
		// var jsonresp coinpayments.TransactionResult
		resp, err := Client.CallCreateTransaction(&coinpayments.TransactionRequest{
			Amount: amount, Currency1: "USD", Currency2: coin, BuyerEmail: email})
		if err != nil {
			fmt.Println("Could not call create transaction: %s", err.Error())
		}

		// _, err = Client.CallGetTxInfo(&coinpayments.TxInfoRequest{TxID: resp.TxnID})

		// if err != nil {
		// 	fmt.Println("Could not call get tx info: %s", err.Error())
		// }

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8999")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
		w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:8999")
		//http.Redirect(w, r, resp.StatusURL, 302)
		fmt.Println("emailemailemailemail", email)
		fmt.Println("amountamount", amount)
		fmt.Println("coincoincoincoincoin", coin)
		fmt.Println("PScccccccccccccccc", resp)

		// resp
		resp.Timeout = resp.Timeout + uint32(time.Now().Unix())
		//Marshal or convert user object back to json and write to response
		respJson, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}

		//Set Content-Type header so that clients will know how to read response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//Write json response back to response
		fmt.Println("PScccccccccccccccc", resp)

		w.Write(respJson)
	}
}
