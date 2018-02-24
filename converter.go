package week

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var (
	errWeekDateFormat = errors.New("week date must have the format 'YYYY-Www' or 'YYYYWww'")
	weekDateRegexp    = regexp.MustCompile(`^([0-9]{4})-?W([0-9]{2})$`)
	weekDateFormat    = "%04d-W%02d"
)

// decodeISOWeekDate converts ISO week string representations such as 'YYYY-Www' and 'YYYYWww' to
// year and week number. If the parsed year or week number is invalid an error will be returned.
func decodeISOWeekDate(data []byte) (int, int, error) {

	if !weekDateRegexp.Match(data) {
		return 0, 0, errWeekDateFormat
	}

	year, err := strconv.ParseInt(string(data[0:4]), 10, 32)
	if err != nil {
		return 0, 0, errors.Wrap(err, "unable to parse year")
	}

	week, err := strconv.ParseInt(string(data[len(data)-2:]), 10, 32)
	if err != nil {
		return 0, 0, errors.Wrap(err, "unable to parse week number")
	}

	err = checkYearAndWeek(int(year), int(week))
	if err != nil {
		return 0, 0, err
	}

	return int(year), int(week), nil
}

// encodeISOWeekDate converts an ISO year and week number to a string representation with the
// format 'YYYY-Www'. If the provided year or week number is invalid an error will be returned.
func encodeISOWeekDate(year, week int) ([]byte, error) {

	err := checkYearAndWeek(year, week)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf(weekDateFormat, year, week)), nil
}
