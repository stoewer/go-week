package week

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckYearAndWeek(t *testing.T) {
	tests := []struct {
		Year  int
		Week  int
		Valid bool
	}{
		{Year: -1, Week: 1, Valid: false},
		{Year: 10000, Week: 1, Valid: false},
		{Year: 2003, Week: 0, Valid: false},
		{Year: 2003, Week: 1, Valid: true},
		{Year: 2003, Week: 52, Valid: true},
		{Year: 2003, Week: 53, Valid: false},
		{Year: 2004, Week: 0, Valid: false},
		{Year: 2004, Week: 1, Valid: true},
		{Year: 2004, Week: 52, Valid: true},
		{Year: 2004, Week: 53, Valid: true},
		{Year: 2004, Week: 54, Valid: false},
	}

	for _, tt := range tests {
		err := checkYearAndWeek(tt.Year, tt.Week)
		if tt.Valid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestWeeksInYear(t *testing.T) {
	longYears := []int{
		4, 9, 15, 20, 26, 32, 37, 43, 48, 54, 60, 65, 71, 76, 82, 88, 93, 99, 105, 111, 116,
		122, 128, 133, 139, 144, 150, 156, 161, 167, 172, 178, 184, 189, 195, 201, 207, 212,
		218, 224, 229, 235, 240, 246, 252, 257, 263, 268, 274, 280, 285, 291, 296, 303, 308,
		314, 320, 325, 331, 336, 342, 348, 353, 359, 364, 370, 376, 381, 387, 392, 398}

	for _, year := range longYears {
		assert.Equal(t, 52, weeksInYear(year-1))
		assert.Equal(t, 53, weeksInYear(year))
		assert.Equal(t, 52, weeksInYear(year+1))
	}
}

func TestIsLeapYear(t *testing.T) {
	leapYears := []int{
		1804, 1808, 1812, 1816, 1820, 1824, 1824, 1832, 1836, 1840, 1844, 1848, 1852, 1856,
		1860, 1864, 1868, 1872, 1876, 1880, 1884, 1888, 1892, 1896, 1904, 1908, 1912, 1916,
		1920, 1924, 1928, 1932, 1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964, 1968, 1972,
		1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012, 2016, 2020, 2024, 2028,
		2032, 2036, 2040, 2044, 2048, 2052, 2056, 2060, 2064, 2068, 2072, 2076, 2080, 2084}

	for _, year := range leapYears {
		assert.Equal(t, false, isLeapYear(year-1))
		assert.Equal(t, true, isLeapYear(year))
		assert.Equal(t, false, isLeapYear(year+1))
	}
}
