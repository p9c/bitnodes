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

func AmpSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r){{end}}
	// lang := vars["lang"]
	// site := mod.Site{}

	records, err := JDB.ReadAll("subscriptions")
	if err != nil {
		fmt.Println("Error", err)
	}
	subscriptions := []mod.Subscription{}
	for _, f := range records {
		subscriptionFound := mod.Subscription{}
		if err := json.Unmarshal([]byte(f), &subscriptionFound); err != nil {
			fmt.Println("Error", err)
		}
		subscriptions = append(subscriptions, subscriptionFound)
	}
	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/icons/icons.gohtml",
		"./tpl/amp/lyt/base.gohtml",
		"./tpl/amp/subscriptions.gohtml",
		"./tpl/amp/inc/nav.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/inc/footer.gohtml")
	tmpl.ExecuteTemplate(w, "base", subscriptions)
}

func AdminSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		records, err := JDB.ReadAll("subscriptions")
		if err != nil {
			fmt.Println("Error", err)
		}
		subscriptions := []mod.Subscription{}
		for _, f := range records {
			subscriptionFound := mod.Subscription{}
			if err := json.Unmarshal([]byte(f), &subscriptionFound); err != nil {
				fmt.Println("Error", err)
			}
			subscriptions = append(subscriptions, subscriptionFound)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/subscriptions.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", subscriptions)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminSubscriptionPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// lang := vars["lang"]
	sbsID := vars["sbsid"]
	userName := GetUserName(r)
	var subscription mod.Subscription
	if !hlp.IsEmpty(userName) {
		if err := JDB.Read("subscriptions", sbsID, &subscription); err != nil {
			fmt.Println("Error", err)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/subscription.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", subscription)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
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
