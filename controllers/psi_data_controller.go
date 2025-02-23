package controllers

import (
    "log"
    "regexp"
    "net/http"
	"github.com/gin-gonic/gin"

	"gnz_psi/go_web_service/models"
	"gnz_psi/go_web_service/middlewares"
)

// CreateUserHandler creates a new user
func CreateUserHandler(c *gin.Context) {

    c.JSON(http.StatusCreated, gin.H{"message": "User created"})

}

type PSIDataRequest struct {

	Date string `json:"date" binding:"required"`

}

func GetPSIData(c *gin.Context) {
    
    type PSIRequest struct {

        Date string `json:"date" binding:"required"`

    }
    
    var req PSIRequest

    // Bind and validate the request body
    if err := c.ShouldBindQuery(&req); err != nil {

        c.JSON(http.StatusBadRequest, gin.H{

            "statusCode": http.StatusBadRequest,
            "code": 1,
            "message": "Invalid request parameters:" + err.Error(),
            "data": nil,

        })
        
        return

    }

    db_conn, db_conn_err := middlewares.GetDBInterface()

    if db_conn_err != nil {

        log.Printf("[ERROR] PSI DATA CNTR: GetPSIData: While connecting to database: %v", db_conn_err)

        c.JSON(http.StatusInternalServerError, gin.H{

            "statusCode": http.StatusInternalServerError,
            "code":       2,
            "message":    "Could not establish database connection",
            "data":       nil,

        })

        return

    }

    var psiRecords []models.PSIHourlyDataLatest
    var db_fetch_err error

    // Determine if input is a date or datetime
    datePattern := `^\d{4}-\d{2}-\d{2}$`
    datetimePattern := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`

    if matched, _ := regexp.MatchString(datePattern, req.Date); matched {

        // Input is a date, fetch daily PSI data
        psiRecords, db_fetch_err = models.GetDailyPSIData(db_conn, req.Date)

    } else if matched, _ := regexp.MatchString(datetimePattern, req.Date); matched {

        // Input is a datetime, fetch hourly PSI data
        psiRecords, db_fetch_err = models.GetHourlyPSIData(db_conn, req.Date)

        // if db_fetch_err == nil {

        //     psiRecords = []models.PSIHourlyDataLatest{psiRecord}

        // }

    } else {

        c.JSON(http.StatusBadRequest, gin.H{

            "statusCode": http.StatusBadRequest,
            "code":      1,
            "message":   "Invalid date format. Use YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS.",
            "data":      nil,

        })

        return

    }

    if db_fetch_err != nil {

        c.JSON(http.StatusInternalServerError, gin.H{

            "statusCode": http.StatusInternalServerError,
            "code":      2,
            "message":   "Failed to retrieve data: " + db_fetch_err.Error(),
            "data":      nil,

        })
    
        return

    }

    // Fetch region metadata
    regions, db_fetch_err := models.GetAllRegions(db_conn)
    if db_fetch_err != nil {

        c.JSON(http.StatusInternalServerError, gin.H{

            "statusCode": http.StatusInternalServerError,
            "code":      3,
            "message":   "Failed to retrieve region metadata: " + db_fetch_err.Error(),
            "data":      nil,

        })
        
        return

    }

    regionMetadata := []gin.H{}
    for _, region := range regions {

        regionMetadata = append(regionMetadata, gin.H{

            "regionID":       region.RegionID,
            "name":           region.Name,
            "labelLocation":  gin.H{
                "latitude":   region.Latitude,
                "longitude":  region.Longitude,
            },

        })

    }

    // Transform the data to match the expected response structure
    readings := make(map[string]map[string]int)
    for _, record := range psiRecords {

        if readings["o3_sub_index"] == nil {

            readings["o3_sub_index"] = make(map[string]int)
            readings["no2_one_hour_max"] = make(map[string]int)
            readings["o3_eight_hour_max"] = make(map[string]int)
            readings["psi_twenty_four_hourly"] = make(map[string]int)
            readings["pm10_twenty_four_hourly"] = make(map[string]int)
            readings["pm10_sub_index"] = make(map[string]int)
            readings["pm25_twenty_four_hourly"] = make(map[string]int)
            readings["so2_sub_index"] = make(map[string]int)
            readings["pm25_sub_index"] = make(map[string]int)
            readings["so2_twenty_four_hourly"] = make(map[string]int)
            readings["co_eight_hour_max"] = make(map[string]int)
            readings["co_sub_index"] = make(map[string]int)

        }

        region := record.Region
        readings["o3_sub_index"][region] = record.O3SubIndex
        readings["no2_one_hour_max"][region] = record.No2OneHourMax
        readings["o3_eight_hour_max"][region] = record.O3EightHourMax
        readings["psi_twenty_four_hourly"][region] = record.PsiTwentyFourHourly
        readings["pm10_twenty_four_hourly"][region] = record.Pm10TwentyFourHourly
        readings["pm10_sub_index"][region] = record.Pm10SubIndex
        readings["pm25_twenty_four_hourly"][region] = record.Pm25TwentyFourHourly
        readings["so2_sub_index"][region] = record.So2SubIndex
        readings["pm25_sub_index"][region] = record.Pm25SubIndex
        readings["so2_twenty_four_hourly"][region] = record.So2TwentyFourHourly
        readings["co_eight_hour_max"][region] = record.CoEightHourMax
        readings["co_sub_index"][region] = record.CoSubIndex

    }

    // Construct response
    response := gin.H{

        "statusCode": http.StatusOK,
        "code":      0,
        "message":   "Success",
        "data": gin.H{
            "regionMetadata": regionMetadata,
            "items":         []gin.H{},
        },

    }

    if len(psiRecords) > 0 {

        response["data"].(gin.H)["items"] = []gin.H{
            {

                "date":             req.Date,
                "updatedTimestamp": psiRecords[0].UpdatedTimestamp,
                "timestamp":        psiRecords[0].Timestamp,
                "readings":         readings,
            },
        }

    }

    // Send response
    c.JSON(http.StatusOK, response)
    
}
