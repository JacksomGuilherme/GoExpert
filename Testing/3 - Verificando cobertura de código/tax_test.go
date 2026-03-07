package tax

import (
	"testing"
)

func TestCalculateBatch(t *testing.T) {
	type calcTax struct {
		amount   float64
		expected float64
	}

	table := []calcTax{
		{0.0, 0.0},
		{600.0, 5.0},
		{500.0, 5.0},
		{1200.0, 10.0},
		{-500.0, 5.0},
		{999.0, 5.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Fatalf("Expected %f but got %f", item.expected, result)
		}
	}
}
