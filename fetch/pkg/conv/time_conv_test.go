package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConv(t *testing.T) {
	tests := map[string]struct {
		input    string
		wantYear int
		wantDate int
	}{
		"space and slice": {input: "2022/05/16 19:42:29", wantYear: 2022, wantDate: 16},
		"space":           {input: "2022-05-17 12:12:21", wantYear: 2022, wantDate: 17},
		"with T":          {input: "2022-05-21T05:56:40.089Z", wantYear: 2022, wantDate: 21},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ParseDate(tc.input)
			year := got.Year()
			day := got.Day() // int(got.Weekday())
			assert.Equal(t, tc.wantYear, year, "wrong year")
			assert.Equal(t, tc.wantDate, day, "wrong day")
		})
	}
}
