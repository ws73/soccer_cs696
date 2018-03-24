package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// MatchSchedulesRecord represents the fields in the table
type MatchSchedulesRecord struct{
	ID int
	TeamAway string
	TeamHome string
	Date string
}

// type MatchSchedulesModel represents the model table type
type MatchSchedulesModel struct {
	db *sql.DB	
}

// NewMatchSchedulesModel returns a pointer to a NewMatchSchedulesModel type
func NewMatchSchedulesModel(db *sql.DB) *MatchSchedulesModel {
	return &MatchSchedulesModel{db}
}

// Get all match schedules in the database
func (m *MatchSchedulesModel) GetAllMatchSchedules() []MatchSchedulesRecord{

	var results []MatchSchedulesRecord

	rows, err := m.db.Query(`
			SELECT m.id, a.name AS team_away, b.name AS team_home, m.date 
			FROM match_schedules m
			JOIN teams a ON a.id = m.team_away
			JOIN teams b ON b.id = m.team_home
			ORDER BY m.date;`)
	
	defer rows.Close()

	for rows.Next() {

		var record MatchSchedulesRecord

    		err = rows.Scan(&record.ID, &record.TeamAway, &record.TeamHome, &record.Date)
		if err != nil {
			fmt.Printf("\nError in scanning next record: %s\n", err.Error())
		}else{
			results = append(results, record)
		}
	}

	return results
}
