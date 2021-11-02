package matrixops

import (
	"net/http"

	"github.com/labstack/echo"
	matrixOpsModel "github.com/manuelhr/go-utils-service/modules/matrix-ops/model"
	matrixUtils "github.com/manuelhr/go-utils-service/modules/matrix-ops/utils"
)

type MatrizController struct {
}

func NewMatrizController() *MatrizController {
	return &MatrizController{}
}

func (uc MatrizController) Rotar90GradosAntiHorario(c echo.Context) error {

	request := matrixOpsModel.MatrizOpRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	response := matrixOpsModel.MatrizOpResponse{}
	result, utilsError := matrixUtils.Rotar90GradosAntihorario(request.Input)
	if utilsError != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utilsError.Error())
	}
	response.Output = result
	c.JSON(http.StatusOK, response)
	return nil
}
