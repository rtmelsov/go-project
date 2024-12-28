package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	tests := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: SuccessBody},
	}
	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {

			request := httptest.NewRequest(test.method, "/", nil)
			w := httptest.NewRecorder()

			webhook(w, request)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.expectedCode, res.StatusCode)

			resBody, err := io.ReadAll(res.Body)
			assert.NoError(t, err)
			if test.expectedBody != "" {
				assert.JSONEq(t, test.expectedBody, string(resBody))
			}
		})
	}
}
