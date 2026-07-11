package users

import (
	"auth_jwt/internals/utils"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *UserRepo

	jwtSecret string
}

// DTO for register
type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterAuthResult struct {
	Token string `json:"token"`

	Users Users `json:"user"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginAuthResult struct {
	Token string `json:"token"`
	Users Users  `json:"user"`
}

// function to create a new service
func NewUserService(repo *UserRepo, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}
}

// Now we will be using this function in the handler
func (s *Service) RegisterUserService(ctx context.Context, input RegisterInput) (RegisterAuthResult, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email))

	pass := strings.ToLower(strings.TrimSpace(input.Password))

	// check if email or password is empty
	if email == "" || pass == "" {
		slog.Error("Error registering user", "error", "email or password is empty")
		return RegisterAuthResult{}, errors.New("email or password is empty")
	}

	// Checking min and max length of password
	if len(pass) < 8 || len(pass) > 32 {
		slog.Error("Error registering user", "error", "password must be between 8 and 32 characters")
		return RegisterAuthResult{}, errors.New("password must be between 8 and 32 characters")
	}

	// ckeck if email already exists
	_, err := s.repo.FindUserByEmail(ctx, email)
	if err == nil {
		return RegisterAuthResult{}, errors.New("email already exists")
	}

	if !errors.Is(err, mongo.ErrNoDocuments) {
		return RegisterAuthResult{}, err
	}
	// hash the password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error registering user", "error", "error hashing password")
		return RegisterAuthResult{}, errors.New("error hashing password")
	}

	now := time.Now().UTC()

	// create user
	u := Users{
		Email:        email,
		PasswordHash: string(hashBytes),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// create user in database
	createdUser, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		slog.Error("Error registering user", "error", "error creating user")
		return RegisterAuthResult{}, errors.New("error creating user")
	}

	// Create the token for the user
	token, err := utils.CreateToken(s.jwtSecret, createdUser.Id.Hex(), createdUser.Role)
	if err != nil {
		slog.Error("Error registering user", "error", "error creating token")
		return RegisterAuthResult{}, errors.New("error creating token")
	}

	// return the token and the user
	return RegisterAuthResult{
		Token: token,
		Users: createdUser,
	}, nil

}

func (s *Service) LoginUserService(ctx context.Context, input LoginInput, w http.ResponseWriter) (LoginAuthResult, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email))

	pass := strings.ToLower(strings.TrimSpace(input.Password))

	// check if email or password is empty?
	if email == "" || pass == "" {
		slog.Error("Error registering user", "error", "email or password is empty")
		return LoginAuthResult{}, errors.New("email or password is empty")
	}

	// Checking min and max length of password
	if len(pass) < 8 || len(pass) > 32 {
		slog.Error("Error registering user", "error", "password must be between 8 and 32 characters")
		return LoginAuthResult{}, errors.New("password must be between 8 and 32 characters")
	}

	// finding the user
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		slog.Error("Error registering user", "error", "error finding user")
		return LoginAuthResult{}, errors.New("Invalid email or password")
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pass))
	if err != nil {
		slog.Error("Error registering user", "error", "error comparing password")
		return LoginAuthResult{}, errors.New("Invalid email or password")
	}

	// Create the token for the user
	token, err := utils.CreateToken(s.jwtSecret, user.Id.Hex(), user.Role)
	if err != nil {
		slog.Error("Error registering user", "error", "error creating token")
		return LoginAuthResult{}, errors.New("error creating token")
	}

	// set the cookie
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}
	http.SetCookie(w, cookie)

	slog.Info("User logged in", "email", email)

	// return the token and the user
	return LoginAuthResult{
		Token: token,
		Users: user,
	}, nil

}
