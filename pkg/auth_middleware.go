package pkg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dan6erbond/go-gqlgen-fx-template/pkg/models"

	// "github.com/dan6erbond/go-gqlgen-fx-template/postgres"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const CurrentUserKey = "currentUser"

func GetUserByID(id string) (*models.User, error) {

	// err := r.db.First(&getuser, "email = ?", input.Email).Error
	// if err == nil {
	// 	return nil, errors.New("email already in used")
	// }

	var user models.User
	// err := u.First(&user, "id = ?", id).Error
	err := DB.First(&user, "id = ?", id).Error
	fmt.Println("error inside GETUsSerID |||||||||||||||||||||||||||||")

	return &user, err

}

func AuthMiddleware() func(http.Handler) http.Handler {
	// func AuthMiddleware(repo models.UsersRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := GetUserByID(claims["jti"].(string))
			if err != nil {
				fmt.Println("error herre inside authorization |||||||||||||||||||||||||||||")
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (s string, err error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error:")
}

func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*models.User)
	if !ok || fmt.Sprint(user.ID) == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}
