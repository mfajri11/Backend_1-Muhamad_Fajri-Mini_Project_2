package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/modules/auth"
	"time"
)

type ITokenManager interface {
	GenerateToken(username string, role string, exp time.Time) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type TokenManager struct {
	secretKey string
	issuer    string
}

func NewTokenManager(secretKey string) *TokenManager {
	return &TokenManager{secretKey: secretKey, issuer: "mini.project.com"}
}

func (t *TokenManager) GenerateToken(name, role string, exp time.Time) (string, error) {
	claims := auth.JWTClaims{
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

func (t *TokenManager) ValidateToken(tokenStr string) (*jwt.Token, error) {

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("modules.TokenManager.ValidateToken: miss match algo type got %s want 'HS256:'", token.Method.Alg())
		}

		return []byte(t.secretKey), nil
	})
}
