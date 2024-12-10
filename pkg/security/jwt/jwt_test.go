package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Mona-bele/rote-notify/pkg/env"
	"github.com/stretchr/testify/assert"
)

const (
	testPemBase64 = "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRQ3lNcnRhWncydGtMSHQKQXJqQ1JPdmVVeTFTYnc1TjU1WFVaS1p2R2x6SGg2UnFTOUlJMXpmRnoySEFFNlpURmpJL0dCZ2QwalEvWmI3QwpnZE5DbzZmWGNNSUhaeEtmY0g1K0dMOVVVZ0tnVjR2clppWHR4ekEzb3pDWGFRcGRWVWhRVkNwdmdOejUwdVkzCjVldjRBYzBEa3FLZTBhcDBmcGVYTmdUMUdwdG1oRmhJNC81aGhUTVZwaHVBQjUrNXpiSFVZQ1NoYmhsUytZamMKOFZDNDZDTXFuYk5vd0d3MkpHRnBYN3E0L1lUU2hSNlVHV2NhUXJiR2FFdzMrNDR6ZlFreVNGU1daZDYwVXlFYQpBcHNHcnh4dmNJbXVyL3JWTWdaQU0rZlFIQ0d5Zm9odnNRN3ZEVVQ4bFJVODZ1NlJ4Zm4yalVIN2pFQjhWUHoxCkZpd2lFeWNWQWdNQkFBRUNnZ0VBQXhsYVJFVjRDYTF6UnJPdUttUy85RjhlY2VsUUJwOVJMZjh0SE5BSXpvaUgKbmJpaFY3bUVkUXA1QVBubWdkbm9PRXd2T1MrTGJKSFVxalhQU29DbG5kdk9hQW9OR3h3OSt1dnJ3ZWZtTjBzbQplZ3JOL0k5akFZaXpRdUxYQlE0RFlyQkFCNjNtU2VyMnlZQ1VaTDBGUjN0Z25DSDUxS3BSOFJXRnF3eVNLTXQ3CmUwcmlybzgyVjJEUVZJSUtzNFZMaHJTMk1WZ3N0R21PSGw2S0ZVd2d4Nit5SkEvbVA1SmZIdlpPakRFcnRLS3QKVXl3TWZldUhlS0dZbnBReXcrM29YM3pGQ1lRbmNLQVFVSjF0cmJqZjNpcEFqd2ErTUVuejJtc0VGdHM3c0ZmZwp2VUlBRStBN0xYcHZVcTYxOFpLNStTL1JNbXE3WE40VmF0Skx6Rk5KcVFLQmdRRGkyZkp4Vm9rWUZ3Ui9UWElpCkxNUnduTnd4Mkl3SzdOZVErQ1MrME9sRjdRYnZqTFl1UXcwekhrQklBLzhib2w5Tk1qcXJObUZOYUtwaG4rZm8KQ05COUFYeDlaMEZZdW1RYWhLbDZhQlFTc0dJQVNxYTdHVUpZdVpUMjVmbG9RaUx4eVBBMllUODhuWE5wQld3TgpZY0NhVXZIQXhTRHhEWFg0WkRTUkJSOTZpUUtCZ1FESkdHTFJEbUN5ZVdiL0Fwb3BCdkhGdktmM0NpbUhEUzlGCkFpR200MG1XQmtpeURxNUFGd0JXN2Q2VzdQd3dFWFZsSjJNcFNGVTVDZkpyUWJMZ2V3dTI4a1Fya3RPc1lNMzMKR2UwVWRPMnNzVVc5bXpqOGtZSmZjbEt2SWVFRC9oSjFqQWJudmpINDhGRXFsdzlWQmQ1SDUrK2x4WU1jd3RUWQpyUVdoQ2JwMUxRS0JnR0s5bHdlNk1PWXBicTJ5bWhGQ0J5YzFQNnI2cE1wRW1QZmk2cXViNTAybWhEUlV2UitaCjArOENKZHl5MEtISXBVN0dwRDdONXNCNDVHQ2w1NTFaNk5YZ2hiMVg2bHVpbGR2dERvL1hLWldRN0xkUHh3NzkKU2FHdzlhUWFLZHMxbmx3N3FFTWpSUkV2UDRMZzkwMUQxVC9YQnA5dnJvejkzUEdIUEZJN05wNXhBb0dBS1dTZApvbzZsRk5lc2ZiMVpZaXlOdzdnTGt1eENsQXdBdU9HeGI1ckZZTjQyUklDRkdhZ2laOEphMlJJNjd1SUpHaU03ClpCb0JnTll0VWlxWjJWODZrQlBhT0dYbXNFclUycEk1bk1aY3pmbEhjN25weHdOa3BLVHhwQjhESkVFK0ozZ0YKUzlwNGl0ZGN2Ym1PYkYvaTIwWkFyQXkxNmt1b2FGbGxHVHJaYUprQ2dZRUF6M2JxNTA5dW45Sk00eGxVa3gxSQpnbHI3NXhQYWp1VjloazdsMnZGT1Q1RUJQWUZGeTF4bk5DVG9mRVBkcHJRMkl3RW1CUllDd1FlM0UySzZwTlFwCmVGQUxVL3c1UWJpRFVQZXVXWHdWSjZrSlI5WGtMOWVaeVl4bXE3alRPWGpxWGJRZDVKMEMrTTdQeGJ2RDRlYzkKMkxBTjVTY2dLTzBuR1QwM2l2MC9pZzQ9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K"
)

