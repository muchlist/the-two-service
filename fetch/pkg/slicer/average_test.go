package slicer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	type test struct {
		input []float64
		want  float64
	}

	tests := []test{
		{input: []float64{10, 5, 1, 25}, want: 10.25},
		{input: []float64{10, 1.8, 6.2, 25}, want: 10.75},
	}

	for _, tc := range tests {
		got := Average(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
