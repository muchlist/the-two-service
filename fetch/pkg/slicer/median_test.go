package slicer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian(t *testing.T) {
	type test struct {
		input []float64
		want  float64
	}

	tests := []test{
		{input: []float64{1, 2, 5, 9}, want: 3.5},
		{input: []float64{2, 4, 6}, want: 4},
		{input: []float64{1, 1, 1, 5}, want: 1},
	}

	for _, tc := range tests {
		got := Median(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
