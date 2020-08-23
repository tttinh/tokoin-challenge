package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBool(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Should be true", args{"true"}, true, false},
		{"Should be false", args{"false"}, false, false},
		{"Should be error", args{"x"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToBool(tt.args.value)
			assert.Equal(t, tt.wantErr, err != nil, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
