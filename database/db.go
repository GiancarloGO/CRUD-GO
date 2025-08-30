package database

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./productos.db")
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	createTables()
	log.Println("Base de datos SQLite inicializada correctamente")
}

func createTables() {
	categoriaTable := `
	CREATE TABLE IF NOT EXISTS categorias (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL UNIQUE
	);`

	productoTable := `
	CREATE TABLE IF NOT EXISTS productos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		descripcion TEXT,
		precio REAL NOT NULL,
		imagen_url TEXT,
		categoria_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (categoria_id) REFERENCES categorias (id)
	);`

	// Crear tablas
	if _, err := DB.Exec(categoriaTable); err != nil {
		log.Fatal("Error al crear tabla categorias:", err)
	}

	if _, err := DB.Exec(productoTable); err != nil {
		log.Fatal("Error al crear tabla productos:", err)
	}

	// Insertar datos iniciales si las tablas están vacías
	insertInitialData()
}

func insertInitialData() {
	// Verificar si ya hay categorías
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM categorias").Scan(&count)
	if err != nil {
		log.Fatal("Error al verificar categorías:", err)
	}

	if count == 0 {
		// Insertar categorías iniciales
		categorias := []string{
			"Electrónicos",
			"Ropa",
			"Hogar",
			"Deportes",
			"Libros",
		}

		for _, categoria := range categorias {
			_, err := DB.Exec("INSERT INTO categorias (nombre) VALUES (?)", categoria)
			if err != nil {
				log.Printf("Error al insertar categoría %s: %v", categoria, err)
			}
		}

		// Insertar productos iniciales
		productos := []struct {
			nombre      string
			descripcion string
			precio      float64
			imagenURL   string
			categoriaID int
		}{
			{"Smartphone", "Teléfono inteligente última generación", 299.99, "https://via.placeholder.com/300x200?text=Smartphone", 1},
			{"Camiseta Premium", "Camiseta de algodón 100% orgánico", 19.99, "https://via.placeholder.com/300x200?text=Camiseta", 2},
			{"Lámpara LED", "Lámpara de mesa con tecnología LED", 49.99, "https://via.placeholder.com/300x200?text=Lampara", 3},
		}

		for _, p := range productos {
			_, err := DB.Exec(`
				INSERT INTO productos (nombre, descripcion, precio, imagen_url, categoria_id) 
				VALUES (?, ?, ?, ?, ?)
			`, p.nombre, p.descripcion, p.precio, p.imagenURL, p.categoriaID)
			if err != nil {
				log.Printf("Error al insertar producto %s: %v", p.nombre, err)
			}
		}

		log.Println("Datos iniciales insertados correctamente")
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// CreateUploadsDir crea el directorio uploads si no existe
func CreateUploadsDir() {
	if _, err := os.Stat("static/uploads"); os.IsNotExist(err) {
		err := os.MkdirAll("static/uploads", 0755)
		if err != nil {
			log.Fatal("Error al crear directorio uploads:", err)
		}
		log.Println("Directorio static/uploads creado")
	}
}
