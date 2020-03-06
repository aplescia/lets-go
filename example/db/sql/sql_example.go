// +build example

package main

import (
	"fmt"
	"github.com/Chewy-Inc/lets-go/db/sql/postgres"
	log "github.com/sirupsen/logrus"
)

var(
	thing, _ = postgres.OpenPostgresConnection("postgresql://localhost:5432/denver_replica?sslmode=disable")
)

func main() {
	rows, err := thing.Raw("select * from promotion;").Rows()
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()
	for rows.Next(){
		fmt.Println(rows)
	}
}
