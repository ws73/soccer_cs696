package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// ScoreRecord represents the fields in the table
type ScoreRecord struct{
	TeamAway string
	TeamAwayScore int
	TeamHome string
	TeamHomeScore int
        Date string
}

// type ScoreModel represents the model table type
type ScoreModel struct {
	db *sql.DB	
}

// NewScoreModel returns a pointer to a ScoreModel type
func NewScoreModel(db *sql.DB) *ScoreModel {
	return &ScoreModel{db}
}

// Get all scores in the database
func (s *ScoreModel) GetAllScores() []ScoreRecord{

	var results []ScoreRecord

	rows, err := s.db.Query(`
			SELECT a.name AS team_away, k.team_away_score, 
			b.name As team_home, k.team_home_score, k.date
			FROM scores k
			JOIN teams a ON a.id = k.team_away
			JOIN teams b ON b.id = k.team_home
			ORDER BY k.date
		`)	

	defer rows.Close()

	for rows.Next() {
		var record ScoreRecord
    		err = rows.Scan(&record.TeamAway, 
				&record.TeamAwayScore, 
				&record.TeamHome, 
				&record.TeamHomeScore,
				&record.Date,)
		if err != nil {
			fmt.Printf("\nError in scanning next record: %s\n", err.Error())
		}else{
			results = append(results, record)
		}
	}

	return results
}
