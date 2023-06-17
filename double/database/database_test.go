package database

import (
	"errors"
	"testing"
)

type fakeDB struct {
	Database
	err error
}

func (f *fakeDB) Insert(collection string, data interface{}) error {
	return f.err
}

func TestInsert(t *testing.T) {
	fake := &fakeDB{
		err: nil,
	}

	err := Insert(fake, "product", `{}`)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestInsertError(t *testing.T) {
	e := errors.New("fake error")
	fake := &fakeDB{
		err: e,
	}

	err := Insert(fake, "product", `{}`)

	if err != nil {
		t.Error(err.Error())
	}
}
