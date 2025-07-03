package auth

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

var InvalidCookieErr = errors.New("invalid cookie")

var AuthService Auther = AuthStruct{}

type Auther interface {
	Authenticate(r *http.Request) (*AuthInfo, error)
	GetAuthInfo(cookieValue string, ctx context.Context) (*AuthInfo, error)
}

type AuthInfo struct {
	Session  gen.Session
	User     gen.User
	UserType string
}

type AuthStruct struct {
}

func (a AuthStruct) GetAuthInfo(cookieValue string, ctx context.Context) (*AuthInfo, error) {

	authInfo, err := db.Queries.GetUserBySessionCookie(ctx, cookieValue)
	if err != nil {
		return nil, err
	}

	userType, err := db.Queries.GetUserType(ctx, authInfo.UserID)

	return &AuthInfo{
		Session: gen.Session{
			ID:        authInfo.ID,
			UserID:    authInfo.UserID,
			Cookie:    cookieValue,
			CreatedAt: authInfo.CreatedAt,
			ExpiresAt: authInfo.ExpiresAt,
		},
		User: gen.User{
			ID:        authInfo.UserID,
			FirstName: authInfo.FirstName,
			LastName:  authInfo.LastName,
			Hash:      authInfo.Hash,
			Email:     authInfo.Email,
		},
		UserType: userType,
	}, nil
}

func (a AuthStruct) Authenticate(r *http.Request) (*AuthInfo, error) {

	cookie, err := r.Cookie("session_cookie")
	if err != nil {
		return nil, err
	}

	authInfo, err := a.GetAuthInfo(cookie.Value, r.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, InvalidCookieErr
		}
	}

	return authInfo, nil
}
