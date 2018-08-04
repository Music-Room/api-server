package psql

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TypePSQL struct {
	PGhost string
	PGport string
	PGname string
	PGuser string
	PGpass string
}


func makeMigrations(connection *gorm.DB) {
	migrate := os.Getenv("MIGRATE")

	if migrate == "1" {
		fmt.Println("Migrate")

		//TODO: ADD CONSTRAINTS
		connection.AutoMigrate()

		fmt.Println("Migrations done")
	}
}

func (p *TypePSQL) NewConn() *gorm.DB  {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
				p.PGhost, p.PGport, p.PGname, p.PGuser, p.PGpass)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = connection.DB().Ping()
	if err != nil {
		panic(err)
	}

	//makeMigrations(connection)

	return connection
}


func New() *TypePSQL {
	return &TypePSQL{
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGNAME"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASS"),
	}
}
