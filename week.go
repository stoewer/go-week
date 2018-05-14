// Package week provides a simple data type representing a week date as defined by ISO 8601.
package week

import (
	"database/sql/driver"
	"time"

	"github.com/pkg/errors"
)

// Week represents a week date as defined by ISO 8601. Week can be marshaled to and unmarshaled from
// numerous formats such as plain text or json.
type Week struct {
	year int
	week int
}

// New creates a new Week object from the specified year and week.
func New(year, week int) (Week, error) {

	err := checkYearAndWeek(year, week)
	if err != nil {
		return Week{}, err
	}

	return Week{year: year, week: week}, nil
}

// Year returns the year of the ISO week date.
func (w *Week) Year() int {
	return w.year
}

// Week returns the week of the ISO week date.
func (w *Week) Week() int {
	return w.week
}

// Next calculates and returns the next week. If the next week is invalid (year > 9999) the function
// returns an error.
func (w *Week) Next() (Week, error) {
	var year, week int

	if w.week >= weeksInYear(w.year) {
		year, week = w.year+1, 1
	} else {
		year, week = w.year, w.week+1
	}

	return New(year, week)
}

// Previous calculates and returns the previous week. If the previous week is invalid (year < 0) the
// function returns an error.
func (w *Week) Previous() (Week, error) {
	var year, week int

	if w.week <= 1 {
		year, week = w.year-1, weeksInYear(w.year-1)
	} else {
		year, week = w.year, w.week-1
	}

	return New(year, week)
}

// UnmarshalJSON implements json.Unmarshaler for Week.
func (w *Week) UnmarshalJSON(data []byte) error {

	if data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("unable to unmarshal json: string literal expected")
	}

	year, week, err := decodeISOWeekDate(data[1 : len(data)-1])
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal json")
	}

	w.year, w.week = year, week

	return nil
}

// MarshalJSON implements json.Marshaler for Week.
func (w Week) MarshalJSON() ([]byte, error) {

	raw, err := encodeISOWeekDate(w.year, w.week)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal json")
	}

	json := make([]byte, 0, len(raw)+2)

	json = append(json, '"')
	json = append(json, raw...)
	json = append(json, '"')

	return json, nil
}

// UnmarshalText implements TextUnmarshaler for Week.
func (w *Week) UnmarshalText(data []byte) error {

	year, week, err := decodeISOWeekDate(data)
	if err != nil {
		return errors.Wrap(err, "unable to unmarshal text")
	}

	w.year, w.week = year, week

	return nil
}

// MarshalText implements TextMarshaler for Week.
func (w Week) MarshalText() ([]byte, error) {

	text, err := encodeISOWeekDate(w.year, w.week)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal text")
	}

	return text, nil
}

// Value implements Valuer for Week.
func (w Week) Value() (driver.Value, error) {

	text, err := encodeISOWeekDate(w.year, w.week)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create value")
	}

	return driver.Value(text), nil
}

// Scan implements scanner for Week.
func (w *Week) Scan(src interface{}) error {

	var year int
	var week int
	var err error

	switch val := src.(type) {
	case string:
		year, week, err = decodeISOWeekDate([]byte(val))
	case []byte:
		year, week, err = decodeISOWeekDate(val)
	default:
		return errors.New("unable to scan value: incompatible type")
	}

	if err != nil {
		return errors.Wrap(err, "unable to scan value")
	}

	w.year, w.week = year, week

	return nil
}

// FromTime converts time.Time into a Week
func FromTime(t time.Time) Week {
	year, week := t.ISOWeek()
	return Week{year: year, week: week}
}

// Weekday represents day of the week in ISO: Monday is 1st and Sunday is 7th
type Weekday int

// Values of days of the week in ISO format
const (
	Monday Weekday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Time converts week to time.Time of the specified weekday
// Implementation based on the method on the ordinal day of the year and described here:
// https://en.wikipedia.org/wiki/ISO_week_date#Calculating_a_date_given_the_year,_week_number_and_weekday
func (w *Week) Time(weekday Weekday) time.Time {
	jan4th := time.Date(w.Year(), 1, 4, 0, 0, 0, 0, time.UTC)
	correction := int(convertGoWeekdayToISOWeekday(jan4th.Weekday())) + 3

	ordinal := w.Week()*7 + int(weekday) - correction
	year, ordinal := normalizeOrdinal(w.Year(), ordinal)

	return time.Date(year, 1, ordinal, 0, 0, 0, 0, time.UTC)
}

func normalizeOrdinal(year, ordinal int) (normalizedYear, normalizedOrdinal int) {
	daysInYear := 365
	if ordinal < 1 {
		if isLeapYear(year - 1) {
			daysInYear = 366
		}
		return year - 1, daysInYear + ordinal
	}

	if isLeapYear(year) {
		daysInYear = 366
	}
	if ordinal > daysInYear {
		return year + 1, ordinal - daysInYear
	}
	return year, ordinal
}

func convertGoWeekdayToISOWeekday(goWeekday time.Weekday) Weekday {
	if goWeekday == time.Sunday {
		return Sunday
	}
	return Weekday(int(goWeekday))
}
