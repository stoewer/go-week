package week

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNewNullWeek(t *testing.T) {
	valid := NewNullWeek(Week{year: 2000, week: 1}, true)

	assert.Equal(t, Week{year: 2000, week: 1}, valid.Week)
	assert.True(t, valid.Valid)

	invalid := NewNullWeek(Week{}, false)

	assert.Equal(t, Week{}, invalid.Week)
	assert.False(t, invalid.Valid)
}

func TestNullWeekFrom(t *testing.T) {
	valid := NullWeekFrom(Week{year: 2000, week: 1})

	assert.Equal(t, Week{year: 2000, week: 1}, valid.Week)
	assert.True(t, valid.Valid)
}

func TestNullWeekFromPtr(t *testing.T) {
	valid := NullWeekFromPtr(&Week{year: 2000, week: 1})

	assert.Equal(t, Week{year: 2000, week: 1}, valid.Week)
	assert.True(t, valid.Valid)

	invalid := NullWeekFromPtr(nil)

	assert.Equal(t, Week{}, invalid.Week)
	assert.False(t, invalid.Valid)
}

func TestNullWeek_IsZero(t *testing.T) {
	valid := NullWeek{Week: Week{year: 2000, week: 1}, Valid: true}

	assert.False(t, valid.IsZero())

	invalid := NullWeek{}

	assert.True(t, invalid.IsZero())
}

func TestNullWeek_Ptr(t *testing.T) {
	valid := NullWeek{Week: Week{year: 2000, week: 1}, Valid: true}

	assert.NotNil(t, valid.Ptr())

	invalid := NullWeek{}

	assert.Nil(t, invalid.Ptr())
}

func TestNullWeek_MarshalText(t *testing.T) {

	tests := []struct {
		Week     NullWeek
		Expected string
		Error    bool
	}{
		{Week: NullWeek{Week: Week{year: 2000, week: 10}, Valid: true}, Expected: "2000-W10"},
		{Week: NullWeek{}, Expected: ""},
		{Week: NullWeek{Week: Week{}, Valid: true}, Error: true},
	}

	for _, tt := range tests {
		text, err := tt.Week.MarshalText()
		if tt.Error {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tt.Expected, string(text))
		}
	}
}

func TestNullWeek_UnmarshalText(t *testing.T) {

	tests := []struct {
		Value    string
		Expected NullWeek
		Error    bool
	}{
		{Value: "0001-W01", Expected: NullWeek{Week: Week{year: 1, week: 1}, Valid: true}},
		{Value: "", Expected: NullWeek{Week: Week{}, Valid: false}},
		{Value: "null", Expected: NullWeek{}},
		{Value: "9999-W99", Error: true},
	}

	for _, tt := range tests {
		var week NullWeek
		err := week.UnmarshalText([]byte(tt.Value))

		if tt.Error {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tt.Expected, week)
		}
	}
}

func TestNullWeek_MarshalJSON(t *testing.T) {

	tests := []struct {
		Week     NullWeek
		Expected string
		Error    bool
	}{
		{Week: NullWeek{Week: Week{year: 1, week: 1}, Valid: true}, Expected: `"0001-W01"`},
		{Week: NullWeek{Week: Week{year: 2001, week: 22}, Valid: true}, Expected: `"2001-W22"`},
		{Week: NullWeek{}, Expected: `null`},
		{Week: NullWeek{Week: Week{year: 2001, week: 99}, Valid: true}, Error: true},
	}

	t.Run("method call", func(t *testing.T) {
		for _, tt := range tests {
			result, err := tt.Week.MarshalJSON()

			if tt.Error {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.JSONEq(t, tt.Expected, string(result))
			}
		}
	})

	t.Run("marshal struct", func(t *testing.T) {
		const template = `{"Week":%s}`

		type testType struct {
			Week NullWeek
		}

		for _, tt := range tests {
			testStruct := testType{Week: tt.Week}
			result, err := json.Marshal(testStruct)

			if tt.Error {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.JSONEq(t, fmt.Sprintf(template, tt.Expected), string(result))
			}
		}
	})
}

func TestNullWeek_UnmarshalJSON(t *testing.T) {

	tests := []struct {
		Value    string
		Expected NullWeek
		Error    bool
	}{
		{Value: `"0001-W01"`, Expected: NullWeek{Week: Week{year: 1, week: 1}, Valid: true}},
		{Value: `"2001-W22"`, Expected: NullWeek{Week: Week{year: 2001, week: 22}, Valid: true}},
		{Value: `"9999-W52"`, Expected: NullWeek{Week: Week{year: 9999, week: 52}, Valid: true}},
		{Value: `null`, Expected: NullWeek{}},
		{Value: `"9999-W99"`, Error: true},
	}

	t.Run("method call", func(t *testing.T) {
		for _, tt := range tests {
			var week NullWeek
			err := week.UnmarshalJSON([]byte(tt.Value))

			if tt.Error {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.Expected, week)
			}
		}
	})

	t.Run("unmarshal struct", func(t *testing.T) {
		const template = `{"Week":%s}`

		type testType struct {
			Week NullWeek
		}

		for _, tt := range tests {
			value := fmt.Sprintf(template, tt.Value)

			var testStruct testType
			err := json.Unmarshal([]byte(value), &testStruct)

			if tt.Error {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, testType{Week: tt.Expected}, testStruct)
			}
		}
	})
}

func TestNullWeek_Value(t *testing.T) {

	tests := []struct {
		Week     NullWeek
		Expected interface{}
		Error    bool
	}{
		{Week: NullWeek{Week: Week{year: 1, week: 1}, Valid: true}, Expected: "0001-W01"},
		{Week: NullWeek{Week: Week{year: 2001, week: 22}, Valid: true}, Expected: "2001-W22"},
		{Week: NullWeek{Week: Week{year: 9999, week: 52}, Valid: true}, Expected: "9999-W52"},
		{Week: NullWeek{}, Expected: nil},
		{Week: NullWeek{Week: Week{year: -100, week: 22}, Valid: true}, Error: true},
	}

	for _, tt := range tests {
		result, err := tt.Week.Value()

		if tt.Error {
			assert.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tt.Expected, result)
		}
	}
}

func TestNullWeek_Scan(t *testing.T) {
	const query = `SELECT null_week FROM test_table ORDER BY null_week LIMIT 1`

	tests := []struct {
		Value    driver.Value
		Expected NullWeek
		Error    bool
	}{
		{Value: "0001-W01", Expected: NullWeek{Week: Week{year: 1, week: 1}, Valid: true}},
		{Value: "2001-W22", Expected: NullWeek{Week: Week{year: 2001, week: 22}, Valid: true}},
		{Value: []byte("9999-W52"), Expected: NullWeek{Week: Week{year: 9999, week: 52}, Valid: true}},
		{Value: nil, Expected: NullWeek{}},
		{Value: "9999-W99", Error: true},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)

		mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"null_week"}).AddRow(tt.Value))

		row := db.QueryRow(query)

		var week NullWeek
		err = row.Scan(&week)
		if tt.Error {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			assert.Equal(t, tt.Expected, week)
		}
	}
}
