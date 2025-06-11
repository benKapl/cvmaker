package auth_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/benKapl/cvmaker-api/internal/auth"
	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := auth.HashPassword(password1)
	hash2, _ := auth.HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := auth.CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func ExampleCheckPasswordHash() {
	// Create a password
	password := "securePassword123!"

	// Hash the password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	// Check if correct password matches the hash
	err = auth.CheckPasswordHash(password, hashedPassword)
	fmt.Println("Correct password check result:", err == nil)

	// Check if incorrect password matches the hash
	wrongPassword := "wrongPassword456!"
	err = auth.CheckPasswordHash(wrongPassword, hashedPassword)
	fmt.Println("Incorrect password check result:", err == nil)

	// Output:
	// Correct password check result: true
	// Incorrect password check result: false
}

func TestValidateJwt(t *testing.T) {
	// Create data necessary for test
	userID := uuid.New()
	validSecret := "this_is_the_secret_to_use"
	wrongSecret := "this_is_wrong_secret"
	validExpiration := 24 * time.Hour
	passedExpiration := -24 * time.Hour // yesterday

	// Create tests jwt
	jwt, _ := auth.MakeJWT(userID, validSecret, validExpiration)
	expiredJWT, _ := auth.MakeJWT(userID, validSecret, passedExpiration)

	tests := []struct {
		name    string
		jwt     string
		secret  string
		wantErr bool
	}{
		{
			name:    "Valid JWT",
			jwt:     jwt,
			secret:  validSecret,
			wantErr: false,
		},
		{
			name:    "Wrong secret JWT",
			jwt:     jwt,
			secret:  wrongSecret,
			wantErr: true,
		},
		{
			name:    "Expired JWT",
			jwt:     expiredJWT,
			secret:  validSecret,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := auth.ValidateJWT(tt.jwt, tt.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func ExampleValidateJWT() {
	// Create a test user ID
	userID := uuid.New()

	// Define JWT parameters
	tokenSecret := "my-secure-secret"
	expirationTime := 24 * time.Hour

	// Create a JWT token
	token, err := auth.MakeJWT(userID, tokenSecret, expirationTime)
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return
	}

	// Validate the JWT token
	retrievedUserID, err := auth.ValidateJWT(token, tokenSecret)
	if err != nil {
		fmt.Println("JWT validation failed:", err)
		return
	}

	fmt.Println("JWT validation successful, User ID matches:", retrievedUserID == userID)
	// Output: JWT validation successful, User ID matches: true
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "Valid Bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token"},
			},
			wantToken: "valid_token",
			wantErr:   false,
		},
		{
			name:      "Missing Authorization header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InvalidBearer token"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := auth.GetBearerToken(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("GetBearerToken() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

func ExampleGetBearerToken() {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-token")

	token, err := auth.GetBearerToken(headers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Bearer", token)
	// Output: Bearer my-secret-token
}

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr bool
	}{
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			wantKey: "valid_api_key",
			wantErr: false,
		},
		{
			name:    "Missing Authorization header",
			headers: http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InvalidApiKey key"},
			},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "Empty API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey "},
			},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := auth.GetApiKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetApiKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}

func ExampleGetApiKey() {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-api-key-12345")

	apiKey, err := auth.GetApiKey(headers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("API Key:", apiKey)
	// Output: API Key: my-api-key-12345
}
