package api

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"flag"
)
var Store *Storage

var (
	userfile   = flag.String("u", "skeddy.db", "use existing storage.")
)

func Init(){
	ConfigDatabase("mysql")
	var err error

	Store, err = NewStorage(*userfile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

}


func HandleUsers(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")

	var users []*User
	//users = make([]*User,0)

	users = Store.AllEntries();
	//users = append(users, &User{"ssm","ssmPass"})
	//users = append(users, &User{"linuxs","linuxsPass"})

	resJson, err := json.Marshal(users)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(resJson))
}


func HandleUser(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	id := vars["id"]
	log.Println("Request for:", id)
	res.Header().Set("Content-Type", "application/json")

	var users []*User
	users = Store.AllEntries();

	var user *User
	for i, item := range users {
		log.Println(users[i], item, i)
		if item.Id == id {
			user = item
			break
		}
	}

	resJson, err := json.Marshal(user)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res, string(resJson))
}