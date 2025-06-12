package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"aurora/internal/auth"
	"aurora/internal/utils"
	"errors"
	"time"

	"context"
	"net/http"
)

type UserService struct {
}

func (u UserService) ListUsers(ctx context.Context) ([]gen.User, error) {
	return db.Queries.ListUsers(ctx)
}

type RegisterParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u UserService) Register(params RegisterParams, ctx context.Context) (*http.Cookie, error) {

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

func (u UserService) Login(ctx context.Context, params LoginParams) (*http.Cookie, error) {

	user, err := db.Queries.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	if !auth.CheckPasswordHash(params.Password, user.Hash) {
		return nil, errors.New("invalid password")
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
