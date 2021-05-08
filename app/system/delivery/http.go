package delivery

import (
	"net/http"
	"university/app/system/repo"
	"university/app/system/usecase"
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

	e.GET("/", h.Root)
	e.GET("/h34l7h", h.Health)
}

// Root will let you see what you can slash 🐲
func (h *handler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "CLEAN study brings CLEAN results!"})
}

// Health will let you know the heart beats ❤️
func (h *handler) Health(c echo.Context) error {
	resp, err := h.Usecase.GetHealth()
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
