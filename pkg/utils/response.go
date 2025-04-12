package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response structure standard pour les réponses API
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// Meta pour la pagination
type Meta struct {
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	LastPage int   `json:"last_page"`
}

// SuccessResponse retourne une réponse de succès
func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse retourne une réponse d'erreur
func ErrorResponse(c *fiber.Ctx, status int, message string, err string) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// PaginatedResponse retourne une réponse paginée
func PaginatedResponse(c *fiber.Ctx, status int, message string, data interface{}, meta Meta) error {
	return c.Status(status).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    &meta,
	})
}
