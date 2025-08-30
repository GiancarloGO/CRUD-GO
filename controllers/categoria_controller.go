package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"productos-crud/models"
)

type CategoriaController struct {
	templates *template.Template
}

func (cc *CategoriaController) init() {
	if cc.templates == nil {
		cc.templates = template.Must(template.ParseGlob("views/*.html"))
	}
}

func (cc *CategoriaController) Index(w http.ResponseWriter, r *http.Request) {
	cc.init()

	categorias := models.CategoriaRepo.GetAll()
	err := cc.templates.ExecuteTemplate(w, "categorias.html", categorias)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (cc *CategoriaController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
		return
	}

	categoria := models.Categoria{
		Nombre: r.FormValue("nombre"),
	}

	models.CategoriaRepo.Create(categoria)
	http.Redirect(w, r, "/categorias", http.StatusSeeOther)
}

func (cc *CategoriaController) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cc.Index(w, r)
	} else if r.Method == "POST" {
		cc.Create(w, r)
	}
}

func (cc *CategoriaController) APIIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categorias := models.CategoriaRepo.GetAll()
	json.NewEncoder(w).Encode(categorias)
}
