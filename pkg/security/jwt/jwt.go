package jwt

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/Mona-bele/logutils-go/logutils"
	"github.com/Mona-bele/rote-notify/pkg/env"
	"github.com/golang-jwt/jwt/v5"
)

const (
	ExpireDays = 365
	ExpireTime = time.Hour * 24 * ExpireDays
	Algorithm  = "RS256"
)

// JWT struct
type JWT struct {
	PrivateKey *rsa.PrivateKey
	*env.Env
}

// NewJWT creates a new JWT instance
func NewJWT(privateKey *rsa.PrivateKey, env *env.Env) *JWT {
	return &JWT{PrivateKey: privateKey, Env: env}
}

// GenerateToken generates a JWT token
func (j *JWT) GenerateToken(payload, issuer, audience, subject string) (string, error) {
	claims := jwt.MapClaims{
		"exp":     time.Now().Add(ExpireTime).Unix(),
		"iat":     time.Now().Unix(),
		"sub":     subject,
		"iss":     issuer,
		"aud":     audience,
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = j.JwtKid

	signedToken, err := token.SignedString(j.PrivateKey)
	if err != nil {
		logutils.Error("Failed to sign the token", err, nil)
		return "", err
	}

	return signedToken, nil
}

// ParseToken parses a JWT token
func (j *JWT) ParseToken(tokenString, issuer, audience, subject string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		j.ValidateToken,
		jwt.WithIssuer(issuer),
		jwt.WithSubject(subject),
		jwt.WithAudience(audience),
		jwt.WithValidMethods([]string{Algorithm}),
	)
	if err != nil {
		logutils.Error("Failed to parse the token", err, nil)
		return nil, err
	}

	if !token.Valid {
		logutils.Error("Token is invalid", nil, nil)
		return nil, errors.New("token is invalid")
	}

	return token, nil
}

func (j *JWT) ValidateToken(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		logutils.Error("Unexpected signing method", nil, nil)
		return nil, errors.New("unexpected signing method")
	}

	if _, ok := token.Header["kid"]; !ok {
		logutils.Error("Token kid is missing", nil, nil)
		return nil, errors.New("token kid is missing")
	}

	if token.Header["kid"] != j.JwtKid {
		logutils.Error("Token kid is invalid", nil, nil)
		return nil, errors.New("token kid is invalid")
	}

	return &rsa.PublicKey{
		N: j.PrivateKey.PublicKey.N,
		E: j.PrivateKey.PublicKey.E,
	}, nil
}

// GetPayload returns the payload from a JWT token
func (j *JWT) GetPayload(token *jwt.Token) string {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logutils.Error("Failed to get claims", nil, nil)
		return ""
	}

	payload, ok := claims["payload"].(string)
	if !ok {
		logutils.Error("Failed to get payload", nil, nil)
		return ""
	}

	return payload
}

// ParsePrivateKey parses a private key
func ParsePrivateKey(pem []byte) (*rsa.PrivateKey, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		logutils.Error("Failed to parse the private key", err, nil)
		return nil, err
	}

	return privateKey, nil
}

// ParsePublicKey parses a public key
func ParsePublicKey(pem []byte) (*rsa.PublicKey, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	if err != nil {
		logutils.Error("Failed to parse the public key", err, nil)
		return nil, err
	}

	return publicKey, nil
}

// NewJWTFromEnv creates a new JWT instance from the environment
func NewJWTFromEnv(env *env.Env) (*JWT, error) {
	privateKey, err := ParsePrivateKey([]byte(env.JwtNotifyPrivateKey))
	if err != nil {
		logutils.Error("Failed to parse the private key", err, nil)
		return nil, err
	}

	return NewJWT(privateKey, env), nil
}
