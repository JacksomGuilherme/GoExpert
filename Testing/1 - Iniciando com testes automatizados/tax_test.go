package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amout := 100.00
	expected := 5.0

	result := CalculateTax(amout)

	if expected != result {
		t.Fatalf("Expected %f but got %f", expected, result)
	}
}
