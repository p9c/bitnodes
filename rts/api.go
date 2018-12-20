package rts

import (
	"net/http"
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
