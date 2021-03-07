package delivery

import (
	"net/http"
	"university/app/errors"
	"university/app/system/usecase"

	"github.com/labstack/echo/v4"
)

// SystemHandler  represent the httphandler for order
type SystemHandler struct {
	Usecase usecase.SystemUsecase
}

// NewSystemHandler will initialize the orders/ resources endpoint
func NewSystemHandler(e *echo.Echo, us usecase.SystemUsecase) {
	handler := &SystemHandler{
		Usecase: us,
	}
	e.GET("/", handler.Root)
	e.GET("/h34l7h", handler.Health)
}

// Root will let you see what you can slash 🐲
func (sh *SystemHandler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Kaj korilei khadyo mele!"})
}

// Health will let you know the heart beats ❤️
func (sh *SystemHandler) Health(c echo.Context) error {
	resp, err := sh.Usecase.GetHealth()
	if err != nil {
		return c.JSON(errors.RespondError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
