package usecase

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/dgrijalva/jwt-go/v4"

	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *entity.User `json:"user"`
}

type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLScnd time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLScnd,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, l logger.Interface, username, password string) error {
	//pwd := sha1.New()
	fmt.Println("uc-signUp()-username-passwors", username, password)
	//pwd.Write([]byte(password))
	//pwd.Write([]byte(a.hashSalt))

	user := &entity.User{
		Login: username,
		//Password: fmt.Sprintf("%x", pwd.Sum(nil)),
		Password: password,
	}

	return a.userRepo.CreateUser(ctx, l, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, l logger.Interface, username, password string) (string, error) {
	//pwd := sha1.New()
	//pwd.Write([]byte(password))
	//pwd.Write([]byte(a.hashSalt))
	//password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, l, username, password)
	fmt.Println("auth-uc-user: ", err)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, l logger.Interface, accessToken string) (*entity.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			l.Info("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		l.Info("err in jwt.ParseWithClaims()")
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
