package httpHandler

import (
	"fmt"
	"net/http"
	"log"

	"github.com/linuxssm/frontendServer/jobScheduler"
	"github.com/gorilla/mux"
	"github.com/linuxssm/frontendServer/authentication"
)


func StartAdminInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	jobMux := mux.NewRouter()

	log.Printf("Starting admin interface at http://%s\n", addr)
	http.HandleFunc("/", jobScheduler.ViewHandler)
	http.HandleFunc("/assets/css/", jobScheduler.CssHandler)
	http.HandleFunc("/assets/js/", jobScheduler.JsHandler)
	http.HandleFunc("/new/", jobScheduler.NewHandler)
	http.HandleFunc("/edit/", jobScheduler.EditHandler)
	http.HandleFunc("/save/", jobScheduler.SaveHandler)
	http.HandleFunc("/add/", jobScheduler.AddHandler)
	http.HandleFunc("/delete/", jobScheduler.DeleteHandler)
	http.HandleFunc("/validate/", jobScheduler.ValidateExpression)

	//http.Handle("/", jobMux)
	//go http.ListenAndServe(addr, nil)

	//r := mux.NewRouter()
	jobMux.HandleFunc("/auth/login", authentication.GetLogin)
	jobMux.HandleFunc("/auth/register", authentication.PostRegister).Methods("POST")
	jobMux.HandleFunc("/auth/loginPost", authentication.PostLogin).Methods("POST")
	jobMux.HandleFunc("/auth/admin", authentication.HandleAdmin).Methods("GET")
	jobMux.HandleFunc("/auth/add_user", authentication.PostAddUser).Methods("POST")
	jobMux.HandleFunc("/auth/change", authentication.PostChange).Methods("POST")
	jobMux.HandleFunc("/auth/", authentication.HandlePage).Methods("GET") // authorized page
	jobMux.HandleFunc("/auth/logout", authentication.HandleLogout)


	http.Handle("/auth/", jobMux)
	fmt.Printf("Server running on port %d\n", port)
	//http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	http.ListenAndServe(addr, nil)
    //fmt.Println(http.ListenAndServe(":8080", nil))

	//redirectPath
}
