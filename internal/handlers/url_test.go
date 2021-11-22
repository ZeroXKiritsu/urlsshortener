package handlers

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/ZeroXKiritsu/urlshortener/structs"
	"github.com/ZeroXKiritsu/urlshortener/internal/service"
	mock_service "github.com/ZeroXKiritsu/urlshortener/internal/service/mocks"
)

func TestHandlersCreateShortURL(t *testing.T) {
	type mockBehavior func(s *mock_service.MockShortURL, url structs.Requests)

	testTable := []struct {
		name                string
		inputBody           string
		inputURL            structs.Requests
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"url": "https://www.google.com/"}`,
			inputURL: structs.Requests{
				URL: "https://www.google.com/",
			},
			mockBehavior: func(s *mock_service.MockShortURL, url structs.Requests) {
				s.EXPECT().Create(url).Return("localhost:8080/K4kf_deoA1", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"short_url":"localhost:8080/K4kf_deoA1"}`,
		},
		{
			name:                "Empty request",
			inputBody:           `{"url": ""}`,
			mockBehavior:        func(s *mock_service.MockShortURL, url structs.Requests) {},
			expectedStatusCode:  400,
			expectedRequestBody: "\"Invalid input\"",
		},
		{
			name:                "Bad field name",
			inputBody:           `{"ul": "https://www.google.com/"}`,
			mockBehavior:        func(s *mock_service.MockShortURL, url structs.Requests) {},
			expectedStatusCode:  400,
			expectedRequestBody: "\"Invalid input\"",
		},
		{
			name:      "Service error",
			inputBody: `{"url": "https://www.google.com/"}`,
			inputURL: structs.Requests{
				URL: "https://www.google.com/",
			},
			mockBehavior: func(s *mock_service.MockShortURL, url structs.Requests) {
				s.EXPECT().Create(url).Return("localhost:8080/K4kf_deoA1", errors.New("service error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "\"Internal server error\"",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mock := mock_service.NewMockShortURL(c)
			testCase.mockBehavior(mock, testCase.inputURL)

			services := &service.Service{ShortURL: mock}
			handler := NewHandler(services)

			// test server
			r := gin.New()
			r.POST("/", handler.createShortURL)

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(testCase.inputBody))

			// perform request
			r.ServeHTTP(w, req)

			// check
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}
}

func TestHandlerGetOriginal(t *testing.T) {
	type mockBehavior func(s *mock_service.MockShortURL, shortURL string)

	testTable := []struct {
		name                string
		inputParam          string
		shortURL            string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:       "OK",
			inputParam: "K4kf_deoA1",
			shortURL:   "K4kf_deoA1",
			mockBehavior: func(s *mock_service.MockShortURL, shortURL string) {
				s.EXPECT().GetOriginal(shortURL).Return("https://www.google.com/", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"original":"https://www.google.com/"}`,
		},
		{
			name:       "URL does not exist",
			inputParam: "K4kf_deoA1",
			shortURL:   "K4kf_deoA1",
			mockBehavior: func(s *mock_service.MockShortURL, shortURL string) {
				s.EXPECT().GetOriginal(shortURL).Return("", nil)
			},
			expectedStatusCode:  400,
			expectedRequestBody: "\"URL does not exist\"",
		},
		{
			name:       "Service error",
			inputParam: "K4kf_deoA1",
			shortURL:   "K4kf_deoA1",
			mockBehavior: func(s *mock_service.MockShortURL, shortURL string) {
				s.EXPECT().GetOriginal(shortURL).Return("", errors.New("service error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "\"Internal server error\"",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mock := mock_service.NewMockShortURL(c)
			testCase.mockBehavior(mock, testCase.shortURL)

			services := &service.Service{ShortURL: mock}
			handler := NewHandler(services)

			// test server
			r := gin.New()
			r.GET("/:short_url", handler.getOriginal)

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/"+testCase.inputParam, nil)

			// perform request
			r.ServeHTTP(w, req)

			// check
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}
}