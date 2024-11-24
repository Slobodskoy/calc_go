package api

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler_Calc(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid expression",
			requestBody:    `{"expression": "2 + 2"}`,
			expectedStatus: 200,
			expectedBody:   `{"result": "4.000000"}`,
		},
		{
			name:           "invalid json",
			requestBody:    `{"expression": invalid}`,
			expectedStatus: 400,
			expectedBody:   `{"error": "decode err: invalid character 'i' looking for beginning of value"}`,
		},
		{
			name:           "missing expression field",
			requestBody:    `{}`,
			expectedStatus: 400,
			expectedBody:   `{"error": "invalid expression"}`,
		},
		{
			name:           "invalid expression",
			requestBody:    `{"expression": "2 ++ 2"}`,
			expectedStatus: 400,
			expectedBody:   `{"error": "insufficient operands"}`,
		},
		{
			name:           "empty request body",
			requestBody:    ``,
			expectedStatus: 400,
			expectedBody:   `{"error": "decode err: EOF"}`,
		},
	}

	handler := &CalcHandler{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/calc", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.Calc(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, w.Body.String())
			}
		})
	}
}
