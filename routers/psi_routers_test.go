package routers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "gnz_psi/go_web_service/configs"
)

// TestRouter ensures that the router correctly routes requests.
func TestRouter(t *testing.T) {

    router := SetupRouter()

    req, _ := http.NewRequest("GET", fetch_daily_psi_data, nil)
    rr := httptest.NewRecorder()

    router.ServeHTTP(rr, req) // Testing the router

    if rr.Code != http.StatusOK {

        t.Errorf("Expected status OK, got %d", rr.Code)

    }

}
