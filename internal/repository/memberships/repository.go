package memberships

import (
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	rows, err := db.Query("SELECT id, email from fastcampus.users")
	if err != nil {
		log.Println("failed to query", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println("failed to scan", err)
		}
		log.Println(id, email)
	}

	return &repository{
		db: db,
	}
}
