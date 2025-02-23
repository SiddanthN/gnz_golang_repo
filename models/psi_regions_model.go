package models

import (
    "gorm.io/gorm"
)

// PSIRegionSchema defines the PSI model
type PSIRegion struct {

    RegionID uint `gorm:"primaryKey;autoIncrement" json:"region_id"`
    Name string `gorm:"unique;not null" json:"name"`
    Longitude float64 `gorm:"not null;default:0" json:"longitude"`
    Latitude float64 `gorm:"not null;default:0" json:"latitude"`

}

// CreatePSIRegion inserts a new user into the database
func CreatePSIRegion(db *gorm.DB, region *PSIRegion) error {

    return db.Create(region).Error

}

// GetAllRegions retrieves all regions
func GetAllRegions(db *gorm.DB) ([]PSIRegion, error) {

    var regions []PSIRegion

    err := db.Find(&regions).Error
    return regions, err

}
