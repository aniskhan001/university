package delivery

import (
	"net/http"
	"strconv"
	"university/app/department/usecase"
	"university/app/errors"

	"github.com/labstack/echo/v4"
)

// DeptHandler represent the httphandler
type DeptHandler struct {
	Usecase usecase.DeptUsecase
}

// NewDeptHandler will initialize the endpoints for department domain
func NewDeptHandler(e *echo.Echo, us usecase.DeptUsecase) {
	handler := &DeptHandler{
		Usecase: us,
	}
	e.GET("/list", handler.List)
	e.POST("/create", handler.Insert)
	e.PATCH("/edit/:id", handler.Edit)
	e.POST("/create-many", handler.InsertMany)
	e.GET("/get/:id", handler.GetByID)
}

// List return all departments
func (dh *DeptHandler) List(c echo.Context) error {
	items, err := dh.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, items)
}

// GetByID returns single department by ID
func (dh *DeptHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}
	resp, err := dh.Usecase.Get(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// Insert a single department into the system
func (dh *DeptHandler) Insert(c echo.Context) error {
	resp, err := dh.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// InsertMany departments into the system
func (dh *DeptHandler) InsertMany(c echo.Context) error {
	resp, err := dh.Usecase.InsertMany(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// Edit a single department
func (dh *DeptHandler) Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}
	resp, err := dh.Usecase.Edit(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
