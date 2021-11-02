package matrixops

import (
	"github.com/labstack/echo"
	util "github.com/manuelhr/go-utils-service/modules/matrix-ops/controller"
)

func RegisterModule(e *echo.Echo) {

	currentModuleGroup := e.Group("/matrix-ops")

	utilController := util.NewMatrizController()

	currentModuleGroup.Add("POST", "/rotate-90-ccw", utilController.Rotar90GradosAntiHorario)

}