func TestNewJWT(t *testing.T) {
	t.Run("Test NewJWT", func(t *testing.T) {
		// TestNewJWT tests the NewJWT function
		// It should return a new JWT instance
		// Arrange
		privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		envJWT := &env.Env{
			JwtKid: "JWT_KID_1234",
		}
		// Act
		jwt := NewJWT(privateKey, envJWT)
		// Assert
		if jwt == nil {
			t.Errorf("Expected a new JWT instance, got nil")
		}
	})
}

func TestGenerateToken(t *testing.T) {
	t.Run("Test GenerateToken", func(t *testing.T) {
		// TestGenerateToken tests the GenerateToken function
		// It should return a JWT token
		// Arrange
		file, err := os.Open("private_key.pem")

		body, _ := io.ReadAll(file)

		privateKey, _ := ParsePrivateKey(body)

		envJWT := &env.Env{
			JwtKid: "JWT_KID_1234",
		}
		jwt := NewJWT(privateKey, envJWT)
		// Act
		token, err := jwt.GenerateToken("payloadTESTE", "issuer", "audience", "subject")
		fmt.Println(token)
		// Assert
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if token == "" {
			t.Errorf("Expected a JWT token, got an empty string")
		}
	})
}

func TestParseToken(t *testing.T) {
	t.Run("Test ParseToken", func(t *testing.T) {
		// TestParseToken tests the ParseToken function
		// It should return a JWT token
		// Arrange
		bpk, err := base64.StdEncoding.DecodeString(testPemBase64)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		envJWT := &env.Env{
			JwtNotifyPrivateKey: string(bpk),
			JwtKid:              "JWT_KID_1234",
		}

		jwt, err := NewJWTFromEnv(envJWT)

		token, err := jwt.GenerateToken("payloadTESTE", "issuer", "audience", "subject")
		// Act
		content, err := jwt.ParseToken(token, "issuer", "audience", "subject")
		// Assert
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if content == nil {
			t.Errorf("Expected a JWT token, got nil")
		}

		assert.NotNil(t, interface{}(content))
	})

	t.Run("Test GetPayload", func(t *testing.T) {
		// TestGetPayload tests the GetPayload function
		// It should return a JWT token
		// Arrange
		bpk, err := base64.StdEncoding.DecodeString(testPemBase64)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		privateKey, err := ParsePrivateKey(bpk)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		envJWT := &env.Env{
			JwtKid: "JWT_KID_1234",
		}
		jwt := NewJWT(privateKey, envJWT)
		token, _ := jwt.GenerateToken("payloadTESTE", "issuer", "audience", "subject")
		// Act
		content, _ := jwt.ParseToken(token, "issuer", "audience", "subject")
		payload := jwt.GetPayload(content)
		// Assert
		assert.NotNil(t, interface{}(payload))
		assert.Equal(t, "payloadTESTE", payload)
	})

	t.Run("Test ParseToken with invalid token", func(t *testing.T) {
		// TestParseToken tests the ParseToken function
		// It should return a JWT token
		// Arrange
		bpk, err := base64.StdEncoding.DecodeString(testPemBase64)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		privateKey, err := ParsePrivateKey(bpk)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		envJWT := &env.Env{
			JwtKid: "JWT_KID_1234",
		}
		jwt := NewJWT(privateKey, envJWT)
		// Act
		_, err = jwt.ParseToken("invalidToken", "issuer", "audience", "subject")
		// Assert
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("Test ParseToken with invalid issuer", func(t *testing.T) {
		// TestParseToken tests the ParseToken function
		// It should return a JWT token
		// Arrange
		bpk, err := base64.StdEncoding.DecodeString(testPemBase64)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		privateKey, err := ParsePrivateKey(bpk)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		envJWT := &env.Env{
			JwtKid: "JWT_KID_1234",
		}
		jwt := NewJWT(privateKey, envJWT)
		token, _ := jwt.GenerateToken("payloadTESTE", "issuer", "audience", "subject")
		// Act
		_, err = jwt.ParseToken(token, "invalidIssuer", "audience", "subject")
		// Assert
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
