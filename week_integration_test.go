// +build integration

package week

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// link to pq for testing
	_ "github.com/lib/pq"
)

func TestIntegrationWeek_Scan(t *testing.T) {
	const query = "SELECT week FROM test_table ORDER BY week LIMIT 1"

	db := setupTestDB(t)
	row := db.QueryRow(query)

	var week Week
	err := row.Scan(&week)

	require.NoError(t, err)
	assert.Equal(t, Week{year: 2016, week: 16}, week)
}

func TestIntegrationWeek_Value(t *testing.T) {
	const insert = "INSERT INTO test_table (week, null_week) VALUES ($1, NULL)"

	db := setupTestDB(t)

	week, err := New(2018, 18)
	require.NoError(t, err)

	result, err := db.Exec(insert, week)
	require.NoError(t, err)

	affected, err := result.RowsAffected()
	assert.Equal(t, int64(1), affected)
}

func setupTestDB(t *testing.T) *sql.DB {
	const template = "host=%s port=%s dbname=%s user=%s password=%s sslmode=%s"
	const setup = `
		DROP TABLE IF EXISTS test_table CASCADE;
		CREATE TABLE test_table (week CHAR(8) NOT NULL, null_week CHAR(8) NULL);
		INSERT INTO test_table (week, null_week) VALUES ('2016-W16', NULL);`

	open := fmt.Sprintf(template, env("PGHOST", "localhost"), env("PGPORT", "5432"), env("PGDB", "postgres"),
		env("PGUSER", "postgres"), env("PGPASSWORD", "postgres"), env("PGSSL", "disable"))

	db, err := sql.Open("postgres", open)
	require.NoError(t, err)

	_, err = db.Exec(setup)
	require.NoError(t, err)

	return db
}

func env(key, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}
