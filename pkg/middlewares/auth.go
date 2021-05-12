package middlewares

import (
	"github.com/labstack/echo/v4"
	"myapp/pkg/auth"
	"net/http"
	"os"
	"strings"
)

// Authz validates token and authorizes users
func Authz(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// clientToken := c.Request().Header.Get("Authorization")
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Auth token was not passed in the cookie")
		}

		clientToken := cookie.Value
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Incorrect Format of Authorization Token")
		}

		jwtWrapper := auth.TokenManager{
			SecretKey: os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
			Issuer:    "AuthService",
		}
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		return next(c)
	}
}
