package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/bitnodes/pay"
	"github.com/parallelcointeam/bitnodes/rts"
)

var r = mux.NewRouter()

func main() {

	r.PathPrefix("/s/").Handler(http.StripPrefix("/s/", http.FileServer(http.Dir("./tpl/static/"))))
	// r.PathPrefix("/amp/").Handler(http.StripPrefix("/amp/", http.FileServer(http.Dir("./tpl/amp/"))))

	r.HandleFunc("/admin/login/", rts.LoginPageHandler) // GET
	r.HandleFunc("/admin/login", rts.LoginHandler).Methods("POST")
	r.HandleFunc("/admin/", rts.AdminHomeHandler).Methods("GET")
	r.HandleFunc("/admin/", rts.AdminHandler).Methods("POST")

	r.HandleFunc("/logout", rts.LogoutHandler).Methods("POST")

	r.HandleFunc("/admin/orders/", rts.AdminOrdersHandler).Methods("GET")
	r.HandleFunc("/admin/order/{ordid}", rts.AdminOrderPageHandler).Methods("GET")
	r.HandleFunc("/admin/order", rts.AdminOrderHandler).Methods("POST")

	r.HandleFunc("/admin/subscriptions/", rts.AdminSubscriptionsHandler).Methods("GET")
	r.HandleFunc("/admin/subscription/{sbsid}", rts.AdminSubscriptionPageHandler).Methods("GET")
	r.HandleFunc("/admin/subscription", rts.AdminSubscriptionHandler).Methods("POST")

	r.HandleFunc("/admin/nodes/", rts.AdminNodesHandler).Methods("GET")
	r.HandleFunc("/admin/node/{nodeid}", rts.AdminNodePageHandler).Methods("GET")
	r.HandleFunc("/admin/node", rts.AdminNodeHandler).Methods("POST")

	// r.HandleFunc("/admin/customers", rts.AdminCustomerHandler).Methods("GET")
	// r.HandleFunc("/admin/customer/{customerid}", rts.AdminCustomerPageHandler).Methods("GET")
	// r.HandleFunc("/admin/customer", rts.AdminCustomerHandler).Methods("POST")

	r.HandleFunc("/admin/tutorials/", rts.AdminTutorialsHandler).Methods("GET")
	r.HandleFunc("/admin/tutorial/{tutid}", rts.AdminTutorialPageHandler).Methods("GET")
	r.HandleFunc("/admin/tutorial", rts.AdminTutorialHandler).Methods("POST")

	r.HandleFunc("/admin/announcements", rts.AdminAnnouncementsHandler).Methods("GET")
	r.HandleFunc("/admin/announcement/{annid}", rts.AdminAnnouncementPageHandler).Methods("GET")
	r.HandleFunc("/admin/announcement", rts.AdminAnnouncementHandler).Methods("POST")

	// r.HandleFunc("/api/{lang}", rts.ApiHandler)

	r.HandleFunc("/api/pay", pay.CPay).Methods("GET")
	// r.HandleFunc("/api/ipn", pay.CPayIPN).Methods("POST")

	r.HandleFunc("/amp", rts.AmpHandler)
	// r.HandleFunc("/shop/{page}", rts.ShopHandler)

	// r.HandleFunc("/amp/addtocart", rts.AddToCart)
	// r.HandleFunc("/amp/cart", rts.Cart)

	// r.HandleFunc("/amp/co", rts.AmpCheckoutPostHandler).Methods("POST")
	// r.HandleFunc("/amp/checkout", rts.AmpCheckoutGetHandler).Methods("GET")

	// r.HandleFunc("/amp/{lang}/log/", rts.AmpPostsHandler)

	// http.Handle("/", r)
	// http.ListenAndServe(":8099", nil)

	r.HandleFunc("/", rts.AmpFrontHandler) // GET

	// log.Fatal(http.ListenAndServe(":8099", (r)))
	//	log.Fatal(http.ListenAndServe(":8099", handlers.CORS()(r)))

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "AMP-Access-Control-Allow-Source-Origin"})
	// originsOk := handlers.AllowedOrigins([]string{"http://127.0.0.1:8099"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	// originsOk := handlers.AllowedOrigins([]string{"http://localhost:8099"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	// log.Fatal(http.ListenAndServe(":8099", r))

	log.Fatal(http.ListenAndServe(":8099", handlers.CORS()(r)))
	// log.Fatal(http.ListenAndServe(":8099", r))

	// r.Schemes("https")

	// go log.Fatal(http.ListenAndServe(":80", handlers.CORS()(r)))

	// srv := &http.Server{
	// 	Handler:      handlers.CORS()(r),
	// 	Addr:         "bitnodes.net:443",
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// go log.Fatal(srv.ListenAndServeTLS("c.crt", "k.key"))

}

// func EnableCors(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		origin := GetOrigin(r)
// 		w.Header().Set("Access-Control-Allow-Origin", origin)
// 		// w.Header().Set("Access-Control-Allow-Origin", origin)
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		sourceOrigin := GetSourceOrigin(r)
// 		if sourceOrigin == "" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
// 		w.Header().Set("AMP-Access-Control-Allow-Source-Origin", sourceOrigin)
// 		next.ServeHTTP(w, r)
// 	}
// }

// func GetOrigin(r *http.Request) string {
// 	origin := r.Header.Get("Origin")
// 	if origin != "" {
// 		return origin
// 	}
// 	if r.Header.Get("amp-same-origin") == "true" {
// 		return GetSourceOrigin(r)
// 	}
// 	return "*"
// }

// func GetSourceOrigin(r *http.Request) string {
// 	// TODO perform checks if source origin is allowed
// 	return r.URL.Query().Get("__amp_source_origin")
// }
