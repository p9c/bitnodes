package rts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/bitnodes/cfg"
	"github.com/zeus16/coinpayments"
)

var Client, _ = coinpayments.NewClient(&coinpayments.Config{PublicKey: cfg.PublicKey, PrivateKey: cfg.PrivateKey}, &http.Client{Timeout: 10 * time.Second})

//var Client, _ = coinpayments.NewClient(&coinpayments.Config{PublicKey: cfg.PublicKey, PrivateKey: cfg.PrivateKey, IPNSecret: cfg.IPNSecret, IPNURL: cfg.IPNURL}, &http.Client{Timeout: 10 * time.Second})

func AmpHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := Client.CallRates(&coinpayments.RatesRequest{Short: "0", Accepted: "0"})
	if err != nil {
		fmt.Println("Could not call get tx info: %s", err.Error())
	}
	// fmt.Println("Could not call aaaaalaaaa:", resp)

	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/amp/lyt/base.gohtml",
		"./tpl/amp/home.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/nodes.gohtml",
		"./tpl/amp/inc/header.gohtml",
		"./tpl/amp/inc/nav.gohtml",
		"./tpl/amp/inc/pricing.gohtml",
		"./tpl/amp/inc/products.gohtml",

		"./tpl/amp/inc/features.gohtml",
		"./tpl/amp/inc/nodejson.gohtml",
		"./tpl/amp/inc/front.gohtml",

		"./tpl/amp/inc/order.gohtml",
		"./tpl/amp/inc/checkout.gohtml",
		"./tpl/amp/inc/pay.gohtml",

		"./tpl/amp/inc/amp-ace.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/inc/amp-nodescss.gohtml",
	)
	tmpl.ExecuteTemplate(w, "base", resp)
}

func AmpFrontHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/lyt/front.gohtml",
	)
	tmpl.ExecuteTemplate(w, "front", nil)
}
