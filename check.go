package week

import (
	"time"

	"github.com/pkg/errors"
)

// checkYearAndWeek tests if a week and year are valid values for an ISO 8601 week date. If one of the
// provided values is invalid the function returns a detailed error.
func checkYearAndWeek(year, week int) error {

	if year < 0 || year > 9999 {
		return errors.Errorf("year must be between 0 and 9999 but was %d", year)
	}

	maxWeeks := weeksInYear(year)
	if week < 1 || week > maxWeeks {
		return errors.Errorf("week in %d must be between 1 and %d but was %d", year, maxWeeks, week)
	}

	return nil
}

func weeksInYear(year int) int {

	firstWeekday := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).Weekday()

	if firstWeekday == time.Thursday || (isLeapYear(year) && firstWeekday == time.Wednesday) {
		return 53
	}

	return 52
}

func isLeapYear(year int) bool {
	return (year%400 == 0 || year%100 != 0) && (year%4 == 0)
}
