package services

import (
	"database/sql"
	"errors"
	"49hw/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) RegisterUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.Exec(query, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	query := `SELECT id, username, email, role FROM users WHERE username=$1 AND password=$2`
	row := s.DB.QueryRow(query, username, password)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	return &user, nil
}

func CreateToken(username, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key"))
}
