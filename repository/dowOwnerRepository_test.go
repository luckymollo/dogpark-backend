package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDogOwnerById(t *testing.T) {
	db, mock, err := sqlmock.New()
	setDataBase(db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "age", "email"}).AddRow("name1", 18, "email1")
	mock.ExpectPrepare("select name, age, email from dog_owner").ExpectQuery().WithArgs(1).WillReturnRows(rows)
	dowOwner := FindDogOwnerById(1)

	expectedDogOwner := DogOwner{1, "name1", 18, "email1"}
	assert.Equal(t, expectedDogOwner, dowOwner)
}

func TestSaveDownOwner(t *testing.T) {
	db, mock, err := sqlmock.New()
	setDataBase(db)
	if err != nil {
		t.Fatalf("an error '#{err} was not expected")
	}

	defer db.Close()
	dogOwner := DogOwner{Name: "name1", Age: 18, Email: "email1"}
	mock.ExpectPrepare("insert into dog_owner").
		ExpectExec().WithArgs("name1", 18, "email1").WillReturnResult(sqlmock.NewResult(1, 1))

	id := SaveDownOwner(dogOwner)
	assert.Equal(t, int64(1), id)
}
