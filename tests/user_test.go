package tests

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"myapp/internal/models"
	"myapp/internal/validators"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func (s *APITestSuite) TestUserSignUp() {
	r := s.Require()
	e := echo.New()
	validators.InitValidator(e)
	s.handler.InitRouter(e)

	email, password, firstName, lastName := "bro55@ffffff.com", "12345678", "bro", "jjjj"
	signUpData := fmt.Sprintf(`{"email":"%s","password":"%s","first_name":"%s","last_name":"%s"}`,
		email, password, firstName, lastName)

	req, err := http.NewRequest(echo.POST, "/api/v1/sign-up", strings.NewReader(signUpData))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	r.Equal(http.StatusCreated, w.Result().StatusCode)

	// check if user is in db
	var userFromDB models.User
	result := s.db.Where("email = ?", email).Find(&userFromDB)
	s.NoError(result.Error)

	r.Equal(firstName, userFromDB.FirstName)
	r.Equal(lastName, userFromDB.LastName)
	r.Equal(email, userFromDB.Email)

	_, err = s.hasher.HashPassword(password)
	s.NoError(err)
	err = s.hasher.CheckPassword(userFromDB.Password, password)
	s.NoError(err)

	s.db.Unscoped().Delete(&userFromDB)
}

func (s *APITestSuite) TestUserSignIn() {
	r := s.Require()
	e := echo.New()
	validators.InitValidator(e)
	s.handler.InitRouter(e)

	email, password, firstName, lastName := "bro55@ffffff.com", "12345678", "bro", "jjjj"
	hashedPassword, err := s.hasher.HashPassword(password)
	s.NoError(err)

	testUser := &models.User{FirstName: firstName, LastName: lastName, Email: email, Password: hashedPassword}
	result := s.db.Create(&testUser)
	s.NoError(result.Error)

	signInData := fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)
	req, err := http.NewRequest(echo.POST, "/api/v1/sign-in", strings.NewReader(signInData))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	resp := w.Result()
	authCookie := resp.Cookies()[0]
	authCookieValueBeginning := authCookie.Value[:6]

	r.Equal(http.StatusOK, resp.StatusCode)
	r.Equal(authCookie.Name, "Authorization")
	r.Equal(authCookie.HttpOnly, true)
	r.Equal(authCookieValueBeginning, "Bearer")

	s.db.Unscoped().Delete(&testUser)
}

func (s *APITestSuite) TestUserProfile() {
	r := s.Require()
	e := echo.New()
	validators.InitValidator(e)
	s.handler.InitRouter(e)

	email, password, firstName, lastName := "bro3@ffffff.com", "12345678", "bro", "jjjj"
	hashedPassword, err := s.hasher.HashPassword(password)
	s.NoError(err)
	testUser := &models.User{FirstName: firstName, LastName: lastName, Email: email, Password: hashedPassword}
	result := s.db.Create(&testUser)
	s.NoError(result.Error)

	req, err := http.NewRequest(echo.GET, "/api/v1/user/profile", strings.NewReader(""))
	s.NoError(err)

	signedToken, err := s.tokenManager.GenerateToken(testUser.Email, testUser.ID)
	s.NoError(err)

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + signedToken,
		HttpOnly: true, // disabling JavaScript access to cookie
		Expires:  time.Now().Add(24 * time.Hour),
	}
	req.AddCookie(cookie)

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	resp := w.Result()
	respData, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var returnedUser models.User
	err = json.Unmarshal(respData, &returnedUser)
	s.NoError(err)

	r.Equal(http.StatusOK, resp.StatusCode)
	r.Equal(returnedUser.Password, "")
	r.Equal(returnedUser.Email, email)
	r.Equal(returnedUser.FirstName, firstName)
	r.Equal(returnedUser.LastName, lastName)

	s.db.Unscoped().Delete(&testUser)
}