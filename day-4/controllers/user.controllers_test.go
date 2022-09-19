package controllers

import (
	"alterra-agmc-day2/config"
	"alterra-agmc-day2/lib/database"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	libUser = database.InitUser(connection)
	u       = InitUser(libUser)
)

type UserTestSuite struct {
	suite.Suite
	Echo *echo.Echo
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	config.InitDB()
	connection.Exec("TRUNCATE TABLE users")

	s.Echo = echo.New()
	s.Echo.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}

func (s *UserTestSuite) TestGetAllUsers() {
	path := "/jwt/users"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_get_all_users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":[{`,
			token:                jToken,
		},
		{
			name:                 "failed_get_all_users",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.GET(path, u.GetUsersControllers)

			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *UserTestSuite) TestGetUserById() {
	path := "/jwt/users/"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
		withCreatedUser      bool
		userId               string
	}{
		{
			name:                 "success_get_user_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":{`,
			token:                jToken,
			withCreatedUser:      true,
			userId:               "1",
		},
		{
			name:                 "failed_get_user_by_id_not_found",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: `{"msg":"record not found`,
			token:                jToken,
			userId:               "13",
		},
		{
			name:                 "failed_get_user_by_id_false_id",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: `{"msg":"strconv`,
			token:                jToken,
			userId:               "a",
		},
		{
			name:                 "failed_get_user_by_id_unauthorized",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
			userId:               "1",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.withCreatedUser {
				connection.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c');")
			}

			s.Echo.GET(path+":id", u.GetUserByIdControllers)

			req := httptest.NewRequest(http.MethodGet, path+tc.userId, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			log.Println(body)
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *UserTestSuite) TestCreateUser() {
	path := "/users"

	var testCases = []struct {
		name                 string
		expectStatus         int
		payload              string
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_create_user",
			expectStatus:         http.StatusOK,
			payload:              `"name":"John","email":"a@b.com","password":"12345678"`,
			expectBodyStartsWith: `{"msg":`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_required_name",
			expectStatus:         http.StatusBadRequest,
			payload:              `"email":"a@b.com","password":"12345678"`,
			expectBodyStartsWith: `{"msg":"Name is required`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_required_email",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","password":"12345678"`,
			expectBodyStartsWith: `{"msg":"Email is required`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_required_password",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","email":"a@b.com"`,
			expectBodyStartsWith: `{"msg":"Password is required`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_false_email",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","email":"a.b.com","password":"12345678"`,
			expectBodyStartsWith: `{"msg":"Email is not valid email`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_unauthorized",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.POST(path, u.CreateUserControllers)

			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{"+tc.payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *UserTestSuite) TestUpdateUserById() {
	path := "/jwt/users/"
	payload := `"name":"John","email":"a@b.c","password":"1234"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
		withCreatedUser      bool
		userId               string
	}{
		{
			name:                 "success_update_user_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":`,
			token:                jToken,
			withCreatedUser:      true,
			userId:               "1",
		},
		{
			name:                 "failed_update_user_by_id_false_id",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: `{"msg":"strconv`,
			token:                jToken,
			userId:               "a",
		},
		{
			name:                 "failed_update_user_by_id_forbidden",
			expectStatus:         http.StatusForbidden,
			expectBodyStartsWith: `{"msg":"Forbidden`,
			token:                jToken,
			userId:               "2",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.withCreatedUser {
				connection.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c');")
			}

			s.Echo.PUT(path+":id", u.UpdateUserControllers)

			req := httptest.NewRequest(http.MethodPut, path+tc.userId, strings.NewReader("{"+payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			log.Println(body)
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *UserTestSuite) TestDeleteUserById() {
	path := "/jwt/users/"

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
		withCreatedUser      bool
		userId               string
	}{
		{
			name:                 "success_delete_user_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":"success delete user"}`,
			token:                jToken,
			withCreatedUser:      true,
			userId:               "1",
		},
		{
			name:                 "failed_delete_user_by_id_false_id",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: `{"msg":"strconv`,
			token:                jToken,
			userId:               "a",
		},
		{
			name:                 "failed_delete_user_by_id_forbidden",
			expectStatus:         http.StatusForbidden,
			expectBodyStartsWith: `{"msg":"Forbidden`,
			token:                jToken,
			userId:               "2",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.withCreatedUser {
				connection.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c')")
			}

			s.Echo.DELETE(path+":id", u.DeletedUserControllers)

			req := httptest.NewRequest(http.MethodDelete, path+tc.userId, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func TestLogin(t *testing.T) {
	path := "/login"
	payload := `"email":"stheven@gmail.com","password":"123123"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		payload              string
		expectBodyStartsWith string
	}{
		{
			name:                 "success_login",
			expectStatus:         http.StatusOK,
			payload:              payload,
			expectBodyStartsWith: `{"msg":"Success Login"`,
		},
		{
			name:                 "failed_login_not_found",
			expectStatus:         http.StatusBadRequest,
			payload:              `"email":"a.b.com","password":"12345678"`,
			expectBodyStartsWith: `{"msg":"record not found`,
		},
	}

	e := echo.New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e.POST(path, u.LoginUserControllers)

			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{"+tc.payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e.ServeHTTP(res, req)

			assert.Equal(t, tc.expectStatus, res.Code)
			body := res.Body.String()
			assert.True(t, strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}
