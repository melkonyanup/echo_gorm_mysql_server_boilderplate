package v1

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/helpers"
	"myapp/internal/shared/payloads"
	"net/http"
	"time"
)

// Signup godoc
// @Summary create user
// @Description creates a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user_info body payloads.SignUpPayload true "Sign up the user"
// @Success 201 {object} helpers.Response
// @Failure 400,422 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /sign-up [post]
func (h *Handler) signUp(c echo.Context) error {
	payload := new(payloads.SignUpPayload)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if  err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	err := h.services.User.SignUp(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, helpers.Res("user was created"))
}

// Signin godoc
// @Summary User SignIn
// @Tags users
// @Description logs user in
// @Accept  json
// @Produce  json
// @Param input body payloads.SignInPayload true "sign in the user"
// @Success 200 {object} helpers.Response
// @Failure 400,401 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /sign-in [post]
func (h *Handler) signIn(c echo.Context) error {
	payload := new(payloads.SignInPayload)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	signedToken, err := h.services.User.SignIn(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + signedToken,
		HttpOnly: true, // disabling JavaScript access to cookie
		Expires:  time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, helpers.Res("ok"))
}

// Profile godoc
// @Security UserAuth
// @Summary User Profile
// @Tags users
// @Description gets info about the user
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400,401,404 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /user/profile [get]
func (h *Handler) getUserProfile(c echo.Context) error {
	email := c.Get("email") // from the authorization middleware
	user, err := h.services.User.GetUserProfile(email.(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
