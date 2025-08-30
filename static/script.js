// static/script.js

// Validación del formulario de productos
function validarFormulario() {
    const nombre = document.getElementById('nombre');
    const descripcion = document.getElementById('descripcion');
    const precio = document.getElementById('precio');
    const categoria = document.getElementById('categoria_id');
    
    let esValido = true;
    
    // Limpiar mensajes de error previos
    limpiarErrores();
    
    // Validar nombre
    if (!nombre.value.trim()) {
        mostrarError('nombre-error', 'El nombre es obligatorio');
        esValido = false;
    } else if (nombre.value.trim().length < 2) {
        mostrarError('nombre-error', 'El nombre debe tener al menos 2 caracteres');
        esValido = false;
    }
    
    // Validar descripción
    if (!descripcion.value.trim()) {
        mostrarError('descripcion-error', 'La descripción es obligatoria');
        esValido = false;
    } else if (descripcion.value.trim().length < 10) {
        mostrarError('descripcion-error', 'La descripción debe tener al menos 10 caracteres');
        esValido = false;
    }
    
    // Validar precio
    const precioVal = parseFloat(precio.value);
    if (!precio.value || isNaN(precioVal)) {
        mostrarError('precio-error', 'El precio es obligatorio y debe ser un número');
        esValido = false;
    } else if (precioVal <= 0) {
        mostrarError('precio-error', 'El precio debe ser mayor que 0');
        esValido = false;
    } else if (precioVal > 999999.99) {
        mostrarError('precio-error', 'El precio no puede ser mayor a $999,999.99');
        esValido = false;
    }
    
    // Validar categoría
    if (!categoria.value) {
        mostrarError('categoria-error', 'Debe seleccionar una categoría');
        esValido = false;
    }
    
    // Validar URL de imagen (opcional pero si se proporciona debe ser válida)
    const imagenUrl = document.getElementById('imagen_url');
    if (imagenUrl.value && !esUrlValida(imagenUrl.value)) {
        mostrarError('imagen-error', 'La URL de la imagen no es válida');
        esValido = false;
    }
    
    return esValido;
}

// Validación del formulario de categorías
function validarCategoria() {
    const nombre = document.getElementById('nombre');
    let esValido = true;
    
    // Limpiar mensajes de error previos
    limpiarErrores();
    
    // Validar nombre de categoría
    if (!nombre.value.trim()) {
        mostrarError('categoria-nombre-error', 'El nombre de la categoría es obligatorio');
        esValido = false;
    } else if (nombre.value.trim().length < 2) {
        mostrarError('categoria-nombre-error', 'El nombre debe tener al menos 2 caracteres');
        esValido = false;
    } else if (nombre.value.trim().length > 50) {
        mostrarError('categoria-nombre-error', 'El nombre no puede tener más de 50 caracteres');
        esValido = false;
    }
    
    return esValido;
}

// Función para mostrar errores
function mostrarError(elementId, mensaje) {
    const errorElement = document.getElementById(elementId);
    if (errorElement) {
        errorElement.textContent = mensaje;
        errorElement.style.display = 'block';
    }
}

// Función para limpiar errores
function limpiarErrores() {
    const errorElements = document.querySelectorAll('.error-message');
    errorElements.forEach(element => {
        element.textContent = '';
        element.style.display = 'none';
    });
}

// Validar URL
function esUrlValida(url) {
    try {
        new URL(url);
        return url.startsWith('http://') || url.startsWith('https://');
    } catch {
        return false;
    }
}

// Confirmar eliminación
function confirmarEliminacion() {
    return confirm('¿Estás seguro de que deseas eliminar este producto? Esta acción no se puede deshacer.');
}

// Filtrar productos por categoría
function filtrarPorCategoria() {
    const select = document.getElementById('categoriaFilter');
    const categoriaId = select.value;
    
    if (categoriaId) {
        window.location.href = `/productos?categoria=${categoriaId}`;
    } else {
        window.location.href = '/productos';
    }
}

