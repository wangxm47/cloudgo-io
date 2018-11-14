package service

import (
	"html/template"
	"net/http"

	"github.com/unrolled/render"
)

func indexHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./templates/index.html")
		t.Execute(w, nil)
	}
}
func signupHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./templates/form.html")
		t.Execute(w, nil)
	}
}
func submitHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		name := r.FormValue("name")
		age := r.FormValue("age")
		email := r.FormValue("email")

		formatter.HTML(w, http.StatusOK, "table", struct {
			Name  string
			Age   string
			Email string
		}{Name: name, Age: age, Email: email})
	}
}
func unknowHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./templates/unknow.html")
		t.Execute(w, nil)
	}
}
