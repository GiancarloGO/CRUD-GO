package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"productos-crud/models"
	"productos-crud/utils"
	"strconv"
	"strings"
)

type ProductoController struct {
	templates *template.Template
}

func (pc *ProductoController) init() {
	if pc.templates == nil {
		pc.templates = template.Must(template.ParseGlob("views/*.html"))
	}
}

func (pc *ProductoController) Index(w http.ResponseWriter, r *http.Request) {
	pc.init()

	categoriaID := r.URL.Query().Get("categoria")
	var productos []models.Producto

	if categoriaID != "" {
		catID, err := strconv.Atoi(categoriaID)
		if err == nil {
			productos = models.ProductoRepo.GetByCategoria(catID)
		} else {
			productos = models.ProductoRepo.GetAll()
		}
	} else {
		productos = models.ProductoRepo.GetAll()
	}

	data := struct {
		Productos  []models.Producto
		Categorias []models.Categoria
	}{
		Productos:  productos,
		Categorias: models.CategoriaRepo.GetAll(),
	}

	err := pc.templates.ExecuteTemplate(w, "productos.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *ProductoController) Show(w http.ResponseWriter, r *http.Request, id int) {
	pc.init()

	producto, err := models.ProductoRepo.GetByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	categoria, _ := models.CategoriaRepo.GetByID(producto.CategoriaID)
	categoriaNombre := ""
	if categoria != nil {
		categoriaNombre = categoria.Nombre
	}

	data := struct {
		Producto  *models.Producto
		Categoria string
	}{
		Producto:  producto,
		Categoria: categoriaNombre,
	}

	err = pc.templates.ExecuteTemplate(w, "detalle_producto.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *ProductoController) New(w http.ResponseWriter, r *http.Request) {
	pc.init()

	data := struct {
		Categorias []models.Categoria
		Producto   *models.Producto
		Editar     bool
	}{
		Categorias: models.CategoriaRepo.GetAll(),
		Producto:   nil,
		Editar:     false,
	}

	err := pc.templates.ExecuteTemplate(w, "form_producto.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *ProductoController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB máximo
	if err != nil {
		http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
		return
	}

	precio, err := strconv.ParseFloat(r.FormValue("precio"), 64)
	if err != nil {
		http.Error(w, "Precio inválido", http.StatusBadRequest)
		return
	}

	categoriaID, err := strconv.Atoi(r.FormValue("categoria_id"))
	if err != nil {
		http.Error(w, "Categoría inválida", http.StatusBadRequest)
		return
	}

	// Manejar imagen
	var imagenURL string
	file, fileHeader, err := r.FormFile("imagen")
	if err == nil && fileHeader != nil {
		defer file.Close()
		// Subir archivo
		imagenURL, err = utils.SaveUploadedFile(fileHeader)
		if err != nil {
			log.Printf("Error al subir imagen: %v", err)
			http.Error(w, "Error al subir imagen: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// Si no hay archivo, usar URL proporcionada
		imagenURL = r.FormValue("imagen_url")
	}

	producto := models.Producto{
		Nombre:      r.FormValue("nombre"),
		Descripcion: r.FormValue("descripcion"),
		Precio:      precio,
		ImagenURL:   imagenURL,
		CategoriaID: categoriaID,
	}

	_, err = models.ProductoRepo.Create(producto)
	if err != nil {
		http.Error(w, "Error al crear producto: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	http.Redirect(w, r, "/productos", http.StatusSeeOther)
}

func (pc *ProductoController) Edit(w http.ResponseWriter, r *http.Request, id int) {
	pc.init()

	producto, err := models.ProductoRepo.GetByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	data := struct {
		Categorias []models.Categoria
		Producto   *models.Producto
		Editar     bool
	}{
		Categorias: models.CategoriaRepo.GetAll(),
		Producto:   producto,
		Editar:     true,
	}

	err = pc.templates.ExecuteTemplate(w, "form_producto.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *ProductoController) Update(w http.ResponseWriter, r *http.Request, id int) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB máximo
	if err != nil {
		http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
		return
	}

	precio, err := strconv.ParseFloat(r.FormValue("precio"), 64)
	if err != nil {
		http.Error(w, "Precio inválido", http.StatusBadRequest)
		return
	}

	categoriaID, err := strconv.Atoi(r.FormValue("categoria_id"))
	if err != nil {
		http.Error(w, "Categoría inválida", http.StatusBadRequest)
		return
	}

	// Obtener producto actual para preservar la imagen si no se sube una nueva
	productoActual, err := models.ProductoRepo.GetByID(id)
	if err != nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	// Manejar imagen
	var imagenURL string
	file, fileHeader, err := r.FormFile("imagen")
	if err == nil && fileHeader != nil {
		defer file.Close()
		// Eliminar imagen anterior si no es una URL externa
		if productoActual.ImagenURL != "" && !strings.HasPrefix(productoActual.ImagenURL, "http") {
			utils.DeleteFile(productoActual.ImagenURL)
		}
		// Subir nueva imagen
		imagenURL, err = utils.SaveUploadedFile(fileHeader)
		if err != nil {
			log.Printf("Error al subir imagen: %v", err)
			http.Error(w, "Error al subir imagen: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// Si no hay archivo, usar URL proporcionada o mantener la actual
		if r.FormValue("imagen_url") != "" {
			imagenURL = r.FormValue("imagen_url")
		} else {
			imagenURL = productoActual.ImagenURL
		}
	}

	producto := models.Producto{
		Nombre:      r.FormValue("nombre"),
		Descripcion: r.FormValue("descripcion"),
		Precio:      precio,
		ImagenURL:   imagenURL,
		CategoriaID: categoriaID,
	}

	err = models.ProductoRepo.Update(id, producto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/productos", http.StatusSeeOther)
}

func (pc *ProductoController) Delete(w http.ResponseWriter, r *http.Request, id int) {
	// Obtener producto para eliminar su imagen
	producto, err := models.ProductoRepo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Eliminar imagen si no es una URL externa
	if producto.ImagenURL != "" && !strings.HasPrefix(producto.ImagenURL, "http") {
		utils.DeleteFile(producto.ImagenURL)
	}

	err = models.ProductoRepo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/productos", http.StatusSeeOther)
}

func (pc *ProductoController) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/producto/nuevo" {
		if r.Method == "GET" {
			pc.New(w, r)
		} else if r.Method == "POST" {
			pc.Create(w, r)
		}
		return
	}

	// Manejar /producto/{id} y /producto/{id}/accion
	if strings.HasPrefix(path, "/producto/") {
		segments := strings.Split(path, "/")
		if len(segments) >= 3 {
			idStr := segments[2]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}

			if len(segments) == 3 {
				// /producto/{id}
				pc.Show(w, r, id)
			} else if len(segments) == 4 {
				action := segments[3]
				switch action {
				case "editar":
					if r.Method == "GET" {
						pc.Edit(w, r, id)
					} else if r.Method == "POST" {
						pc.Update(w, r, id)
					}
				case "borrar":
					if r.Method == "POST" {
						pc.Delete(w, r, id)
					}
				}
			}
		}
		return
	}

	http.NotFound(w, r)
}

func (pc *ProductoController) APIIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productos := models.ProductoRepo.GetAll()
	json.NewEncoder(w).Encode(productos)
}
