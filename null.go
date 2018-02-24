package week

import (
	"database/sql/driver"
)

// NullWeek is a nullable Week representation.
type NullWeek struct {
	Week  Week
	Valid bool
}

// NewNullWeek creates a new NullWeek.
func NewNullWeek(week Week, valid bool) NullWeek {
	return NullWeek{Week: week, Valid: valid}
}

// NullWeekFrom creates a new NullWeek that will always be valid.
func NullWeekFrom(week Week) NullWeek {
	return NewNullWeek(week, true)
}

// NullWeekFromPtr creates a new NullWeek that may be null if week is nil.
func NullWeekFromPtr(week *Week) NullWeek {
	if week == nil {
		return NullWeek{}
	}
	return NewNullWeek(*week, true)
}

// MarshalText implements the encoding TextMarshaler interface.
func (n NullWeek) MarshalText() ([]byte, error) {

	if !n.Valid {
		return []byte{}, nil
	}

	return n.Week.MarshalText()
}

// UnmarshalText implements the encoding TextUnmarshaler interface.
func (n *NullWeek) UnmarshalText(text []byte) error {
	str := string(text)

	if str == "" || str == "null" {
		n.Week, n.Valid = Week{}, false
		return nil
	}

	err := n.Week.UnmarshalText(text)
	n.Valid = err == nil

	return err
}

// MarshalJSON implements the json Marshaler interface.
func (n NullWeek) MarshalJSON() ([]byte, error) {

	if !n.Valid {
		return []byte("null"), nil
	}

	return n.Week.MarshalJSON()
}

// UnmarshalJSON implements the json Unmarshaler interface.
func (n *NullWeek) UnmarshalJSON(data []byte) error {
	str := string(data)

	if str == "null" {
		n.Week, n.Valid = Week{}, false
		return nil
	}

	err := n.Week.UnmarshalJSON(data)
	n.Valid = err == nil

	return err
}

// Scan implements the sql Scanner interface.
func (n *NullWeek) Scan(value interface{}) error {

	if value == nil {
		n.Week, n.Valid = Week{}, false
		return nil
	}

	err := n.Week.Scan(value)
	if err != nil {
		return err
	}

	n.Valid = true

	return nil
}

// Value implements the driver Valuer interface.
func (n NullWeek) Value() (driver.Value, error) {

	if !n.Valid {
		return nil, nil
	}

	return n.Week.Value()
}

// Ptr returns a pointer to this NullWeek's value, or a nil pointer it it is invalid.
func (n NullWeek) Ptr() *Week {

	if !n.Valid {
		return nil
	}

	return &n.Week
}

// IsZero returns true for invalid NullWeeks.
func (n NullWeek) IsZero() bool {
	return !n.Valid
}
