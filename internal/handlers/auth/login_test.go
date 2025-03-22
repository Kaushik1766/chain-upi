package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	err := godotenv.Load(`D:\chain-upi\.env`)
	if err != nil {
		log.Fatal("Cant find .env")
	}
	err = db.InitDB()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB connected.")
	}

	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid login",
			requestBody: map[string]interface{}{
				"password": "1234",
				"email":    "kaushiksaha004@gmail.com",
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"token": "some-token",
			},
		},
		{
			name: "Invalid credentials",
			requestBody: map[string]interface{}{
				"password": "wrongpassword",
				"email":    "valid@example",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]interface{}{
				"error": "Invalid credentials",
			},
		},
		{
			name: "Missing fields",
			requestBody: map[string]interface{}{
				"email": "test@example",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid data",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.Default()
			authGroup := r.Group("/api/auth")
			authGroup.POST("/login", Login())

			jsonBody, _ := json.Marshal(tc.requestBody)

			req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				for _, cookie := range w.Result().Cookies() {
					if cookie.Name == "token" {
						assert.NotEmpty(t, cookie.Value, "empty token")
					}
				}
			}
		})
	}
}
