package models

// import (
//     "gorm.io/gorm"
// )

// // PSIRegionSchema defines the PSI model
// type PSIRegionSchema struct {

//     LocationID uint `gorm:"primaryKey;autoIncrement" json:"location_id"`
//     Name string `gorm:"unique;not null" json:"name"`
//     Longitude float64 `gorm:"not null;default:0" json:"longitude"`
//     Latitude float64 `gorm:"not null;default:0" json:"latitude"`

// }

// // CreatePSIRegion inserts a new user into the database
// func CreatePSIRegion(db *gorm.DB, region *PSIRegionSchema) error {

//     return db.Create(region).Error

// }

// // // GetPSIRegion retrieves a user by their ID
// // func GetPSIRegion(db *gorm.DB, id uint) (*PSIRegionSchema, error) {

// //     var region PSIRegionSchema

// //     err := db.First(&region, id).Error
// //     return &region, err

// // }

// // GetAllUsers retrieves all users
// func GetAllUsers(db *gorm.DB) ([]PSIRegionSchema, error) {

//     var regions []PSIRegionSchema

//     err := db.Find(&regions).Error
//     return regions, err

// }
