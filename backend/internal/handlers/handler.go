// internal/handlers/handlers.go
package handlers

import (
	"calculator/internal/calculator"
	"encoding/json"
	"net/http"
)

// Handler manages HTTP requests
type Handler struct {
	calc *calculator.Service
}

// NewHandler creates a new HTTP handler
func NewHandler() *Handler {
	return &Handler{
		calc: calculator.NewService(),
	}
}

// BinaryRequest represents a request with two operands
type BinaryRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

// UnaryRequest represents a request with one operand
type UnaryRequest struct {
	A float64 `json:"a"`
}

// Response represents a calculation response
type Response struct {
	Result    float64 `json:"result"`
	Operation string  `json:"operation"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// sendJSON sends a JSON response
func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// sendError sends an error response
func sendError(w http.ResponseWriter, status int, message string) {
	sendJSON(w, status, ErrorResponse{Error: message})
}

// validateMethod checks if the HTTP method is POST
func validateMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "method not allowed")
		return false
	}
	return true
}

// decodeBinaryRequest decodes a binary operation request
func decodeBinaryRequest(w http.ResponseWriter, r *http.Request) (*BinaryRequest, bool) {
	var req BinaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "invalid request body")
		return nil, false
	}
	return &req, true
}

// decodeUnaryRequest decodes a unary operation request
func decodeUnaryRequest(w http.ResponseWriter, r *http.Request) (*UnaryRequest, bool) {
	var req UnaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "invalid request body")
		return nil, false
	}
	return &req, true
}

// Add handles addition endpoint
func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}

	result := h.calc.Add(req.A, req.B)
	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "addition",
	})
}

// Subtract handles subtraction endpoint
func (h *Handler) Subtract(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}

	result := h.calc.Subtract(req.A, req.B)
	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "subtraction",
	})
}

// Multiply handles multiplication endpoint
func (h *Handler) Multiply(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}

	result := h.calc.Multiply(req.A, req.B)
	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "multiplication",
	})
}

// Divide handles division endpoint
func (h *Handler) Divide(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}

	result, err := h.calc.Divide(req.A, req.B)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "division",
	})
}

// Power handles exponentiation endpoint
func (h *Handler) Power(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}

	result := h.calc.Power(req.A, req.B)
	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "power",
	})
}

// SquareRoot handles square root endpoint
func (h *Handler) SquareRoot(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeUnaryRequest(w, r)
	if !ok {
		return
	}

	result, err := h.calc.SquareRoot(req.A)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "sqrt",
	})
}

// Percentage handles percentage endpoint
func (h *Handler) Percentage(w http.ResponseWriter, r *http.Request) {
	if !validateMethod(w, r) {
		return
	}

	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}

	result := h.calc.Percentage(req.A, req.B)
	sendJSON(w, http.StatusOK, Response{
		Result:    result,
		Operation: "percentage",
	})
}
