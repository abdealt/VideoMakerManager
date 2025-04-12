package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Pagination structure pour les paramètres de pagination
type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

// GetPagination extraire les paramètres de pagination de la requête
func GetPagination(c *fiber.Ctx) Pagination {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	sort := c.Query("sort", "created_at desc")

	// Limiter la taille maximale de la page
	if limit > 100 {
		limit = 100
	}

	if page < 1 {
		page = 1
	}

	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

// Paginate applique la pagination à une requête GORM
func Paginate(db *gorm.DB, pagination Pagination) *gorm.DB {
	offset := (pagination.Page - 1) * pagination.Limit
	return db.Offset(offset).Limit(pagination.Limit).Order(pagination.Sort)
}
