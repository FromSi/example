package tools

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	ExpirationInDays             = 30
	CommonJWTClaimIssuer         = "iss"
	CommonJWTClaimAudience       = "aud"
	CommonJWTClaimSubject        = "sub"
	CommonJWTClaimIssuedAt       = "iat"
	CommonJWTClaimExpirationTime = "exp"
)

type SessionJWT struct {
	Issuer         string
	Audience       string
	Subject        string
	IssuedAt       time.Time
	ExpirationTime time.Time
}

func NewSessionJWT(issuer string, audience string, subject string, timeNow time.Time) *SessionJWT {
	return &SessionJWT{
		Issuer:         issuer,
		Audience:       audience,
		Subject:        subject,
		IssuedAt:       timeNow,
		ExpirationTime: timeNow.AddDate(0, 0, ExpirationInDays),
	}
}

func NewSessionJWTFromString(tokenJWT string, secretKey string) (*SessionJWT, error) {
	token, err := jwt.Parse(tokenJWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid claims")
	}

	sessionJWT := SessionJWT{
		Issuer:         claims[CommonJWTClaimIssuer].(string),
		Audience:       claims[CommonJWTClaimAudience].(string),
		Subject:        claims[CommonJWTClaimSubject].(string),
		IssuedAt:       time.Unix(int64(claims[CommonJWTClaimIssuedAt].(float64)), 0),
		ExpirationTime: time.Unix(int64(claims[CommonJWTClaimExpirationTime].(float64)), 0),
	}

	if time.Now().Before(sessionJWT.IssuedAt) {
		return nil, errors.New("token used before issued")
	}

	if time.Now().After(sessionJWT.ExpirationTime) {
		return nil, errors.New("token has expired")
	}

	return &sessionJWT, nil
}

func (sessionJWT SessionJWT) GetJWT(secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CommonJWTClaimIssuer:         sessionJWT.Issuer,
		CommonJWTClaimAudience:       sessionJWT.Audience,
		CommonJWTClaimSubject:        sessionJWT.Subject,
		CommonJWTClaimIssuedAt:       sessionJWT.IssuedAt.Unix(),
		CommonJWTClaimExpirationTime: sessionJWT.ExpirationTime.Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
