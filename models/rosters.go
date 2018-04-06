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
	Player1_ID int
	Player2 string
	Player2_ID int
	Player3 string
	Player3_ID int
	Player4 string
	Player4_ID int
	Player5 string
	Player5_ID int
	Player6 string
	Player6_ID int
	Player7 string
	Player7_ID int
	Player8 string
	Player8_ID int
	Player9 string
	Player9_ID int
	Player10 string
	Player10_ID int
	Player11 string
	Player11_ID int
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
		CONCAT(p1.first_name, ' ', p1.last_name) as p1, p1.id as p1_id,
		CONCAT(p2.first_name, ' ', p2.last_name) as p2, p2.id as p2_id,
		CONCAT(p3.first_name, ' ', p3.last_name) as p3, p3.id as p3_id,
		CONCAT(p4.first_name, ' ', p4.last_name) as p4, p4.id as p4_id,
		CONCAT(p5.first_name, ' ', p5.last_name) as p5, p5.id as p5_id,
		CONCAT(p6.first_name, ' ', p6.last_name) as p6, p6.id as p6_id,
		CONCAT(p7.first_name, ' ', p7.last_name) as p7, p7.id as p7_id,
		CONCAT(p8.first_name, ' ', p8.last_name) as p8, p8.id as p8_id,
		CONCAT(p9.first_name, ' ', p9.last_name) as p9, p9.id as p9_id,
		CONCAT(p10.first_name, ' ', p10.last_name) as p10, p10.id as p10_id,
		CONCAT(p11.first_name, ' ', p11.last_name) as p11, p11.id as p11_id
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
				&record.Player1_ID, 
				&record.Player2, 
				&record.Player2_ID, 
				&record.Player3, 
				&record.Player3_ID, 
				&record.Player4, 
				&record.Player4_ID, 
				&record.Player5, 
				&record.Player5_ID, 
				&record.Player6, 
				&record.Player6_ID, 
				&record.Player7, 
				&record.Player7_ID, 
				&record.Player8, 
				&record.Player8_ID, 
				&record.Player9, 
				&record.Player9_ID, 
				&record.Player10, 
				&record.Player10_ID, 
				&record.Player11, 
				&record.Player11_ID, 
			)

		if err != nil {
			fmt.Printf("\nError in scanning next record: %s\n", err.Error())
		}else{
			results = append(results, record)
		}
	}

	return results
}
