package rts

import (
	"net/http"

	"github.com/alecthomas/template"
)

func AmpHandler(w http.ResponseWriter, r *http.Request) {
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
		"./tpl/amp/inc/order.gohtml",
		"./tpl/amp/inc/selection.gohtml",
		"./tpl/amp/inc/checkout.gohtml",
		"./tpl/amp/inc/features.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/inc/amp-nodescss.gohtml",
	)
	tmpl.ExecuteTemplate(w, "base", nil)
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
