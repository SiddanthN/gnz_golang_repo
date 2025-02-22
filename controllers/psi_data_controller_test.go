
package controllers

import (
    "net/http"
    "net/http/httptest"
    "testing"
	"encoding/json"

    "gnz_psi/go_web_service/configs"
)

// InnerResult defines the structure inside the 'result' key.
type InnerResult struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Details string `json:"details"`
}

// MockAPIResponse defines the outer structure of the API response.
type MockAPISuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	ErrorMsg   string     `json:"errorMsg"`
	Result     InnerResult `json:"result"`
}



func TestFetchCntrForSuccessOutput(t *testing.T) {

    tests := []struct {
        testTitle            string
        inputDateTime        string
        hasCorrectStructure  bool
    }{
        {"Valid datetime with no hour value", "2025-02-24", true},
        {"Valid datetime with imperfect hour value", "2025-02-24 13:15:00", true}
    }

    for _, tt := range tests {

        t.Run(tt.testTitle, func(t *testing.T) {

            var fetchedRecord 
            err := db.Find(&fetchedRecord, tt.inputDateTime).Error

            if (err == nil) != tt.shouldExist {
                t.Errorf("Test %s failed: expected existence=%v, got error %v", tt.testTitle, tt.shouldExist, err)
            }

        })

    }




}




