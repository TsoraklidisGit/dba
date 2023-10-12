package repository

import (
	"database/sql"
	"dba/models"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Repository() {
	db, err := sql.Open("sqlite3", "myarmy.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	err = initializeDatabase(db, "resource/create_tables.sql") // Provide the SQL file path.
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	err = populateDatabase(db)
	if err != nil {
		fmt.Println("Error populating the database:", err)
		return
	}
}

// InitializeDatabase creates the necessary tables in the database.
func initializeDatabase(db *sql.DB, sqlFile string) error {
	// Read the SQL script from the specified file
	sqlScript, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	// Execute the SQL script to create the tables
	_, err = db.Exec(string(sqlScript))
	if err != nil {
		return err
	}

	log.Println("Database initialization complete.")
	return nil
}

// PopulateDatabase reads and executes the sample_data.sql script to populate the database.
func populateDatabase(db *sql.DB) error {
	sampleDataSQL, err := os.ReadFile("resource/sample_data.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sampleDataSQL))
	if err != nil {
		return err
	}

	log.Println("Database populated with sample data.")
	return nil
}

// TO-DO
// AddUnitsToArmy adds a unit to an army with a specified element order.
func AddUnitsToArmy(db *sql.DB, armyID, unitID, elementOrder int) error {
	insertSQL := `
		INSERT INTO army_elements (army_id, unit_id, element_order)
		VALUES (?, ?, ?)
	`
	_, err := db.Exec(insertSQL, armyID, unitID, elementOrder)
	if err != nil {
		return err
	}
	return nil
}

func QueryUnits(armyID int) ([]models.Unit, error) {
	db, err := sql.Open("sqlite3", "myarmy.db")
	if err != nil {
		return nil, err
	}
	defer db.Close() // Close the database connection when done.

	querySQL, err := os.ReadFile("resource/query_units.sql")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(string(querySQL), armyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []models.Unit
	for rows.Next() {
		var unit models.Unit
		//var specialAttr sql.NullString // Use sql.NullString for special_attr

		err := rows.Scan(&unit.Name, &unit.CombatVsFoot, &unit.CombatVsMtn,
			&unit.GoodGoingMove, &unit.BadGoingMove, &unit.SpecialAttr, &unit.ElementOrder)
		if err != nil {
			return nil, err
		}

		//FIX-ME : NULL IS A PROBLEM
		// if specialAttr.Valid {
		// 	unit.SpecialAttr = specialAttr.String
		// } else {
		// 	unit.SpecialAttr = "" // Handle NULL values appropriately
		// }

		units = append(units, unit)

	}
	return units, nil
}
