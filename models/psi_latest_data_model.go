
package models

import (

    "time"
    "errors"
    "gorm.io/gorm"

)

type PSIHourlyDataLatest struct {
    RecordID             uint   `gorm:"primaryKey;autoIncrement" json:"record_id"`
    Region               string `gorm:"not null" json:"region"`
    CreatedDate          string `gorm:"not null" json:"__date"`
    Timestamp            int    `gorm:"not null;index" json:"timestamp"`
    UpdatedTimestamp     int    `gorm:"not null;default:0" json:"updated_timestamp"`
    CoSubIndex           int    `gorm:"not null;default:0" json:"co_sub_index"`
    So2TwentyFourHourly  int    `gorm:"not null;default:0" json:"so2_twenty_four_hourly"`
    So2SubIndex          int    `gorm:"not null;default:0" json:"so2_sub_index"`
    PsiThreeHourly       int    `gorm:"not null;default:0" json:"psi_three_hourly"`
    CoEightHourMax       int    `gorm:"not null;default:0" json:"co_eight_hour_max"`
    No2OneHourMax        int    `gorm:"not null;default:0" json:"no2_one_hour_max"`
    Pm10SubIndex         int    `gorm:"not null;default:0" json:"pm10_sub_index"`
    Pm25SubIndex         int    `gorm:"not null;default:0" json:"pm25_sub_index"`
    O3EightHourMax       int    `gorm:"not null;default:0" json:"o3_eight_hour_max"`
    PsiTwentyFourHourly  int    `gorm:"not null;default:0" json:"psi_twenty_four_hourly"`
    O3SubIndex           int    `gorm:"not null;default:0" json:"o3_sub_index"`
    Pm25TwentyFourHourly int    `gorm:"not null;default:0" json:"pm25_twenty_four_hourly"`
    Pm10TwentyFourHourly int    `gorm:"not null;default:0" json:"pm10_twenty_four_hourly"`
}

func CreatePSIDataRecord(db *gorm.DB, psiData *PSIHourlyDataLatest) error {

    return db.Create(psiData).Error

}

func GetHourlyPSIData(db *gorm.DB, dateTimeValue string) (*PSIHourlyDataLatest, error) {

    var dataRecord PSIHourlyDataLatest

    dateTime, err := time.Parse("2006-01-02 15:04:05", dateTimeValue)

    if err != nil {
        return nil, errors.New("invalid datetime format, expected it to be in the format: YYYY-MM-DD HH:MM:SS")
    }

    timestamp := dateTime.Truncate(time.Hour).Unix()

    return &dataRecord, db.Where("timestamp = ?", timestamp).Find(&dataRecord).Error

}

func GetDailyPSIData(db *gorm.DB, dateValue string) ([]PSIHourlyDataLatest, error) {

    var dataRecords []PSIHourlyDataLatest

    if _, err := time.Parse("2006-01-02", dateValue); err != nil {
        return nil, errors.New("invalid date format, expected it to be in the format: YYYY-MM-DD")
    }

    return dataRecords, db.Where("created_date = ?", dateValue).Find(&dataRecords).Error

}


