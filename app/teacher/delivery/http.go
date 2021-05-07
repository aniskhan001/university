package delivery

import (
	"net/http"
	"strconv"
	deptRepo "university/app/department/repo"
	"university/app/teacher/repo"
	"university/app/teacher/usecase"
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
			deptRepo.Init(db),
		),
	}

	e.GET("/teacher/:id", h.GetByID)
	e.GET("/teacher/:id/details", h.GetDetailsByID)
	e.GET("/teachers", h.List)
	// todo: move this to department domain?
	e.GET("/department/:id/teachers", h.ListFromDept)
	e.POST("/teacher", h.Insert)
	e.POST("/teachers", h.InsertMany)
	e.PATCH("/teacher/:id", h.Edit)
}

// List return all teachers
func (h *handler) List(c echo.Context) error {
	resp, err := h.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// ListFromDept return all teachers from a specific department
func (h *handler) ListFromDept(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	resp, err := h.Usecase.ListByDept(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// GetByID returns single teacher by ID
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

// GetDetailsByID returns single department by ID
func (h *handler) GetDetailsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	resp, err := h.Usecase.GetDetails(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// Insert a single teacher into the system
func (h *handler) Insert(c echo.Context) error {
	resp, err := h.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// InsertMany teachers into the system
func (h *handler) InsertMany(c echo.Context) error {
	resp, err := h.Usecase.InsertMany(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// Edit a single teacher
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
