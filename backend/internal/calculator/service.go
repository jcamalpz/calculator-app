// internal/calculator/service.go
package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrNegativeSqrt   = errors.New("cannot calculate square root of negative number")
	ErrInvalidInput   = errors.New("invalid input")
)

// Service handles calculator business logic
type Service struct{}

// NewService creates a new calculator service
func NewService() *Service {
	return &Service{}
}

// Add performs addition of two numbers
func (s *Service) Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction of two numbers
func (s *Service) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication of two numbers
func (s *Service) Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division of two numbers
// Returns error if divisor is zero
func (s *Service) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// Power calculates a raised to the power of b
func (s *Service) Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// SquareRoot calculates the square root of a number
// Returns error if number is negative
func (s *Service) SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSqrt
	}
	return math.Sqrt(a), nil
}

// Percentage calculates a% of b
// Example: Percentage(20, 100) returns 20 (which is 20% of 100)
func (s *Service) Percentage(a, b float64) float64 {
	return (a / 100) * b
}
