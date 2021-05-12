package helpers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

func LoadConfig() error {
	_, b, _, _ := runtime.Caller(0)
	pathToCurrentFile := filepath.Dir(b)

	viper.AddConfigPath(pathToCurrentFile + "/../../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func LoadEnvVariables() error {
	_, b, _, _ := runtime.Caller(0)
	pathToCurrentFile := filepath.Dir(b)

	err := godotenv.Load(pathToCurrentFile + "/../../.env")
	if err != nil {
		return err
	}

	return nil
}

// HTTPCode returns the HTTP code of a given custom HTTP error, with 500 as default.
func GetResCode(err error) int {
	code := 500
	e, ok := err.(*echo.HTTPError)
	if ok {
		code = e.Code
	}
	return code
}