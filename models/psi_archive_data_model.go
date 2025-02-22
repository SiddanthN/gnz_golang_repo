package models

import "gorm.io/gorm"

type PSIRollupTableArchiveSchema struct {
    RecordID             uint    `gorm:"primaryKey;autoIncrement" json:"record_id"`
    Region               string  `gorm:"not null" json:"region"`
    CreatedDate          string  `gorm:"not null" json:"__date"`
    Timestamp            int     `gorm:"not null;index" json:"timestamp"`
    CoSubIndex           float64 `gorm:"not null;default:0" json:"co_sub_index"`
    So2TwentyFourHourly  float64 `gorm:"not null;default:0" json:"so2_twenty_four_hourly"`
    So2SubIndex          float64 `gorm:"not null;default:0" json:"so2_sub_index"`
    PsiThreeHourly       float64 `gorm:"not null;default:0" json:"psi_three_hourly"`
    CoEightHourMax       float64 `gorm:"not null;default:0" json:"co_eight_hour_max"`
    No2OneHourMax        float64 `gorm:"not null;default:0" json:"no2_one_hour_max"`
    Pm10SubIndex         float64 `gorm:"not null;default:0" json:"pm10_sub_index"`
    Pm25SubIndex         float64 `gorm:"not null;default:0" json:"pm25_sub_index"`
    O3EightHourMax       float64 `gorm:"not null;default:0" json:"o3_eight_hour_max"`
    PsiTwentyFourHourly  float64 `gorm:"not null;default:0" json:"psi_twenty_four_hourly"`
    O3SubIndex           float64 `gorm:"not null;default:0" json:"o3_sub_index"`
    Pm25TwentyFourHourly float64 `gorm:"not null;default:0" json:"pm25_twenty_four_hourly"`
    Pm10TwentyFourHourly float64 `gorm:"not null;default:0" json:"pm10_twenty_four_hourly"`
}


func CreatePSIArchiveRecord(db *gorm.DB, region *PSIRollupTableArchiveSchema) error {

    return db.Create(region).Error

}

// // GetPSIRegion retrieves a user by their ID
// func GetPSIRegion(db *gorm.DB, id uint) (*PSIRegionSchema, error) {

//     var region PSIRegionSchema

//     err := db.First(&region, id).Error
//     return &region, err

// }

// GetAllUsers retrieves all users
func GetNearestRecord(db *gorm.DB) (*PSIRollupTableArchiveSchema, error) {

    var nearestRecord PSIRollupTableArchiveSchema

    err := db.Find(&nearestRecord).Error
    return &nearestRecord, err

}
