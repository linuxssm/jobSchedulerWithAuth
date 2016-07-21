package routes

import (
	"html/template"
	"net/http"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base")
	template.Must(tmpl.Parse(LoginTmplStr))
	tmpl.Execute(w, nil)
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base")
	template.Must(tmpl.Parse(signupTmplStr))
	tmpl.Execute(w, nil)
}
func profileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base")
	template.Must(tmpl.Parse(profileTmplStr))
	tmpl.Execute(w, nil)
}
//func logoutHandler(w http.ResponseWriter, r *http.Request) {
//	tmpl := template.New("base")
//	template.Must(tmpl.Parse(logBaseTmplStr))
//	tmpl.Execute(w, nil)
//}
