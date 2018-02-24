package week

import "github.com/pkg/errors"

// checkYearAndWeek tests if a week and year are valid values for an ISO 8601 week date. If one of the
// provided values is invalid the function returns a detailed error.
func checkYearAndWeek(year, week int) error {

	if year < 0 || year > 9999 {
		return errors.Errorf("year must be between 0 and 9999 but was %d", year)
	}

	if week < 1 || week > 53 {
		return errors.Errorf("week must be between 1 and 53 but was %d", week)
	}

	return nil
}
