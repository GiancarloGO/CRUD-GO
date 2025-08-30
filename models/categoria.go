package models

import (
	"database/sql"
	"errors"
	"productos-crud/database"
)

type Categoria struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type CategoriaModel struct{}

var CategoriaRepo *CategoriaModel

func init() {
	CategoriaRepo = &CategoriaModel{}
}

func (cm *CategoriaModel) GetAll() []Categoria {
	rows, err := database.DB.Query("SELECT id, nombre FROM categorias ORDER BY nombre")
	if err != nil {
		return []Categoria{}
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var c Categoria
		err := rows.Scan(&c.ID, &c.Nombre)
		if err != nil {
			continue
		}
		categorias = append(categorias, c)
	}
	return categorias
}

func (cm *CategoriaModel) GetByID(id int) (*Categoria, error) {
	var c Categoria
	err := database.DB.QueryRow("SELECT id, nombre FROM categorias WHERE id = ?", id).Scan(&c.ID, &c.Nombre)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("categoría no encontrada")
		}
		return nil, err
	}
	return &c, nil
}

func (cm *CategoriaModel) Create(categoria Categoria) (*Categoria, error) {
	result, err := database.DB.Exec("INSERT INTO categorias (nombre) VALUES (?)", categoria.Nombre)
	if err != nil {
		return nil, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	
	categoria.ID = int(id)
	return &categoria, nil
}

func (cm *CategoriaModel) Update(id int, categoria Categoria) error {
	_, err := database.DB.Exec("UPDATE categorias SET nombre = ? WHERE id = ?", categoria.Nombre, id)
	return err
}

func (cm *CategoriaModel) Delete(id int) error {
	result, err := database.DB.Exec("DELETE FROM categorias WHERE id = ?", id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return errors.New("categoría no encontrada")
	}
	
	return nil
}