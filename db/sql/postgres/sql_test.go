package postgres_test

import (
	"github.com/Chewy-Inc/lets-go/db/sql/postgres"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type MyDummyStruct struct{
	ID int `gorm:"default:'0'"`
	Name string
}

func NewDummyStruct(name string) *MyDummyStruct {
	d := MyDummyStruct{Name: name}
	return &d
}

func TestQueryAllFromDatabase(t *testing.T) {
	conn, err := postgres.OpenPostgresConnection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		t.Fatal(err)
	}
	_ = postgres.InsertIntoDatabase(conn, &MyDummyStruct{ Name: "Ptr"})
	_ = postgres.InsertIntoDatabase(conn, NewDummyStruct("Bobby"))
	bobert := MyDummyStruct{Name: "Hey!"}
	err = postgres.InsertIntoDatabase(conn, bobert)
	assert.NotNil(t, err)
	if err != nil {
		t.Log(err)
	}
	err = postgres.InsertIntoDatabase(conn, &bobert)
	var str []MyDummyStruct
	postgres.QueryAllFromDatabase(conn, &str)
	t.Log(str)
}
