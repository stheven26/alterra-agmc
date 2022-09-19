package controllers

import (
	"alterra-agmc-day2/config"
	"alterra-agmc-day2/lib/database"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

var (
	connection = config.InitDB()
	libBook    = database.InitBook(connection)
	b          = InitBook(libBook)
	jToken     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjU5MzkyMjMsInVzZXJJZCI6MX0.I_qMVlgghBUW-tr7yqVumMMZDOSmjlLui2gYa2tASzw"
)

type BookTestSuite struct {
	suite.Suite
	Echo *echo.Echo
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}

func (s *BookTestSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	config.InitDB()
	s.Echo = echo.New()
	s.Echo.Use(middleware.JWT([]byte(config.LoadEnv().GetString("JWT_KEY"))))
}

func (s *BookTestSuite) TestGetAllBooks() {
	path := "/books"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_get_all_books",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":[{`,
			token:                jToken,
		},
		{
			name:                 "failed_get_all_books",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.GET(path, b.GetAllBookControllers)

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

func (s *BookTestSuite) TestGetBookById() {
	path := "/books/2"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_get_book_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":{"id":1`,
			token:                jToken,
		},
		{
			name:                 "failed_get_book_by_id",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.GET(path, b.GetBookByIdControllers)

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

func (s *BookTestSuite) TestCreateBook() {
	path := "/jwt/books"
	payload := `"title":"GO","isbn":"x.x.x.x","writer":"Erlangga"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_create_book",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":{"id":3,` + payload,
			token:                jToken,
		},
		{
			name:                 "failed_create_book",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.POST(path, b.PostBookControllers)

			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{"+payload+"}"))
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

func (s *BookTestSuite) TestUpdateBookById() {
	path := "/jwt/books/2"
	payload := `"title":"GO","isbn":"x.x.x.x","writer":"Erlangga"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_update_book_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"msg":{"id":1,` + payload,
			token:                jToken,
		},
		{
			name:                 "failed_update_book_by_id",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"msg":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.PUT("/books/:id", b.UpdateBookControllers)

			req := httptest.NewRequest(http.MethodPut, path, strings.NewReader("{"+payload+"}"))
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

func (s *BookTestSuite) TestDeleteBookById() {
	path := "/jwt/books/1"

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_delete_book_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"status":"success delete book"}`,
			token:                jToken,
		},
		{
			name:                 "failed_delete_book_by_id",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"status":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.DELETE(path, b.DeleteBookControllers)

			req := httptest.NewRequest(http.MethodDelete, path, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}
