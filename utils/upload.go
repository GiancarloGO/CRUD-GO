package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveUploadedFile guarda un archivo subido y retorna la ruta relativa
func SaveUploadedFile(fileHeader *multipart.FileHeader) (string, error) {
	// Verificar que el directorio uploads existe
	uploadDir := "static/uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return "", fmt.Errorf("error al crear directorio uploads: %v", err)
		}
	}

	// Validar tipo de archivo
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("tipo de archivo no permitido: %s", ext)
	}

	// Generar nombre único para evitar colisiones
	timestamp := time.Now().Unix()
	newFilename := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)
	
	// Limpiar el nombre de archivo de caracteres especiales
	newFilename = strings.ReplaceAll(newFilename, " ", "_")
	
	// Ruta completa del archivo
	filePath := filepath.Join(uploadDir, newFilename)

	// Abrir el archivo subido
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("error al abrir archivo subido: %v", err)
	}
	defer src.Close()

	// Crear el archivo de destino
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error al crear archivo de destino: %v", err)
	}
	defer dst.Close()

	// Copiar el contenido
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", fmt.Errorf("error al copiar archivo: %v", err)
	}

	// Retornar la ruta relativa para guardar en BD
	return "/" + strings.ReplaceAll(filePath, "\\", "/"), nil
}

// DeleteFile elimina un archivo del sistema de archivos
func DeleteFile(imagePath string) error {
	if imagePath == "" || strings.HasPrefix(imagePath, "http") {
		return nil // No es un archivo local
	}

	// Remover la barra inicial si existe
	if strings.HasPrefix(imagePath, "/") {
		imagePath = imagePath[1:]
	}

	// Verificar que el archivo existe
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return nil // El archivo ya no existe, no hay error
	}

	// Eliminar el archivo
	return os.Remove(imagePath)
}