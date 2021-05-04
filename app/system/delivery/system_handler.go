package delivery

import (
	"net/http"
	"university/app/errors"
	"university/app/system/repo"
	"university/app/system/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// systemHandler represent the httphandler
type systemHandler struct {
	Usecase usecase.SystemUsecase
}

// RegisterEndpoints register all the listed endpoints with application server
func RegisterEndpoints(e *echo.Echo, db *gorm.DB) {
	handler := &systemHandler{
		Usecase: usecase.NewSystemUsecase(
			repo.NewSystemRepository(db),
		),
	}

	e.GET("/", handler.Root)
	e.GET("/h34l7h", handler.Health)
}

// Root will let you see what you can slash 🐲
func (sh *systemHandler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "CLEAN study brings CLEAN results!"})
}

// Health will let you know the heart beats ❤️
func (sh *systemHandler) Health(c echo.Context) error {
	resp, err := sh.Usecase.GetHealth()
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
