package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"aurora/internal/auth"
	"aurora/internal/utils"

	"context"
	"errors"
	"net/http"
	"time"
)

type UserServicer interface {
	ListUsers(ctx context.Context) ([]gen.User, error)
	Register(params RegisterParams, ctx context.Context) (*http.Cookie, error)
	Login(params LoginParams, ctx context.Context) (*http.Cookie, error)
	GetAuthInfo(cookie string, ctx context.Context) (AuthInfo, error)
}

var UserService UserServicer = UserServiceStruct{}

type UserServiceStruct struct {
}

func (u UserServiceStruct) ListUsers(ctx context.Context) ([]gen.User, error) {
	return db.Queries.ListUsers(ctx)
}

type RegisterParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u UserServiceStruct) Register(params RegisterParams, ctx context.Context) (*http.Cookie, error) {

	if params.FirstName == "" || params.LastName == "" || params.Email == "" || params.Password == "" {
		return nil, errors.New("no register parameter can be empty")
	}

	if !utils.ValidateEmail(params.Email) {
		return nil, errors.New("invalid email")
	}

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		return nil, errors.New("unhashable password")
	}

	user, err := db.Queries.CreateUser(ctx, gen.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Hash:      hash,
	})

	if err != nil {
		return nil, errors.New("failed to create a new user in the database")
	}

	cookieValue, err := auth.NewSessionCookie()
	if err != nil {
		return nil, errors.New("failed to generate a cookie")
	}

	cookie := http.Cookie{
		Name:     "session_cookie",
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * 7 * 24),
	}

	_, err = db.Queries.CreateSession(ctx, gen.CreateSessionParams{
		UserID:    user.ID,
		Cookie:    cookieValue,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: cookie.Expires.Unix(),
	})

	if err != nil {
		return nil, errors.New("failed to create a new session in the database")
	}

	return &cookie, nil
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserServiceStruct) Login(params LoginParams, ctx context.Context) (*http.Cookie, error) {

	user, err := db.Queries.GetUserByEmail(ctx, params.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &UnknownEmail{params.Email}
		}
		return nil, err
	}

	if len(params.Password) < 6 {
		return nil, &BadPasswordError{"length must be at least 6"}
	}

	if !auth.CheckPasswordHash(params.Password, user.Hash) {
		return nil, &BadPasswordError{"unhashable"}
	}

	cookieValue, err := auth.NewSessionCookie()
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     "session_cookie",
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * 7 * 24),
	}

	_, err = db.Queries.CreateSession(ctx, gen.CreateSessionParams{
		UserID:    user.ID,
		Cookie:    cookieValue,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: cookie.Expires.Unix(),
	})

	if err != nil {
		return nil, err
	}

	return &cookie, nil
}

type AuthInfo struct {
	Session  gen.Session
	User     gen.User
	UserType string
}

func (u UserServiceStruct) GetAuthInfo(cookie string, ctx context.Context) (AuthInfo, error) {

	authInfo, err := db.Queries.GetUserBySessionCookie(ctx, cookie)
	if err != nil {
		return AuthInfo{}, err
	}

	userType, err := db.Queries.GetUserType(ctx, authInfo.UserID)

	return AuthInfo{
		Session: gen.Session{
			ID:        authInfo.ID,
			UserID:    authInfo.UserID,
			Cookie:    cookie,
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
