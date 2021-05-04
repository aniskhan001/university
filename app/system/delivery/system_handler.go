package delivery

import (
	"net/http"
	"university/app/errors"
	"university/app/system/repo"
	"university/app/system/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// SystemHandler represent the httphandler
type SystemHandler struct {
	Usecase usecase.SystemUsecase
}

func RegisterSystemEndpoints(e *echo.Echo, db *gorm.DB) {
	handler := &SystemHandler{
		Usecase: usecase.NewSystemUsecase(
			repo.NewSystemRepository(db),
		),
	}

	e.GET("/", handler.Root)
	e.GET("/h34l7h", handler.Health)
}

// Root will let you see what you can slash 🐲
func (sh *SystemHandler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "CLEAN study brings CLEAN results!"})
}

// Health will let you know the heart beats ❤️
func (sh *SystemHandler) Health(c echo.Context) error {
	resp, err := sh.Usecase.GetHealth()
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
