package models

import (
	"database/sql"
	"errors"
	"productos-crud/database"
	"time"
)

type Producto struct {
	ID          int       `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Precio      float64   `json:"precio"`
	ImagenURL   string    `json:"imagen_url"`
	CategoriaID int       `json:"categoria_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductoModel struct{}

var ProductoRepo *ProductoModel

func init() {
	ProductoRepo = &ProductoModel{}
}

func (pm *ProductoModel) GetAll() []Producto {
	rows, err := database.DB.Query(`
		SELECT id, nombre, descripcion, precio, imagen_url, categoria_id, created_at, updated_at 
		FROM productos ORDER BY created_at DESC
	`)
	if err != nil {
		return []Producto{}
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.ImagenURL, &p.CategoriaID, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}
		productos = append(productos, p)
	}
	return productos
}

func (pm *ProductoModel) GetByID(id int) (*Producto, error) {
	var p Producto
	err := database.DB.QueryRow(`
		SELECT id, nombre, descripcion, precio, imagen_url, categoria_id, created_at, updated_at 
		FROM productos WHERE id = ?
	`, id).Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.ImagenURL, &p.CategoriaID, &p.CreatedAt, &p.UpdatedAt)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("producto no encontrado")
		}
		return nil, err
	}
	return &p, nil
}

func (pm *ProductoModel) GetByCategoria(categoriaID int) []Producto {
	rows, err := database.DB.Query(`
		SELECT id, nombre, descripcion, precio, imagen_url, categoria_id, created_at, updated_at 
		FROM productos WHERE categoria_id = ? ORDER BY created_at DESC
	`, categoriaID)
	if err != nil {
		return []Producto{}
	}
	defer rows.Close()

	var productos []Producto
	for rows.Next() {
		var p Producto
		err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.ImagenURL, &p.CategoriaID, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}
		productos = append(productos, p)
	}
	return productos
}

func (pm *ProductoModel) Create(producto Producto) (*Producto, error) {
	result, err := database.DB.Exec(`
		INSERT INTO productos (nombre, descripcion, precio, imagen_url, categoria_id) 
		VALUES (?, ?, ?, ?, ?)
	`, producto.Nombre, producto.Descripcion, producto.Precio, producto.ImagenURL, producto.CategoriaID)
	
	if err != nil {
		return nil, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	producto.ID = int(id)
	producto.CreatedAt = time.Now()
	producto.UpdatedAt = time.Now()
	
	return &producto, nil
}

func (pm *ProductoModel) Update(id int, producto Producto) error {
	_, err := database.DB.Exec(`
		UPDATE productos 
		SET nombre = ?, descripcion = ?, precio = ?, imagen_url = ?, categoria_id = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, producto.Nombre, producto.Descripcion, producto.Precio, producto.ImagenURL, producto.CategoriaID, id)
	
	if err != nil {
		return err
	}
	
	return nil
}

func (pm *ProductoModel) Delete(id int) error {
	result, err := database.DB.Exec("DELETE FROM productos WHERE id = ?", id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("producto no encontrado")
	}
	
	return nil
}