package handlers

import (
	"crud-echo/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// Handler struct will hold the database connection
type Handler struct {
	DB *gorm.DB
}

// HomeHandler returns a handler function with access to the Handler struct
func (db *Handler) HomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := types.SampleResponse{
			Message: "This is a sample response for home handler yes baby",
			Status:  http.StatusOK,
		}

		// Use h.DB to interact with the database
		// Example: result := h.DB.Find(...)
		return c.JSON(http.StatusOK, response)
	}
}
