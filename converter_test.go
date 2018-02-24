package week

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeISOWeekDate(t *testing.T) {
	tests := []struct {
		Year     int
		Week     int
		Expected string
		Error    bool
	}{
		{Year: 0, Week: 1, Expected: `0000-W01`},
		{Year: 9999, Week: 52, Expected: `9999-W52`},
		{Year: 1999, Week: 21, Expected: `1999-W21`},
		{Year: 1999, Week: 0, Error: true},
		{Year: -1, Week: 21, Error: true},
	}

	for _, tt := range tests {
		result, err := encodeISOWeekDate(tt.Year, tt.Week)

		if tt.Error {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tt.Expected, string(result))
		}
	}
}

func TestDecodeISOWeekDate(t *testing.T) {

	tests := []struct {
		Value        string
		ExpectedYear int
		ExpectedWeek int
		Error        bool
	}{
		{Value: `0000-W01`, ExpectedYear: 0, ExpectedWeek: 1},
		{Value: `9999-W52`, ExpectedYear: 9999, ExpectedWeek: 52},
		{Value: `1800-W11`, ExpectedYear: 1800, ExpectedWeek: 11},
		{Value: `0000W01`, ExpectedYear: 0, ExpectedWeek: 1},
		{Value: `9999W52`, ExpectedYear: 9999, ExpectedWeek: 52},
		{Value: `1800W11`, ExpectedYear: 1800, ExpectedWeek: 11},
		{Value: `18000-w11`, Error: true},
		{Value: `0000-W00`, Error: true},
		{Value: `weekdate`, Error: true},
		{Value: ``, Error: true},
		{Value: `-100-W-1`, Error: true},
	}

	for _, tt := range tests {
		year, week, err := decodeISOWeekDate([]byte(tt.Value))

		if tt.Error {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tt.ExpectedYear, year)
			assert.Equal(t, tt.ExpectedWeek, week)
		}
	}
}
