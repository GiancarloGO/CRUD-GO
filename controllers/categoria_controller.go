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

// Listado de categorías
func (cc *CategoriaController) Index(w http.ResponseWriter, r *http.Request) {
	cc.init()

	categorias := models.CategoriaRepo.GetAll()
	err := cc.templates.ExecuteTemplate(w, "categorias.html", categorias)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Crear categoría
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

// Mostrar formulario de edición
func (cc *CategoriaController) Edit(w http.ResponseWriter, r *http.Request, id int) {
	cc.init()

	categoria, _ := models.CategoriaRepo.GetByID(id)
	err := cc.templates.ExecuteTemplate(w, "editar_categoria.html", categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Actualizar categoría
func (cc *CategoriaController) Update(w http.ResponseWriter, r *http.Request, id int) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
		return
	}

	categoria := models.Categoria{
		Nombre: r.FormValue("nombre"),
	}

	models.CategoriaRepo.Update(id, categoria)
	http.Redirect(w, r, "/categorias", http.StatusSeeOther)
}

// Eliminar categoría
func (cc *CategoriaController) Delete(w http.ResponseWriter, r *http.Request, id int) {
	models.CategoriaRepo.Delete(id)
	http.Redirect(w, r, "/categorias", http.StatusSeeOther)
}

// Manejo general de categorías (listar / crear)
func (cc *CategoriaController) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cc.Index(w, r)
	} else if r.Method == "POST" {
		cc.Create(w, r)
	}
}

// API (JSON)
func (cc *CategoriaController) APIIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categorias := models.CategoriaRepo.GetAll()
	json.NewEncoder(w).Encode(categorias)
}
