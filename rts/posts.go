package rts

import (
	"net/http"

	"github.com/parallelcointeam/bitnodes/hlp"
)

func AmpPostsHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r){{end}}
	// lang := vars["lang"]
	// site := mod.Site{}
	// if err := JDB.Read("site", "meta", &site); err != nil {
	// 	fmt.Println("Error", err)
	// }
	// records, err := JDB.ReadAll("post")
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// posts := []mod.Post{}
	// for _, f := range records {
	// 	postFound := mod.Post{}
	// 	if err := json.Unmarshal([]byte(f), &postFound); err != nil {
	// 		fmt.Println("Error", err)
	// 	}
	// 	posts = append(posts, postFound)
	// }
	// data := struct {
	// 	Site  mod.Site   `json:"site"`
	// 	Lang  string     `json:"lang"`
	// 	Posts []mod.Post `json:"posts"`
	// }{
	// 	site,
	// 	lang,
	// 	posts,
	// }
	// //	tmpl, _ := template.New("").ParseFiles("./tpl/amp/index.gohtml", "./tpl/amp/lyt/home.gohtml", "./tpl/amp/inc/amp.gohtml", "./tpl/amp/inc/nav.gohtml", "./tpl/amp/inc/amp-css-home.gohtml", "./tpl/amp/inc/footer.gohtml")
	// tmpl, _ := template.New("").ParseFiles(
	// 	"./tpl/icons/logo.gohtml",
	// 	"./tpl/icons/icons.gohtml",
	// 	"./tpl/amp/lyt/base.gohtml",
	// 	"./tpl/amp/posts.gohtml",
	// 	"./tpl/amp/inc/nav.gohtml",
	// 	"./tpl/amp/inc/amp.gohtml",
	// 	"./tpl/amp/inc/amp-basecss.gohtml",
	// 	"./tpl/amp/inc/amp-basecssplgs.gohtml",
	// 	"./tpl/amp/inc/amp-css.gohtml",
	// 	"./tpl/amp/inc/footer.gohtml")
	// //tmpl, _ := template.New("").ParseFiles("./tpl/amp/lyt/base.gohtml")
	// tmpl.ExecuteTemplate(w, "base", data)
}

func AdminPostsPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		// records, err := JDB.ReadAll("post")
		// if err != nil {
		// 	fmt.Println("Error", err)
		// }
		// posts := []mod.Post{}
		// for _, f := range records {
		// 	postFound := mod.Post{}
		// 	if err := json.Unmarshal([]byte(f), &postFound); err != nil {
		// 		fmt.Println("Error", err)
		// 	}
		// 	posts = append(posts, postFound)
		// }
		// tmpl, _ := template.New("").ParseFiles("./tpl/admin/posts.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		// tmpl.ExecuteTemplate(w, "base", posts)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminPostPageHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// lang := vars["lang"]
	// postID := vars["post"]
	userName := GetUserName(r)
	// var post mod.Post
	if !hlp.IsEmpty(userName) {
		// if err := JDB.Read("post", postID, &post); err != nil {
		// 	fmt.Println("Error", err)
		// }
		// data := struct {
		// 	Lang string   `json:"lang"`
		// 	Post mod.Post `json:"post"`
		// }{
		// 	lang,
		// 	post,
		// }
		// tmpl, _ := template.New("").ParseFiles("./tpl/admin/post.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/nav.gohtml")
		// tmpl.ExecuteTemplate(w, "base", data)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminPostHandler(w http.ResponseWriter, r *http.Request) {
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
