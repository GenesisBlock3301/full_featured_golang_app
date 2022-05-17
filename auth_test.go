package main

import (
	"bookshop/config"
	controller "bookshop/controllers"
	"bookshop/services"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
)

func getLoginPayload() string {
	params := url.Values{}
	params.Add("name", "nase111")
	params.Add("email", "nas111@gmail.com")
	params.Add("password", "nas12345")
	return params.Encode()
}

func TestUserExistence(t *testing.T) {
	config.ConnectWithDB()
	// User exist or not
	_, err := services.GetUserById(1)
	fmt.Println("first case", err)
	if err != nil {
		t.Fail()
	}
	// Email and password valid or not
	isValidCredential, _ := services.VerifiyCredentialService("nas@gmail.com", "nas12345")
	if !isValidCredential {
		t.Fail()
	}
	// Generate Token from user ID
	token, refresh, err := services.GenerateTokenPair(1)
	fmt.Println("test case 3: ", err, token)
	if err != nil {
		t.Error("generated token", token, refresh)
	}

	// user_id, err := services.ExtractTokenID(ctx)

}



func TestUserRegistration(t *testing.T) {
	config.ConnectWithDB()
	r := gin.New()

	r.POST("/api/v1/auth/register", controller.Register)
	Payload := map[string]string{
		"name":     "nas",
		"email":    "nur11@gmail.com",
		"password": "nas12345",
	}
	jsonString, _ := json.Marshal(Payload)
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(string(jsonString)))
	req.Header.Add("content-type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("user test authentication", w.Code,http.StatusOK)
	if w.Code != http.StatusOK {
		t.Fail()
	}
}

func TestUserLogin(t *testing.T) {
	config.ConnectWithDB()

	r := gin.New()
	r.POST("/api/v1/auth/login", controller.Login)
	payload := map[string]string{
		"name":     "nas",
		"email":    "nur11@gmail.com",
		"password": "nas12345",
	}
	jsonString, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(string(jsonString)))
	req.Header.Add("content-type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("user test authentication", w.Code,http.StatusOK)
	if w.Code != http.StatusOK {
		t.Fail()
	}

}

func TestGetUser(t *testing.T) {
	config.ConnectWithDB()
	r := getRouter(true)
	r.GET("/api/v1/auth/user")
	req, _ := http.NewRequest("GET", "/api/v1/auth/user", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTI3NzM2MjksInVzZXJfaWQiOjF9.n2PHd--bi4jDqzyyHhoKEyRUxZEzGdAMjMDtHCLNh8w")
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}

// func TestAllPosts
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}
