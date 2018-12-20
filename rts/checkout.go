package rts

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
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

	// w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://127.0.0.1:8099")
	w.Header().Set("Access-Control-Allow-Origin", "127.0.0.1:8099")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "127.0.0.1:8099")

	r.ParseForm()
	test := r.Form
	fmt.Println("aaaadsadsdf", test)
	tmpl.ExecuteTemplate(w, "base", nil)
}

func AmpCheckoutPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "127.0.0.1:8099")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "127.0.0.1:8099")
	// r.ParseForm()

	// name := r.FormValue("name")
	// pass := r.FormValue("password")
	fmt.Println("ssssssssssssssssadasdasdasdasfdsgfdhgfjghkghkfgh")

	http.Redirect(w, r, "/amp/checkout", 302)
}
