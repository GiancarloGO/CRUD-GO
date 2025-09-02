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
├── main.go                       # Punto de entrada de la aplicación y definición de rutas.
├── go.mod                        # Archivo de configuración que gestiona las dependencias del módulo Go.
├── go.sum                        # Archivo de verificación de sumas de control para las dependencias.
├── README.md                     # Documentación del proyecto (descripción, instrucciones de uso, etc.).
|
├── controllers/                  # Contiene la lógica de negocio y maneja las peticiones HTTP.
│   ├── categoria_controller.go   # Maneja las operaciones CRUD (Crear, Leer, Actualizar, Borrar) para las categorías.
│   └── producto_controller.go    # Maneja las operaciones CRUD para los productos.
|
├── models/                       # Define las estructuras de datos (modelos) que representan las tablas de la base de datos.
│   ├── categoria.go              # Modelo para la entidad `Categoria`.
│   └── producto.go               # Modelo para la entidad `Producto`.
|
├── database/                     # Contiene los archivos relacionados con la conexión y la configuración de la base de datos.
│   └── db.go                     # Lógica para establecer y gestionar la conexión a la base de datos.
|
├── static/                       # Contiene los archivos estáticos que se sirven directamente al navegador.
│   ├── images/                   # Directorio para imágenes del sitio web.
│   │   └── master_color_logo.jpg # Un archivo de imagen de ejemplo.
│   ├── uploads/                  # Directorio para archivos subidos por los usuarios.
│   ├── script.js                 # Scripts de JavaScript para funcionalidades del lado del cliente.
│   └── style.css                 # Hojas de estilo CSS para el diseño de la interfaz.
|
├── utils/                        # Contiene funciones de utilidad que pueden ser usadas en múltiples partes de la aplicación.
│   └── upload.go                 # Lógica para manejar la subida de archivos (por ejemplo, imágenes de productos).
|
└── views/                        # Almacena los templates HTML que se renderizan para mostrar la interfaz al usuario.
    ├── categorias.html           # Vista para mostrar la lista de categorías.
    ├── detalle_producto.html     # Vista para mostrar los detalles de un producto específico.
    ├── editar_categoria.html     # Vista con un formulario para editar una categoría existente.
    ├── form_producto.html        # Vista con un formulario para crear un nuevo producto.
    ├── layout.html               # Plantilla base (layout) que define la estructura principal de las páginas.
    └── productos.html            # Vista para mostrar la lista de todos los productos.


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

