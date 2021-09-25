package delivery

import (
	"net/http"
	"strconv"
	deptRepo "university/app/department/repo"
	"university/app/student/repo"
	"university/app/student/usecase"
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

	e.GET("/student/:id", h.GetByID)
	e.GET("/students", h.List)
	// todo: move this to department domain?
	e.GET("/department/:id/students", h.ListFromDept)
	e.POST("/student", h.Insert)
	e.POST("/students", h.InsertMany)
	e.PATCH("/student/:id", h.Edit)
}

// List return all students
func (h *handler) List(c echo.Context) error {
	resp, err := h.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// ListFromDept return all students from a specific department
func (h *handler) ListFromDept(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	items, err := h.Usecase.ListByDept(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, items)
}

// GetByID returns single student by ID
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

// Insert a single student into the system
func (h *handler) Insert(c echo.Context) error {
	resp, err := h.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// InsertMany students into the system
func (h *handler) InsertMany(c echo.Context) error {
	resp, err := h.Usecase.InsertMany(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// Edit a single student
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
