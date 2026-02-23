package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateImage(file *multipart.FileHeader) (bool, error) {
	if file.Size > maxImageSize {
		return false, fmt.Errorf("Avatar image must be less than %dMB", maxImageSize/(1024*1024))
	}

	f, err := file.Open()
	if err != nil {
		return false, fmt.Errorf("Failed to open avatar image: %v", err)
	}

	defer f.Close()
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return false, fmt.Errorf("Failed to read avatar image: %v", err)
	}

	contentType := http.DetectContentType(buffer)
	if !validImageContentTypes[contentType] {
		return false, fmt.Errorf("Avatar image must be a valid image type (JPEG, PNG, GIF)")
	}

	ext := getFileExtension(file.Filename)
	if !slices.Contains(validImage, ext) {
		return false, fmt.Errorf("Avatar image must be one of the following types: %s", strings.Join(validImage, ", "))
	}

	file.Filename = fmt.Sprintf("%s_%s", uuid.New().String(), file.Filename)
	return true, nil
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return ""
	}
	return "." + strings.ToLower(parts[len(parts)-1])
}

func SaveFile(file *multipart.FileHeader, uploadDir string, c *gin.Context) (string, error) {
	uploadPath := filepath.ToSlash(uploadDir)
	publicPath := strings.TrimPrefix(uploadPath, "uploads")
	publicPath = strings.TrimPrefix(publicPath, "/")

	publicURL := "http://localhost:8080/images/" + file.Filename
	if publicPath != "" {
		publicURL = "http://localhost:8080/images/" + publicPath + "/" + file.Filename
	}

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	dst := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}
	return publicURL, nil
}
