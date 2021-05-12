package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func RunServer(e *echo.Echo) error {
	port := ":" + viper.GetString("http.port")
	readTimeout, _ := strconv.Atoi(viper.GetString("http.readTimeout"))
	writeTimeoutInSeconds, _ := strconv.Atoi(viper.GetString("http.writeTimeout"))

	e.Server.ReadTimeout = time.Duration(readTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(writeTimeoutInSeconds) * time.Second

	return e.Start(port)
}

func StopServer(e *echo.Echo, ctx context.Context) error {
	return e.Shutdown(ctx)
}
