package delivery

import (
	"net/http"
	"strconv"
	"university/app/department/repo"
	"university/app/department/usecase"
	"university/errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// handler represent the httphandler
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

	e.GET("/depts", h.List)
	e.POST("/dept", h.Insert)
	e.PATCH("/dept/:id", h.Edit)
	e.POST("/depts", h.InsertMany)
	e.GET("/dept/:id", h.GetByID)
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

// InsertMany items into the system
func (h *handler) InsertMany(c echo.Context) error {
	resp, err := h.Usecase.InsertMany(c)
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
