package delivery

import (
	"net/http"
	"strconv"
	"university/app/errors"
	"university/app/teacher/usecase"

	"github.com/labstack/echo/v4"
)

// TeacherHandler represent the httphandler
type TeacherHandler struct {
	Usecase usecase.TeacherUsecase
}

// NewTeacherHandler will initialize the endpoints for teacher domain
func NewTeacherHandler(e *echo.Echo, us usecase.TeacherUsecase) {
	handler := &TeacherHandler{
		Usecase: us,
	}
	e.GET("/teacher/get/:id", handler.GetByID)
	e.GET("/teacher/list", handler.List)
	e.POST("/teacher/create", handler.Insert)
	e.POST("/teacher/create-many", handler.InsertMany)
	e.PATCH("/teacher/edit/:id", handler.Edit)
}

// List return all teachers
func (th *TeacherHandler) List(c echo.Context) error {
	items, err := th.Usecase.List(c)
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

// Edit a single teacher into the system
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
