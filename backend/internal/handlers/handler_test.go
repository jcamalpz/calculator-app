// internal/handlers/handlers_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	h := NewHandler()
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedResult float64
	}{
		{
			name:           "valid addition",
			method:         "POST",
			body:           `{"a": 10, "b": 5}`,
			expectedStatus: http.StatusOK,
			expectedResult: 15,
		},
		{
			name:           "invalid method",
			method:         "GET",
			body:           `{"a": 10, "b": 5}`,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "invalid json",
			method:         "POST",
			body:           `{invalid}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/v1/calculate/add", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.Add(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var resp Response
				if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if resp.Result != tt.expectedResult {
					t.Errorf("expected result %v, got %v", tt.expectedResult, resp.Result)
				}
			}
		})
	}
}

func TestSubtractHandler(t *testing.T) {
	h := NewHandler()
	req := httptest.NewRequest("POST", "/api/v1/calculate/subtract", bytes.NewBufferString(`{"a": 10, "b": 5}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Subtract(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 5 {
		t.Errorf("expected result 5, got %v", resp.Result)
	}
}

func TestMultiplyHandler(t *testing.T) {
	h := NewHandler()
	req := httptest.NewRequest("POST", "/api/v1/calculate/multiply", bytes.NewBufferString(`{"a": 10, "b": 5}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Multiply(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 50 {
		t.Errorf("expected result 50, got %v", resp.Result)
	}
}

func TestDivideHandler(t *testing.T) {
	h := NewHandler()
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "valid division",
			body:           `{"a": 10, "b": 5}`,
			expectedStatus: http.StatusOK,
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "division by zero",
			body:           `{"a": 10, "b": 0}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/calculate/divide", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.Divide(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectError {
				var resp Response
				if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if resp.Result != tt.expectedResult {
					t.Errorf("expected result %v, got %v", tt.expectedResult, resp.Result)
				}
			}
		})
	}
}

func TestPowerHandler(t *testing.T) {
	h := NewHandler()
	req := httptest.NewRequest("POST", "/api/v1/calculate/power", bytes.NewBufferString(`{"a": 2, "b": 8}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Power(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 256 {
		t.Errorf("expected result 256, got %v", resp.Result)
	}
}

func TestSquareRootHandler(t *testing.T) {
	h := NewHandler()
	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "valid square root",
			body:           `{"a": 16}`,
			expectedStatus: http.StatusOK,
			expectedResult: 4,
			expectError:    false,
		},
		{
			name:           "negative number",
			body:           `{"a": -16}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/api/v1/calculate/sqrt", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			h.SquareRoot(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if !tt.expectError {
				var resp Response
				if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				if resp.Result != tt.expectedResult {
					t.Errorf("expected result %v, got %v", tt.expectedResult, resp.Result)
				}
			}
		})
	}
}

func TestPercentageHandler(t *testing.T) {
	h := NewHandler()
	req := httptest.NewRequest("POST", "/api/v1/calculate/percentage", bytes.NewBufferString(`{"a": 20, "b": 100}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Percentage(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Result != 20 {
		t.Errorf("expected result 20, got %v", resp.Result)
	}
}
