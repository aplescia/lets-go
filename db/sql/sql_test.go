package sql_test

import (
	"github.com/Chewy-Inc/lets-go/db/sql"
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
	conn, err := sql.OpenPostgresConnection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		t.Fatal(err)
	}
	_ = sql.InsertIntoDatabase(conn, &MyDummyStruct{ Name: "Ptr"})
	_ = sql.InsertIntoDatabase(conn, NewDummyStruct("Bobby"))
	bobert := MyDummyStruct{Name: "Hey!"}
	err = sql.InsertIntoDatabase(conn, bobert)
	assert.NotNil(t, err)
	if err != nil {
		t.Log(err)
	}
	err = sql.InsertIntoDatabase(conn, &bobert)
	var str []MyDummyStruct
	sql.QueryAllFromDatabase(conn, &str)
	t.Log(str)
}
