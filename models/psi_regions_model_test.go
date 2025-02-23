package models

import (
    "testing"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// setupTestDB initializes an in-memory SQLite database
func setupTestDB(t *testing.T) *gorm.DB {

    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

    if err != nil {

        t.Fatalf("Failed to connect to test database: %v", err)

    }

    // Auto-migrate schema
    if err := db.AutoMigrate(&PSIRegionSchema{}); err != nil {

        t.Fatalf("Failed to migrate test database: %v", err)

    }

    return db

}

// TestUserOperations performs multiple test cases using table-driven tests
func TestUserOperations(t *testing.T) {

    db := setupTestDB(t)

    // Table-driven test cases
    tests := []struct {
        testTitle     string
        inputRegion   PSIRegionSchema
        expectedError bool
    }{
        {"Create valid region record", PSIRegionSchema{Name: "North", Longitude: 9.543, Latitude: 149.441}, false},
        {"Create second valid region record", PSIRegionSchema{Name: "South", Longitude: 9.543, Latitude: 149.441}, false},
        {"Create duplicate region name", PSIRegionSchema{Name: "North", Longitude: 9.543, Latitude: 149.441}, true},
        {"Create record with NULL/NIL value for longitude", PSIRegionSchema{Name: "East", Latitude: 149.441}, false},
        {"Create record with NULL/NIL value for latitude", PSIRegionSchema{Name: "West", Longitude: 9.543}, false},
    }

    for _, tt := range tests {

        t.Run(tt.testTitle, func(t *testing.T) {

            err := db.Create(&tt.inputRegion).Error

            if (err != nil) != tt.expectedError {
                t.Errorf("Test %s failed: expected error=%v, got %v", tt.testTitle, tt.expectedError, err)
            }

        })

    }

}

// TestGetUserByID uses table-driven tests for fetching users
func TestGetUserByID(t *testing.T) {

    db := setupTestDB(t)

    // Seed database with test users
    regions := []PSIRegionSchema{

        {Name: "North", Longitude: 9.543, Latitude: 149.441},
        {Name: "South", Longitude: 2.243, Latitude: 108.276},
        {Name: "West", Longitude: 4.987, Latitude: 127.112},
        {Name: "East", Longitude: 7.599, Latitude: 164.998},
        {Name: "Central", Longitude: 3.208, Latitude: 130.009},

    }

    for i := range regions {

        db.Create(&regions[i])

    }

    // Table-driven test cases
    tests := []struct {
        testTitle   string
        regionID    uint
        shouldExist bool
    }{
        {"Valid region record: South", regions[0].LocationID, true},
        {"Valid region: East", regions[1].LocationID, true},
        {"Non existent region: North-West", 997, false},
        {"Non existent region: East-Central", 279, false},
        {"Non existent region: West", 601, false},
        {"Valid region: West", regions[2].LocationID, true},
    }

    for _, tt := range tests {

        t.Run(tt.testTitle, func(t *testing.T) {

            var fetchedRegion PSIRegionSchema
            err := db.First(&fetchedRegion, tt.regionID).Error

            if (err == nil) != tt.shouldExist {
                t.Errorf("Test %s failed: expected existence=%v, got error %v", tt.testTitle, tt.shouldExist, err)
            }

        })

    }

}
