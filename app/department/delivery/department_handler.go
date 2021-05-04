package delivery

import (
	"net/http"
	"strconv"
	"university/app/department/repo"
	"university/app/department/usecase"
	"university/app/errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// deptHandler represent the httphandler
type deptHandler struct {
	Usecase usecase.DeptUsecase
}

// RegisterEndpoints register all the listed endpoints with application server
func RegisterEndpoints(e *echo.Echo, db *gorm.DB) {
	handler := &deptHandler{
		Usecase: usecase.NewDeptUsecase(
			repo.NewDeptRepository(db),
		),
	}

	e.GET("/list", handler.List)
	e.POST("/create", handler.Insert)
	e.PATCH("/edit/:id", handler.Edit)
	e.POST("/create-many", handler.InsertMany)
	e.GET("/get/:id", handler.GetByID)
}

// List return all departments
func (dh *deptHandler) List(c echo.Context) error {
	items, err := dh.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, items)
}

// GetByID returns single department by ID
func (dh *deptHandler) GetByID(c echo.Context) error {
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
func (dh *deptHandler) Insert(c echo.Context) error {
	resp, err := dh.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// InsertMany departments into the system
func (dh *deptHandler) InsertMany(c echo.Context) error {
	resp, err := dh.Usecase.InsertMany(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}

// Edit a single department
func (dh *deptHandler) Edit(c echo.Context) error {
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
