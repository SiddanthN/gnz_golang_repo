package models

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Helper function to create an in-memory database for testing
func setupTestArchiveDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&PSIRollupTableArchiveSchema{})

	return db, nil

}

// Helper function to insert test data
func insertTestArchiveData(db *gorm.DB) {

	testData := []PSIRollupTableArchiveSchema{
		{Region: "South", CreatedDate: "2025-02-22 09:44:52", Timestamp: 1740217492}, // Feb 20, 2025, 03:00 UTC
		{Region: "East", CreatedDate: "2025-02-20 06:00:00", Timestamp: 1708418400}, // Feb 20, 2025, 06:00 UTC
	}

	for _, record := range testData {

		db.Create(&record)

	}

}

// Test querying data on a single exact date value
func TestQueryArchiveExactDate(t *testing.T) {

	tests := []struct {
		testTitle    string
		dateTimeStr  string
		wantError    bool
	}{
		{
			testTitle:      "Valid Datetime value",
			dateTimeStr:   "2025-02-22 09:44:52",
			wantError: false,
		},
		{
			testTitle:      "Non-Existent Datetime value",
			dateTimeStr:   "2025-02-21 04:00:00",
			wantError: true,
		},
	}

	db, err := setupTestArchiveDB()

	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	insertTestArchiveData(db)

	// Reference datetime object for formatting
	// Golang weirldy does not recognize a placeholder format
	placeholder_value := "2006-01-02 15:04:05"

	for _, tt := range tests {

		t.Run(tt.testTitle, func(t *testing.T) {

			dateTime, _ := time.Parse(placeholder_value, tt.dateTimeStr)
			timestamp := dateTime.Unix()

			var result PSIRollupTableArchiveSchema
			err := db.Where("timestamp = ?", timestamp).First(&result).Error

			if (err != nil) != tt.wantError {
				t.Errorf("Test %s failed: expected error = %v, got error = %v", tt.testTitle, tt.wantError, err)
			}

		})

	}

}


// Test querying data for a missing date and finding the closest match
func TestQueryArchiveClosestDate(t *testing.T) {

	db, err := setupTestArchiveDB()

	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}

	insertTestArchiveData(db)

	tests := []struct {
		testTitle    string
		dateTimeStr  string
	}{
		{
			testTitle:    "Closest Timestamp",
			dateTimeStr: "2025-02-20 04:00:00",
		},
	}

	placeholder_value:= "2006-01-02 15:04:05"

	for _, tt := range tests {

		t.Run(tt.testTitle, func(t *testing.T) {

			dateTime, _ := time.Parse(placeholder_value, tt.dateTimeStr)
			queryTimestamp := dateTime.Unix()

			var closestRecord PSIRollupTableArchiveSchema
			err := db.Raw(`
				SELECT * FROM psi_rollup_table_archive_schemas 
				ORDER BY ABS(timestamp - ?) LIMIT 1`, queryTimestamp).Scan(&closestRecord).Error

			if err != nil {

				t.Errorf("Failed to find closest timestamp: %v", err)

			} else if closestRecord.Timestamp == 0 {

				t.Errorf("Expected to find the closest timestamp but found none")

			} else {

				fmt.Printf("Closest timestamp to %d is %d\n", queryTimestamp, closestRecord.Timestamp)

			}

		})

	}

}

