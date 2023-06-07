package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenManager struct {
	secretKey string
	issuer    string
}

func NewTokenManager(secretKey string) *TokenManager {
	return &TokenManager{secretKey: secretKey, issuer: "mini.project.com"}
}

func (t *TokenManager) GenerateToken(name, role string, exp time.Time) (string, error) {
	claims := &JWTClaims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  exp.Unix(),
			Issuer:    t.issuer,
			Id:        name,
			ExpiresAt: exp.Add(48 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", fmt.Errorf("modules.TokenManager.GenerateToken: error signing token: %w", err)
	}

	return tokenStr, nil
}

func (t *TokenManager) ValidateToken(tokenStr string) (any, error) {
	claim := &JWTClaims{}
	payload, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("modules.TokenManager.ValidateToken: miss match algo type got %s want 'HS256:'", token.Method.Alg())
		}

		return []byte(t.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("modules.TokenManager.ValidateToken: error validate token: %w", err)
	}
	claim, ok := payload.Claims.(*JWTClaims)
	if !ok {
		return nil, fmt.Errorf("modules.TokenManager.ValidateToken: error invalid claim")
	}

	return claim, nil
}

type JWTClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
