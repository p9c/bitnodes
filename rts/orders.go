package rts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/bitnodes/hlp"
	"github.com/parallelcointeam/bitnodes/mod"
)

func AmpOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r){{end}}
	// lang := vars["lang"]
	// site := mod.Site{}

	records, err := JDB.ReadAll("orders")
	if err != nil {
		fmt.Println("Error", err)
	}
	orders := []mod.Order{}
	for _, f := range records {
		orderFound := mod.Order{}
		if err := json.Unmarshal([]byte(f), &orderFound); err != nil {
			fmt.Println("Error", err)
		}
		orders = append(orders, orderFound)
	}
	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/icons/icons.gohtml",
		"./tpl/amp/lyt/base.gohtml",
		"./tpl/amp/orders.gohtml",
		"./tpl/amp/inc/nav.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/inc/footer.gohtml")
	tmpl.ExecuteTemplate(w, "base", orders)
}

func AdminOrdersHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		records, err := JDB.ReadAll("orders")
		if err != nil {
			fmt.Println("Error", err)
		}
		orders := []mod.Order{}
		for _, f := range records {
			orderFound := mod.Order{}
			if err := json.Unmarshal([]byte(f), &orderFound); err != nil {
				fmt.Println("Error", err)
			}
			orders = append(orders, orderFound)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/orders.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", orders)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminOrderPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// lang := vars["lang"]
	ordID := vars["ordid"]
	userName := GetUserName(r)
	var order mod.Order
	if !hlp.IsEmpty(userName) {
		if err := JDB.Read("orders", ordID, &order); err != nil {
			fmt.Println("Error", err)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/order.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", order)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminOrderHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		// r.ParseForm()
		// lang := r.FormValue("lang")
		// id := r.FormValue("id")
		// date := r.FormValue("date")
		// title := r.FormValue("title")
		// excerpt := r.FormValue("excerpts")
		// content := r.FormValue("content")
		// image := r.FormValue("image")

		// var post mod.Post = mod.Post{
		// 	Date:    date,
		// 	Title:   title,
		// 	Excerpt: excerpt,
		// 	Content: content,
		// 	Image:   image,
		// }
		// JDB.Write("post", id, post)
		// http.Redirect(w, r, "/admin/"+lang+"/home", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
