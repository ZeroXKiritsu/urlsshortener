package service

import (
	"errors"
	"testing"

	"github.com/ZeroXKiritsu/urlshortener/internal/repository"
	mock_repository "github.com/ZeroXKiritsu/urlshortener/internal/repository/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestServiceGetOriginal(t *testing.T) {
	type mockBehavior func(r *mock_repository.MockShortURL, shortURL string)

	testTable := []struct {
		name             string
		input            string
		mockBehavior     mockBehavior
		expectedOriginal string
		expectedError    error
	}{
		{
			name:  "OK",
			input: "K4kf_deoA1",
			mockBehavior: func(r *mock_repository.MockShortURL, shortURL string) {
				r.EXPECT().SearchShortURL(shortURL).Return("https://www.google.com/", nil)
			},
			expectedOriginal: "https://www.google.com/",
			expectedError:    nil,
		},
		{
			name:  "Repository error",
			input: "K4kf_deoA1",
			mockBehavior: func(r *mock_repository.MockShortURL, shortURL string) {
				r.EXPECT().SearchShortURL(shortURL).Return("", errors.New("Repository error"))
			},
			expectedOriginal: "",
			expectedError:    errors.New("Repository error"),
		},
		{
			name:  "Original URL does not exists",
			input: "K4kf_deoA1",
			mockBehavior: func(r *mock_repository.MockShortURL, shortURL string) {
				r.EXPECT().SearchShortURL(shortURL).Return("", nil)
			},
			expectedOriginal: "",
			expectedError:    nil,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mock := mock_repository.NewMockShortURL(c)
			testCase.mockBehavior(mock, testCase.input)

			repository := &repository.Repository{ShortURL: mock}
			service := NewService(repository)

			// perform
			original, error := service.GetOriginal(testCase.input)

			// check
			assert.Equal(t, testCase.expectedOriginal, original)
			assert.Equal(t, testCase.expectedError, error)
		})
	}
}
