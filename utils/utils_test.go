package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBool(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input         string
		expectedValue bool
		expectedErr   bool
	}{
		{"true", true, false},
		{"false", false, false},
		{"1", false, true},
	}

	for _, test := range tests {
		output, err := StringToBool(test.input)
		if test.expectedErr {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(output, test.expectedValue)
		}
	}
}
