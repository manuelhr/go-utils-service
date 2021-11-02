package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"

	matrixops "github.com/manuelhr/go-utils-service/modules/matrix-ops"
)

func main() {
	e := echo.New()

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	ConfigurarServicios(e)

	//Registro de módulos de negocio
	matrixops.RegisterModule(e)
	//tasas.RegisterModule(e)
	//plazos.RegisterModule(e)
	//etc.RegisterModule(e)

	//Inicio de servicio
	e.Logger.Fatal(e.Start(fmt.Sprint(":", httpPort)))
	fmt.Println("Aplicacion iniciada y escuchando en el puerto ", httpPort)
}

func ConfigurarServicios(e *echo.Echo) {
	e.Use(middleware.Logger())
	//Custom logger for console
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}-${status} at > ${uri} < in ${response_time} - ${response_size} bytes\n",
	}))
}