// Validación en tiempo real
document.addEventListener('DOMContentLoaded', function() {
    // Validación en tiempo real para el formulario de productos
    const productoForm = document.getElementById('productoForm');
    if (productoForm) {
        const inputs = productoForm.querySelectorAll('input, textarea, select');
        
        inputs.forEach(input => {
            input.addEventListener('blur', function() {
                validarCampoIndividual(this);
            });
            
            input.addEventListener('input', function() {
                // Limpiar error cuando el usuario empieza a escribir
                const errorId = this.id + '-error';
                const errorElement = document.getElementById(errorId);
                if (errorElement) {
                    errorElement.textContent = '';
                    errorElement.style.display = 'none';
                }
            });
        });
    }
    
    // Validación en tiempo real para el formulario de categorías
    const categoriaForm = document.getElementById('categoriaForm');
    if (categoriaForm) {
        const nombreInput = document.getElementById('nombre');
        if (nombreInput) {
            nombreInput.addEventListener('blur', function() {
                validarCampoCategoria(this);
            });
            
            nombreInput.addEventListener('input', function() {
                const errorElement = document.getElementById('categoria-nombre-error');
                if (errorElement) {
                    errorElement.textContent = '';
                    errorElement.style.display = 'none';
                }
            });
        }
    }
    
    // Establecer filtro actual en el selector
    const urlParams = new URLSearchParams(window.location.search);
    const categoriaActual = urlParams.get('categoria');
    if (categoriaActual) {
        const select = document.getElementById('categoriaFilter');
        if (select) {
            select.value = categoriaActual;
        }
    }
    
    // Previsualización de imagen
    const imagenUrlInput = document.getElementById('imagen_url');
    if (imagenUrlInput) {
        imagenUrlInput.addEventListener('input', function() {
            previsualizarImagen(this.value);
        });
        
        // Previsualizar imagen inicial si existe
        if (imagenUrlInput.value) {
            previsualizarImagen(imagenUrlInput.value);
        }
    }
});

// Validar campo individual
function validarCampoIndividual(campo) {
    const valor = campo.value.trim();
    const errorId = campo.id + '-error';
    
    switch(campo.id) {
        case 'nombre':
            if (!valor) {
                mostrarError(errorId, 'El nombre es obligatorio');
            } else if (valor.length < 2) {
                mostrarError(errorId, 'El nombre debe tener al menos 2 caracteres');
            }
            break;
            
        case 'descripcion':
            if (!valor) {
                mostrarError(errorId, 'La descripción es obligatoria');
            } else if (valor.length < 10) {
                mostrarError(errorId, 'La descripción debe tener al menos 10 caracteres');
            }
            break;
            
        case 'precio':
            const precio = parseFloat(campo.value);
            if (!campo.value || isNaN(precio)) {
                mostrarError(errorId, 'El precio es obligatorio y debe ser un número');
            } else if (precio <= 0) {
                mostrarError(errorId, 'El precio debe ser mayor que 0');
            } else if (precio > 999999.99) {
                mostrarError(errorId, 'El precio no puede ser mayor a $999,999.99');
            }
            break;
            
        case 'categoria_id':
            if (!campo.value) {
                mostrarError(errorId, 'Debe seleccionar una categoría');
            }
            break;
            
        case 'imagen_url':
            if (valor && !esUrlValida(valor)) {
                mostrarError(errorId, 'La URL de la imagen no es válida');
            }
            break;
    }
}

// Validar campo de categoría
function validarCampoCategoria(campo) {
    const valor = campo.value.trim();
    
    if (!valor) {
        mostrarError('categoria-nombre-error', 'El nombre de la categoría es obligatorio');
    } else if (valor.length < 2) {
        mostrarError('categoria-nombre-error', 'El nombre debe tener al menos 2 caracteres');
    } else if (valor.length > 50) {
        mostrarError('categoria-nombre-error', 'El nombre no puede tener más de 50 caracteres');
    }
}

// Previsualizar imagen (desde URL o archivo)
function previsualizarImagen(inputElementOrUrl) {
    const previewContainer = document.getElementById('preview-container');
    if (!previewContainer) return;
    
    // Limpiar preview anterior
    previewContainer.innerHTML = '';
    
    if (typeof inputElementOrUrl === 'string') {
        // Es una URL
        const url = inputElementOrUrl;
        if (url && esUrlValida(url)) {
            const img = document.createElement('img');
            img.src = url;
            img.alt = 'Previsualización';
            img.style.cssText = 'max-width: 200px; max-height: 200px; border-radius: 0.5rem; box-shadow: var(--shadow); object-fit: contain;';
            img.onerror = function() {
                previewContainer.innerHTML = '<p style="color: var(--danger-color); font-size: 0.875rem;">No se pudo cargar la imagen</p>';
            };
            
            const label = document.createElement('p');
            label.textContent = 'Previsualización:';
            label.style.cssText = 'margin-bottom: 0.5rem; font-weight: 500; color: var(--text-secondary);';
            
            previewContainer.appendChild(label);
            previewContainer.appendChild(img);
        }
    } else {
        // Es un input file
        const input = inputElementOrUrl;
        const file = input.files[0];
        
        if (file) {
            // Validar tipo de archivo
            const validTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'];
            if (!validTypes.includes(file.type)) {
                mostrarError('imagen-error', 'Tipo de archivo no válido. Use JPG, PNG, GIF o WebP');
                input.value = '';
                return;
            }
            
            // Validar tamaño (máximo 5MB)
            if (file.size > 5 * 1024 * 1024) {
                mostrarError('imagen-error', 'La imagen es demasiado grande. Máximo 5MB');
                input.value = '';
                return;
            }
            
            const reader = new FileReader();
            reader.onload = function(e) {
                const img = document.createElement('img');
                img.src = e.target.result;
                img.alt = 'Previsualización';
                img.style.cssText = 'max-width: 200px; max-height: 200px; border-radius: 0.5rem; box-shadow: var(--shadow); object-fit: contain;';
                
                const label = document.createElement('p');
                label.textContent = 'Previsualización del archivo:';
                label.style.cssText = 'margin-bottom: 0.5rem; font-weight: 500; color: var(--text-secondary);';
                
                const info = document.createElement('small');
                info.textContent = `${file.name} (${(file.size / 1024 / 1024).toFixed(2)} MB)`;
                info.style.cssText = 'color: var(--text-secondary); display: block; margin-top: 0.5rem;';
                
                previewContainer.appendChild(label);
                previewContainer.appendChild(img);
                previewContainer.appendChild(info);
            };
            reader.readAsDataURL(file);
            
            // Limpiar el campo URL si se sube un archivo
            const urlInput = document.getElementById('imagen_url');
            if (urlInput) {
                urlInput.value = '';
            }
        }
    }
}

