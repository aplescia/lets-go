package sql

import (
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

//Opens a postgres connection using the SQL ORM library Gorm. Example connString
//format:
//postgresql://localhost:5432/denver_replica?sslmode=disable
func OpenConnection(connString string) (*gorm.DB, error) {
	conn, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Panic(err)
	}
	conn.SingularTable(true)
	return conn, err
}