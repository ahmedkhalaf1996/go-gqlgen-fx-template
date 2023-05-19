package models

import (
	// "github.com/dan6erbond/go-gqlgen-fx-template/pkg/models"

	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type UsersRepo struct {
	DB *gorm.DB
}

func (u *UsersRepo) GetUserByID(id string) (*User, error) {
	var user User
	err := u.DB.First(&user, "id = ?", id).Error
	return &user, err

}

// func (u *User) HashPassword(password string) error {
// 	bytePassword := []byte(password)
// 	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	u.Password = string(passwordHash)

// 	return nil
// }

// func (u *User) GenToken() (*AuthToken, error) {
// 	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		ExpiresAt: expiredAt.Unix(),
// 		Id:        fmt.Sprint(u.ID),
// 		IssuedAt:  time.Now().Unix(),
// 		Issuer:    "meetmeup",
// 	})

// 	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &AuthToken{
// 		AccessToken: accessToken,
// 		ExpiredAt:   expiredAt,
// 	}, nil
// }

// func (u *User) ComparePassword(password string) error {
// 	bytePassword := []byte(password)
// 	byteHashedPassword := []byte(u.Password)
// 	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
// }
