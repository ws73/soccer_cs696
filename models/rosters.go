package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// RostersRecord represents the fields in the table
type RostersRecord struct{
	TeamName string
	HeadCoach string
	AssistantCoach string
	Player1 string
	Player2 string
	Player3 string
	Player4 string
	Player5 string
	Player6 string
	Player7 string
	Player8 string
	Player9 string
	Player10 string
	Player11 string
}

// type RostersModel represents the model table type
type RostersModel struct {
	db *sql.DB	
}

// NewRostersModel returns a pointer to a RostersModel type
func NewRostersModel(db *sql.DB) *RostersModel {
	return &RostersModel{db}
}

// Get all rosters in the database
func (m *RostersModel) GetAllRosters() []RostersRecord{

	var results []RostersRecord

	rows, err := m.db.Query(`
	SELECT 
		t.name as team_name,
		CONCAT(c.first_name, ' ', c.last_name) as head_coach,
		CONCAT(c2.first_name, ' ', c2.last_name) as assistant_coach,
		CONCAT(p1.first_name, ' ', p1.last_name) as p1,
		CONCAT(p2.first_name, ' ', p2.last_name) as p2,
		CONCAT(p3.first_name, ' ', p3.last_name) as p3,
		CONCAT(p4.first_name, ' ', p4.last_name) as p4,
		CONCAT(p5.first_name, ' ', p5.last_name) as p5,
		CONCAT(p6.first_name, ' ', p6.last_name) as p6,
		CONCAT(p7.first_name, ' ', p7.last_name) as p7,
		CONCAT(p8.first_name, ' ', p8.last_name) as p8,
		CONCAT(p9.first_name, ' ', p9.last_name) as p9,
		CONCAT(p10.first_name, ' ', p10.last_name) as p10,
		CONCAT(p11.first_name, ' ', p11.last_name) as p11
	FROM rosters r
	JOIN teams t ON r.team_id = t.id
	JOIN players p1 ON p1.id = r.player_1
	JOIN players p2 ON p2.id = r.player_2
	JOIN players p3 ON p3.id = r.player_3
	JOIN players p4 ON p4.id = r.player_4
	JOIN players p5 ON p5.id = r.player_5
	JOIN players p6 ON p6.id = r.player_6
	JOIN players p7 ON p7.id = r.player_7
	JOIN players p8 ON p8.id = r.player_8
	JOIN players p9 ON p9.id = r.player_9
	JOIN players p10 ON p10.id = r.player_10
	JOIN players p11 ON p11.id = r.player_11
	JOIN coaches c ON c.id = r.head_coach
	JOIN coaches c2 ON c2.id = r.assistant_coach`)	

	defer rows.Close()

	for rows.Next() {
		var record RostersRecord
		err = rows.Scan(
			&record.TeamName, 
			&record.HeadCoach, 
			&record.AssistantCoach, 
			&record.Player1, 
			&record.Player2, 
			&record.Player3, 
			&record.Player4, 
			&record.Player5, 
			&record.Player6, 
			&record.Player7, 
			&record.Player8, 
			&record.Player9, 
			&record.Player10, 
			&record.Player11, 
			)

		if err != nil {
			fmt.Printf("\nError in scanning next record: %s\n", err.Error())
		}else{
			results = append(results, record)
		}
	}

	return results
}
