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

func AmpTutorialsHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r){{end}}
	// lang := vars["lang"]
	// site := mod.Site{}

	records, err := JDB.ReadAll("tutorials")
	if err != nil {
		fmt.Println("Error", err)
	}
	tutorials := []mod.Tutorial{}
	for _, f := range records {
		tutorialFound := mod.Tutorial{}
		if err := json.Unmarshal([]byte(f), &tutorialFound); err != nil {
			fmt.Println("Error", err)
		}
		tutorials = append(tutorials, tutorialFound)
	}
	tmpl, _ := template.New("").ParseFiles(
		"./tpl/icons/logo.gohtml",
		"./tpl/icons/icons.gohtml",
		"./tpl/amp/lyt/base.gohtml",
		"./tpl/amp/tutorials.gohtml",
		"./tpl/amp/inc/nav.gohtml",
		"./tpl/amp/inc/amp.gohtml",
		"./tpl/amp/inc/amp-basecss.gohtml",
		"./tpl/amp/inc/amp-basecssplgs.gohtml",
		"./tpl/amp/inc/amp-css.gohtml",
		"./tpl/amp/inc/footer.gohtml")
	tmpl.ExecuteTemplate(w, "base", tutorials)
}

func AdminTutorialsHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		records, err := JDB.ReadAll("tutorials")
		if err != nil {
			fmt.Println("Error", err)
		}
		tutorials := []mod.Tutorial{}
		for _, f := range records {
			tutorialFound := mod.Tutorial{}
			if err := json.Unmarshal([]byte(f), &tutorialFound); err != nil {
				fmt.Println("Error", err)
			}
			tutorials = append(tutorials, tutorialFound)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/tutorials.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", tutorials)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminTutorialPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// lang := vars["lang"]
	tutID := vars["tutid"]
	userName := GetUserName(r)
	var tutorial mod.Tutorial
	if !hlp.IsEmpty(userName) {
		if err := JDB.Read("tutorials", tutID, &tutorial); err != nil {
			fmt.Println("Error", err)
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/tutorial.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", tutorial)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminTutorialHandler(w http.ResponseWriter, r *http.Request) {
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