// Función para buscar productos (funcionalidad adicional)
function buscarProductos() {
    const termino = document.getElementById('busqueda').value.toLowerCase();
    const productos = document.querySelectorAll('.producto-card');
    
    productos.forEach(producto => {
        const nombre = producto.querySelector('h3').textContent.toLowerCase();
        const descripcion = producto.querySelector('.descripcion').textContent.toLowerCase();
        
        if (nombre.includes(termino) || descripcion.includes(termino)) {
            producto.style.display = 'block';
        } else {
            producto.style.display = 'none';
        }
    });
}

// Función para formatear precio mientras se escribe
function formatearPrecio(input) {
    let valor = input.value.replace(/[^\d.]/g, '');
    
    // Asegurar que solo hay un punto decimal
    const partes = valor.split('.');
    if (partes.length > 2) {
        valor = partes[0] + '.' + partes.slice(1).join('');
    }
    
    // Limitar decimales a 2 dígitos
    if (partes[1] && partes[1].length > 2) {
        valor = partes[0] + '.' + partes[1].substring(0, 2);
    }
    
    input.value = valor;
}

// Event listeners adicionales
document.addEventListener('DOMContentLoaded', function() {
    // Formatear precio en tiempo real
    const precioInput = document.getElementById('precio');
    if (precioInput) {
        precioInput.addEventListener('input', function() {
            formatearPrecio(this);
        });
    }
    
    // Autofocus en el primer campo del formulario
    const primerInput = document.querySelector('form input:not([type="hidden"])');
    if (primerInput) {
        primerInput.focus();
    }
    
    // Confirmación al salir si hay cambios sin guardar
    let formData = {};
    const form = document.querySelector('form');
    if (form) {
        // Guardar datos iniciales del formulario
        const inputs = form.querySelectorAll('input, textarea, select');
        inputs.forEach(input => {
            formData[input.name] = input.value;
        });
        
        // Verificar cambios antes de salir
        window.addEventListener('beforeunload', function(e) {
            let hayCambios = false;
            inputs.forEach(input => {
                if (formData[input.name] !== input.value) {
                    hayCambios = true;
                }
            });
            
            if (hayCambios) {
                e.preventDefault();
                e.returnValue = '';
            }
        });
        
        // Limpiar alerta de cambios al enviar el formulario
        form.addEventListener('submit', function() {
            window.removeEventListener('beforeunload', arguments.callee);
        });
    }
});

// Función para mostrar notificaciones (opcional)
function mostrarNotificacion(mensaje, tipo = 'success') {
    const notificacion = document.createElement('div');
    notificacion.className = `notificacion notificacion-${tipo}`;
    notificacion.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 1rem 1.5rem;
        background-color: ${tipo === 'success' ? 'var(--success-color)' : 'var(--danger-color)'};
        color: white;
        border-radius: 0.5rem;
        box-shadow: var(--shadow-lg);
        z-index: 1000;
        opacity: 0;
        transform: translateX(100%);
        transition: all 0.3s ease;
    `;
    notificacion.textContent = mensaje;
    
    document.body.appendChild(notificacion);
    
    // Mostrar notificación
    setTimeout(() => {
        notificacion.style.opacity = '1';
        notificacion.style.transform = 'translateX(0)';
    }, 100);
    
    // Ocultar notificación después de 3 segundos
    setTimeout(() => {
        notificacion.style.opacity = '0';
        notificacion.style.transform = 'translateX(100%)';
        setTimeout(() => {
            document.body.removeChild(notificacion);
        }, 300);
    }, 3000);
}