package postgres

import (
	"errors"
	"fmt"
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

//Opens a postgres connection using the SQL ORM library Gorm. Example connString
//format:
//postgresql://localhost:5432/denver_replica?sslmode=disable
func OpenPostgresConnection(connString string) (*gorm.DB, error) {
	conn, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Panic(err)
	}
	conn.SingularTable(true)
	return conn, err
}

//Use a given DB connection to query all of type desiredStruct from the database.
//Expects a slice.
func QueryAllFromDatabase(db *gorm.DB, desiredStruct interface{}) {
	if nil == desiredStruct {
		return
	}
	if reflect.TypeOf(desiredStruct).Kind() == reflect.Struct {
		db.Find(&desiredStruct)
	} else if reflect.TypeOf(desiredStruct).Kind() == reflect.Ptr {
		db.Find(desiredStruct)
	}
}

//Use a given DB connection to insert a pointer into a database. Uses the type of the passed
//pointer to automigrate the schema and insert it into the database.
func InsertIntoDatabase(db *gorm.DB, ptrToInsert interface{}) error {
	if nil == ptrToInsert || nil == db {
		return errors.New(fmt.Sprintf("input object(s) nil, %v %v", db, ptrToInsert))
	}
	if reflect.TypeOf(ptrToInsert).Kind() == reflect.Struct && reflect.ValueOf(ptrToInsert).CanAddr() {
		db.AutoMigrate(&ptrToInsert)
		db.Create(&ptrToInsert)
	} else if reflect.TypeOf(ptrToInsert).Kind() == reflect.Ptr {
		db.AutoMigrate(ptrToInsert)
		db.Create(ptrToInsert)
	} else {
		return errors.New("must pass in a pointer, passed in object is not addressable")
	}
	return nil
}
