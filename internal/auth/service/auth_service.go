package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/davidPardoC/go-chat/cmd/chat/api/http/dtos"
	"github.com/davidPardoC/go-chat/internal/auth/model"
	"github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/davidPardoC/go-chat/pkg/constants"
	"github.com/davidPardoC/go-chat/pkg/errs"
	"github.com/davidPardoC/go-chat/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	userModels "github.com/davidPardoC/go-chat/internal/user/model"
)

type AuthService struct {
	userRepo repository.IUserRepository
}

func NewAuthService(userRepo repository.IUserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) LoginUser(email string, password string) (*model.Credentials, *errs.Error) {
	user, err := s.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, errs.NewUnauthorizedError("Not user foud")
	}

	if isValid := utils.VerifyPassword(password, user.Password); !isValid {
		return nil, errs.NewUnauthorizedError("Invalid Credndetials")
	}

	accesTokenClaims := utils.JwtClaims{
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			Issuer:    constants.TOKEN_ISSUER,
			Subject:   fmt.Sprint(user.ID),
		},
	}

	refreshTokenClaims := utils.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			Issuer:    constants.TOKEN_ISSUER,
		},
	}

	accesToken, err := utils.SignToken(accesTokenClaims)
	refreshToken, refreshError := utils.SignToken(refreshTokenClaims)

	if err != nil || refreshError != nil {
		return nil, errs.NewUnauthorizedError("Invalid Credndetials")
	}

	user.RefreshToken = &sql.NullString{String: refreshToken}

	s.userRepo.UpdateRefresToken(*user, refreshToken)

	return &model.Credentials{AccesToken: accesToken, RefreshToken: refreshToken}, nil
}

func (s *AuthService) SignupUser(signupDto dtos.SignUpDto) (*userModels.User, *errs.Error) {
	hashedPassword, _ := utils.HashPassword(signupDto.Password)

	user := userModels.User{Username: signupDto.Username, Email: signupDto.Email, Password: hashedPassword}
	resul, err := s.userRepo.CreateUser(&user)

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errs.NewConflictError("Email alredy in use")
	}

	return resul.WithoutPassword(), nil
}
