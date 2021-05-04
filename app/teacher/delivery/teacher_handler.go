package delivery

import (
	"net/http"
	"strconv"
	deptRepo "university/app/department/repo"
	"university/app/errors"
	"university/app/teacher/repo"
	"university/app/teacher/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// TeacherHandler represent the httphandler
type TeacherHandler struct {
	Usecase usecase.TeacherUsecase
}

func RegisterTeacherEndpoints(e *echo.Echo, db *gorm.DB) {
	handler := &TeacherHandler{
		Usecase: usecase.NewTeacherUsecase(
			repo.NewTeacherRepository(db),
			deptRepo.NewDeptRepository(db),
		),
	}

	e.GET("/teacher/:id", handler.GetByID)
	e.GET("/teacher/:id/details", handler.GetDetailsByID)
	e.GET("/teachers", handler.List)
	// todo: move this to department domain?
	e.GET("/department/:id/teachers", handler.ListFromDept)
	e.POST("/teacher", handler.Insert)
	e.POST("/teachers", handler.InsertMany)
	e.PATCH("/teacher/:id", handler.Edit)
}

// List return all teachers
func (th *TeacherHandler) List(c echo.Context) error {
	items, err := th.Usecase.List(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, items)
}

// ListFromDept return all teachers from a specific department
func (th *TeacherHandler) ListFromDept(c echo.Context) error {
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

// GetByID returns single teacher by ID
func (th *TeacherHandler) GetByID(c echo.Context) error {
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
func (th *TeacherHandler) GetDetailsByID(c echo.Context) error {
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

// Insert a single teacher into the system
func (th *TeacherHandler) Insert(c echo.Context) error {
	resp, err := th.Usecase.Insert(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// InsertMany teachers into the system
func (th *TeacherHandler) InsertMany(c echo.Context) error {
	resp, err := th.Usecase.InsertMany(c)
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}

	return c.JSON(http.StatusOK, resp)
}

// Edit a single teacher
func (th *TeacherHandler) Edit(c echo.Context) error {
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
