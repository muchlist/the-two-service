package slicer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMin(t *testing.T) {
	type test struct {
		input     []float64
		operation string
		want      float64
	}

	tests := []test{
		{input: []float64{10, 5, 1, 25}, operation: "max", want: 25},
		{input: []float64{10, 5, 1, 25}, operation: "min", want: 1},
		{input: []float64{10, 1.8, 6.2, 25}, operation: "max", want: 25},
		{input: []float64{10, 1.8, 6.2, 25}, operation: "min", want: 1.8},
	}

	for _, tc := range tests {
		if tc.operation == "max" {
			got := Max(tc.input)
			assert.Equal(t, tc.want, got, "maxTest failed")
		} else {
			got := Min(tc.input)
			assert.Equal(t, tc.want, got, "minTest failed")
		}

	}
}
