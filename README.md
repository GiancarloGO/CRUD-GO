# go.mod
module productos-crud

go 1.21

# README.md
# Sistema CRUD de Productos y Categorías

Este proyecto implementa un sistema completo de gestión de productos y categorías usando Go con arquitectura MVC.

## Características

- **CRUD completo** para productos y categorías
- **Arquitectura MVC** con separación clara de responsabilidades
- **Interfaz web moderna** con HTML, CSS y JavaScript
- **Validaciones** tanto en frontend como backend
- **Diseño responsivo** que funciona en dispositivos móviles
- **Filtrado por categorías**
- **Almacenamiento en memoria** (fácil de migrar a base de datos)

## Tecnologías Utilizadas

- **Backend**: Go (Golang) con net/http
- **Frontend**: HTML5, CSS3, JavaScript
- **Templates**: html/template de Go
- **Almacenamiento**: En memoria (arrays/slices)

## Estructura del Proyecto

```
productos-crud/
├── main.go                          # Punto de entrada y rutas
├── go.mod                          # Configuración del módulo
├── controllers/                    # Controladores (lógica de negocio)
│   ├── producto_controller.go     # Controlador de productos
│   └── categoria_controller.go    # Controlador de categorías
├── models/                         # Modelos (datos)
│   ├── producto.go               # Modelo de producto
│   └── categoria.go              # Modelo de categoría
├── views/                          # Templates HTML
│   ├── productos.html
│   ├── detalle_producto.html
│   ├── form_producto.html
│   └── categorias.html
├── static/                         # Archivos estáticos
│   ├── style.css
│   └── script.js
└── README.md

## Modelos de Datos

### Producto
```go
type Producto struct {
    ID          int
    Nombre      string
    Descripcion string
    Precio      float64
    ImagenURL   string
    CategoriaID int
}
```

### Categoría
```go
type Categoria struct {
    ID     int
    Nombre string
}
```

## Rutas Disponibles

| Ruta | Método | Descripción |
|------|--------|-------------|
| `/` | GET | Redirige a `/productos` |
| `/productos` | GET | Lista todos los productos |
| `/productos?categoria=ID` | GET | Filtra productos por categoría |
| `/producto/ID` | GET | Muestra detalle del producto |
| `/producto/nuevo` | GET | Formulario nuevo producto |
| `/producto/nuevo` | POST | Crea nuevo producto |
| `/producto/ID/editar` | GET | Formulario editar producto |
| `/producto/ID/editar` | POST | Actualiza producto |
| `/producto/ID/borrar` | POST | Elimina producto |
| `/categorias` | GET | Lista categorías |
| `/categorias` | POST | Crea nueva categoría |
| `/api/productos` | GET | API JSON de productos |
| `/api/categorias` | GET | API JSON de categorías |

