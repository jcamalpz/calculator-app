// internal/calculator/service_test.go
package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	s := NewService()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 5, 15},
		{"negative numbers", -10, -5, -15},
		{"mixed signs", 10, -5, 5},
		{"with decimals", 10.5, 5.3, 15.8},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.Add(tt.a, tt.b)
			if !almostEqual(result, tt.expected) {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	s := NewService()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 5, 5},
		{"negative numbers", -10, -5, -5},
		{"mixed signs", 10, -5, 15},
		{"with decimals", 10.5, 5.3, 5.2},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.Subtract(tt.a, tt.b)
			if !almostEqual(result, tt.expected) {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	s := NewService()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 5, 50},
		{"negative numbers", -10, -5, 50},
		{"mixed signs", 10, -5, -50},
		{"with decimals", 10.5, 2, 21},
		{"multiply by zero", 10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.Multiply(tt.a, tt.b)
			if !almostEqual(result, tt.expected) {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	s := NewService()
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"positive numbers", 10, 5, 2, false},
		{"negative numbers", -10, -5, 2, false},
		{"mixed signs", 10, -5, -2, false},
		{"with decimals", 10.5, 2, 5.25, false},
		{"division by zero", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := s.Divide(tt.a, tt.b)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Divide(%v, %v) expected error but got none", tt.a, tt.b)
				}
				if err != ErrDivisionByZero {
					t.Errorf("Divide(%v, %v) expected ErrDivisionByZero but got %v", tt.a, tt.b, err)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				}
				if !almostEqual(result, tt.expected) {
					t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

func TestPower(t *testing.T) {
	s := NewService()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive exponent", 2, 3, 8},
		{"zero exponent", 5, 0, 1},
		{"negative exponent", 2, -2, 0.25},
		{"fractional exponent", 4, 0.5, 2},
		{"large numbers", 2, 10, 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.Power(tt.a, tt.b)
			if !almostEqual(result, tt.expected) {
				t.Errorf("Power(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	s := NewService()
	tests := []struct {
		name      string
		a         float64
		expected  float64
		expectErr bool
	}{
		{"positive number", 16, 4, false},
		{"zero", 0, 0, false},
		{"decimal", 2.25, 1.5, false},
		{"large number", 144, 12, false},
		{"negative number", -16, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := s.SquareRoot(tt.a)
			if tt.expectErr {
				if err == nil {
					t.Errorf("SquareRoot(%v) expected error but got none", tt.a)
				}
				if err != ErrNegativeSqrt {
					t.Errorf("SquareRoot(%v) expected ErrNegativeSqrt but got %v", tt.a, err)
				}
			} else {
				if err != nil {
					t.Errorf("SquareRoot(%v) unexpected error: %v", tt.a, err)
				}
				if !almostEqual(result, tt.expected) {
					t.Errorf("SquareRoot(%v) = %v; want %v", tt.a, result, tt.expected)
				}
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	s := NewService()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"20% of 100", 20, 100, 20},
		{"50% of 200", 50, 200, 100},
		{"10% of 50", 10, 50, 5},
		{"0% of 100", 0, 100, 0},
		{"100% of 75", 100, 75, 75},
		{"150% of 100", 150, 100, 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.Percentage(tt.a, tt.b)
			if !almostEqual(result, tt.expected) {
				t.Errorf("Percentage(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// almostEqual checks if two floats are approximately equal (handles floating point precision)
func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}
