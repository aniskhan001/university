package delivery

import (
	"net/http"
	"strconv"
	"university/app/club/repo"
	"university/app/club/usecase"
	"university/errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// handler represent the http handler
type handler struct {
	Usecase usecase.Usecase
}

// RegisterEndpoints register all the listed endpoints with application server
func RegisterEndpoints(e *echo.Echo, db *gorm.DB) {
	h := &handler{
		Usecase: usecase.Init(
			repo.Init(db),
		),
	}

	e.GET("/clubs", h.List)
	e.POST("/club", h.Insert)
	e.PATCH("/club/:id", h.Edit)
	e.GET("/club/:id", h.GetByID)
}

// List return all items
func (h *handler) List(c echo.Context) error {
	resp, err := h.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// GetByID returns single item by ID
func (h *handler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}
	resp, err := h.Usecase.Get(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// Insert a single item into the system
func (h *handler) Insert(c echo.Context) error {
	resp, err := h.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// Edit a single item
func (h *handler) Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}
	resp, err := h.Usecase.Edit(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
