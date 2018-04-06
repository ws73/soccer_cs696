package models

// import packages 
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// PlayerRecord represents the fields in the table
type PlayerRecord struct{
	ID int
	FirstName string
	LastName string
	MiddleName string
	Position string
        DateOfBirth string
}

// type PlayerModel represents the model table type
type PlayerModel struct {
	db *sql.DB	
}

// NewPlayerModel returns a pointer to a PlayerModel type
func NewPlayerModel(db *sql.DB) *PlayerModel {
	return &PlayerModel{db}
}

// GetPlayer returns a player
func (p PlayerModel) GetPlayer(id int64) (PlayerRecord, error) {

	var player PlayerRecord

	err := p.db.QueryRow("SELECT id, first_name, last_name, middle_name, position, date_of_birth FROM players WHERE id = ?", id).Scan(
		
		&player.ID,
		&player.FirstName,
		&player.LastName,
		&player.MiddleName,
		&player.Position,
		&player.DateOfBirth,
	)

	return player, err
}
