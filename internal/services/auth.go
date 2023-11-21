package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/internal/dto"
	"github.com/phamdinhha/go-chat-server/internal/models"
	"github.com/phamdinhha/go-chat-server/internal/repositories"
	"github.com/phamdinhha/go-chat-server/pkg/http_error"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	userRepo repositories.UserRepo
	cfg      *config.Config
}

func NewAuthService(
	userRepo repositories.UserRepo,
	cfg *config.Config,
) AuthenService {
	return &auth{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (a *auth) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}
	expriedAt := time.Now().Add(time.Hour * time.Duration(a.cfg.JWTConfig.JWTTTL)).Unix()
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if errf != nil {
		return dto.LoginResponse{}, http_error.ErrInvalidCredentials
	}
	tk := &models.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expriedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(a.cfg.JWTConfig.SigningMethod), tk)
	tokenString, err := token.SignedString([]byte(a.cfg.JWTConfig.JWTSecret))
	if err != nil {
		return dto.LoginResponse{}, err
	}
	return dto.LoginResponse{
		User:     user,
		JwtToken: tokenString,
	}, nil
}

func (a *auth) SignUp(ctx context.Context, req dto.SignUpRequest) (dto.SignUpResponse, error) {
	userCheck, _ := a.userRepo.GetUserByEmail(ctx, req.Email)
	if userCheck != nil {
		return dto.SignUpResponse{}, http_error.ErrDuplicateEmail
	}
	expiredAt := time.Now().Add(time.Hour * time.Duration(a.cfg.JWTConfig.JWTTTL)).Unix()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.SignUpResponse{}, err
	}
	hPS := string(hashedPassword)
	user := &models.User{
		ID:       req.Email,
		UserName: req.UserName,
		Email:    req.Email,
		Password: hPS,
	}
	user, err = a.userRepo.CreateUser(ctx, user)
	if err != nil {
		return dto.SignUpResponse{}, err
	}
	user.Password = ""
	tk := &models.Token{
		UserID:   user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(a.cfg.JWTConfig.SigningMethod), tk)
	jwtSecret := []byte(a.cfg.JWTConfig.JWTSecret)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return dto.SignUpResponse{}, err
	}
	return dto.SignUpResponse{
		User:     user,
		JwtToken: tokenString,
	}, nil
}
