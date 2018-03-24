package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// CoachRecord represents the fields in the table
type CoachRecord struct{
	id int
	first_name string
	last_name string
	middle_name string
        date_of_birth string
}

// type CoachModel represents the model table type
type CoachModel struct {
	db *sql.DB	
}

// NewCoachchModel returns a pointer to a CoachModel type
func NewCoachModel(db *sql.DB) *CoachModel {
	return &CoachModel{db}
}

// Get all matches in the database
func (m *CoachModel) GetAllCoaches() []CoachRecord{

	var results []CoachRecord

	rows, err := m.db.Query("SELECT * FROM coaches")	
	defer rows.Close()

	for rows.Next() {
		var record CoachRecord
    		err = rows.Scan(&record.id, &record.first_name, &record.last_name, &record.middle_name, &record.date_of_birth)
		if err != nil {
			fmt.Printf("\nError in scanning next record: %s\n", err.Error())
		}else{
			results = append(results, record)
		}
	}

	return results
}
