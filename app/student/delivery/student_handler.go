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

// studentHandler represent the httphandler
type studentHandler struct {
	Usecase usecase.StudentUsecase
}

// RegisterEndpoints register all the listed endpoints with application server
func RegisterEndpoints(e *echo.Echo, db *gorm.DB) {
	handler := &studentHandler{
		Usecase: usecase.NewStudentUsecase(
			repo.NewStudentRepository(db),
			deptRepo.NewDeptRepository(db),
		),
	}

	e.GET("/student/:id", handler.GetByID)
	e.GET("/student/:id/details", handler.GetDetailsByID)
	e.GET("/students", handler.List)
	// todo: move this to department domain?
	e.GET("/department/:id/students", handler.ListFromDept)
	e.POST("/student", handler.Insert)
	e.POST("/students", handler.InsertMany)
	e.PATCH("/student/:id", handler.Edit)
}

// List return all students
func (th *studentHandler) List(c echo.Context) error {
	items, err := th.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, items)
}

// ListFromDept return all students from a specific department
func (th *studentHandler) ListFromDept(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	items, err := th.Usecase.ListByDept(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, items)
}

// GetByID returns single student by ID
func (th *studentHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	resp, err := th.Usecase.Get(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// GetDetailsByID returns single department by ID
func (th *studentHandler) GetDetailsByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	resp, err := th.Usecase.GetDetails(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// Insert a single student into the system
func (th *studentHandler) Insert(c echo.Context) error {
	resp, err := th.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// InsertMany students into the system
func (th *studentHandler) InsertMany(c echo.Context) error {
	resp, err := th.Usecase.InsertMany(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// Edit a single student
func (th *studentHandler) Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(errors.RespondError(errors.ErrBadRequest))
	}

	resp, err := th.Usecase.Edit(c, uint(id))
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}
