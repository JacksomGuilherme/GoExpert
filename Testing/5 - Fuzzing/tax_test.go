package tax

import (
	"testing"
)

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-25.00, -500.00, 0.00, 2.00, 500.00, 999.00, 1000.00, 1500.00}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount < 0 && result != 0 {
			t.Fatalf("Received %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Fatalf("Received %f but expected 20", result)
		}
	})
}
