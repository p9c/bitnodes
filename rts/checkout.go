package rts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/bitnodes/mod"
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
	amnt := r.FormValue("amount")
	eml := r.FormValue("email")
	coin := r.FormValue("coin")
	nick := r.FormValue("nick")
	csc := r.FormValue("source")

	rsc := r.FormValue("resources")
	nds := r.FormValue("nodes")

	if eml != "" {
		pack := rsc + " " + nds + " Pack"
		resp, err := Client.CallCreateTransaction(&coinpayments.TransactionRequest{
			Amount: amnt, Currency1: "USD", Currency2: coin, BuyerEmail: eml})
		if err != nil {
			fmt.Println("Could not call create transaction: %s", err.Error())
		}

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8999")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
		w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:8999")

		fmt.Println("emailemailemailemail", eml)
		fmt.Println("amountamount", amnt)
		fmt.Println("coincoincoincoincoin", coin)
		fmt.Println("PScccccccccccccccc", resp)

		resp.Timeout = resp.Timeout + uint32(time.Now().Unix())

		var order mod.Order = mod.Order{
			TxnID:          resp.TxnID,
			Status:         "0",
			Amount:         resp.Amount,
			Address:        resp.Address,
			Nick:           nick,
			CoinSourceCode: csc,
			Currency:       coin,
			Email:          eml,
			Pack:           pack,
			Time:           uint32(time.Now().Unix()),
			Timeout:        resp.Timeout,
			ConfirmsNeeded: resp.ConfirmsNeeded,
			StatusURL:      resp.StatusURL,
			QRCodeURL:      resp.QRCodeURL,
		}

		JDB.Write("orders", resp.TxnID, order)

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

func CPayIPN(w http.ResponseWriter, r *http.Request) {
	reader, _ := r.GetBody()
	// var depres *coinpayments.IPNDepositResponse{}

	// depres, _ := Client.HandleIPNDeposit(reader)

	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8999")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")

	// w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	// w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:8999")

	// fmt.Println("emailemailemailemail", eml)
	// fmt.Println("amountamount", amnt)
	// fmt.Println("coincoincoincoincoin", coin)
	fmt.Println("PScccccccccccccccc", reader)

	// var order mod.Order = mod.Order{
	// 	type IPNAPIResponse struct {
	// 		Status           string `json:"status"`
	// 		StatusText       string `json:"status_text"`
	// 		TxnID            string `json:"txn_id"`
	// 		Currency1        string `json:"currency1"`
	// 		Currency2        string `json:"currency2"`
	// 		Amount1          string `json:"amount1"`
	// 		Amount2          string `json:"amount2"`
	// 		Fee              string `json:"fee"`
	// 		BuyerName        string `json:"buyer_name"`
	// 		Email            string `json:"email"`
	// 		ItemName         string `json:"item_name"`
	// 		ItemNumber       string `json:"item_number"`
	// 		Invoice          string `json:"invoice"`
	// 		Custom           string `json:"custom"`
	// 		SendTX           string `json:"send_tx"` // the tx id of the payment to the merchant. only included when 'status' >= 100 and the payment mode is set to ASAP or nightly or if the payment is paypal passthru
	// 		ReceivedAmount   string `json:"received_amount"`
	// 		ReceivedConfirms string `json:"received_confirms"`
	// 	}
	// }

	// JDB.Write("orders", resp.TxnID, order)

	// respJson, err := json.Marshal(resp)
	// if err != nil {
	// 	panic(err)
	// }

	// //Set Content-Type header so that clients will know how to read response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// //Write json response back to response
	// fmt.Println("PScccccccccccccccc", resp)

	// w.Write(respJson)

}
