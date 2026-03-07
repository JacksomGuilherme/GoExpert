package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAndSaveTax(t *testing.T) {
	repository := &TaxRepositoryMock{}

	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateAndSaveTax(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateAndSaveTax(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
}
