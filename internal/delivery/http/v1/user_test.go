package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"myapp/internal/helpers"
	"myapp/internal/service"
	"myapp/internal/service/mocks"
	"myapp/internal/shared/payloads"
	"myapp/internal/validators"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	testUser := &payloads.SignUpPayload{
		FirstName: "bro",
		LastName:  "jjjj",
		Email:     "bro22@ffffff.com",
		Password:  "12345678",
	}

	tests := []struct {
		testName     string
		expectations func(ctx context.Context, userMockService *mocks.User)
		input        string
		err          error
		code         int
	}{
		{
			testName: "valid",
			expectations: func(ctx context.Context, userMockService *mocks.User) {
				userMockService.On("SignUp", testUser).Return(nil)
			},
			input: `{"email": "bro22@ffffff.com", "password": "12345678", "first_name": "bro", "last_name": "jjjj"}`,
			code:  http.StatusCreated,
		},
		{
			testName:     "missing parameter",
			expectations: func(ctx context.Context, userMockService *mocks.User) {},
			input:        `{}`,
			err:          errors.New("code=422, message=code=500, message=Key: 'SignUpPayload.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag\nKey: 'SignUpPayload.LastName' Error:Field validation for 'LastName' failed on the 'required' tag\nKey: 'SignUpPayload.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'SignUpPayload.Password' Error:Field validation for 'Password' failed on the 'required' tag"),
			code:         http.StatusUnprocessableEntity,
		},
		{
			testName:     "bad request",
			expectations: func(ctx context.Context, userMockService *mocks.User) {},
			input:        `{some"}`,
			err:          errors.New("code=400, message=code=400, message=Syntax error: offset=2, error=invalid character 's' looking for beginning of object key string, internal=invalid character 's' looking for beginning of object key string"),
			code:         http.StatusBadRequest,
		},
		{
			testName: "service error",
			expectations: func(ctx context.Context, userMockService *mocks.User) {
				userMockService.On("SignUp", testUser).Return(errors.New("bad request"))
			},
			input: `{"email": "bro22@ffffff.com", "password": "12345678", "first_name": "bro", "last_name": "jjjj"}`,
			err:   errors.New("code=400, message=bad request"),
			code:  http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Logf("running %v", test.testName)

		// initialize the echo context to use for the test
		e := echo.New()
		validators.InitValidator(e)
		r, err := http.NewRequest(echo.POST, "/api/v1/sign-up", strings.NewReader(test.input))
		if err != nil {
			t.Fatal("could not create request")
		}
		r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		userMockService := &mocks.User{}

		test.expectations(ctx.Request().Context(), userMockService)

		handler := &Handler{services: &service.Service{User: userMockService}}

		err = handler.signUp(ctx)
		assert.Equal(t, test.err == nil, err == nil)

		if err != nil {
			if test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				t.Errorf("Expected no error, found: %s", err.Error())
			}
			assert.Equal(t, test.code, helpers.GetResCode(err))
		} else {
			assert.Equal(t, test.code, w.Code)
		}
		userMockService.AssertExpectations(t)
	}
}
