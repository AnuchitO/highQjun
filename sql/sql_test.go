package sql

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	lastInsertId int64
	rowsAffected int64
}

func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertId, nil
}

func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m, nil
}

func TestExecQuery(t *testing.T) {
	mock := &mockDB{
		rowsAffected: 3,
	}

	r, err := ExecQuery(mock, "query", "arg1", "arg2")

	if err != nil {
		t.Errorf("ExecQuery() returned an error: %v", err)
	}

	if r != 1 {
		t.Errorf("ExecQuery() returned %d, expected 1", r)
	}
}
