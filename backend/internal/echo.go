package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/light-planck/nemmy/backend/internal/models"
	"github.com/uptrace/bun"
)

func NewEcho(frontendURL string, db *bun.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontendURL},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	h := &Handler{DB: db}
	api := e.Group("/api")

	u := api.Group("/users")
	u.GET("/:username", h.GetUser)
	return e
}

type Handler struct {
	DB *bun.DB
}

func (h *Handler) GetUser(c echo.Context) error {
	p := c.Param("username")

	var u models.User
	err := h.DB.NewSelect().
		Model(&u).
		Where("username = ?", p).
		Scan(c.Request().Context())
	if err != nil {
		return err
	}

	if err := c.JSON(http.StatusOK, u); err != nil {
		return err
	}
	return nil
}
