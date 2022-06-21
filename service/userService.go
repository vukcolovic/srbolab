package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"srbolabApp/database"
	"srbolabApp/model"
	"strconv"
	"time"
)

var (
	UsersService usersServiceInterface = &userService{}
)

type userService struct {
}

type usersServiceInterface interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(model.User) (*model.User, error)
	//UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	//DeleteUser(int64) rest_errors.RestErr
	//Search(string) (users.Users, rest_errors.RestErr)
	Login(model.User) (string, error)
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (s *userService) CreateUser(user model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	_, err = database.Client.Exec(`INSERT INTO users (first_name, last_name, email, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`,
		user.FirstName, user.LastName, user.Email, string(hashedPassword), time.Now(), time.Now())
	if err != nil {
		//loger.Instance().Error("error inserting format", loger.AdditionalFields{"Error": err, "DbKey": formatToAdd.DbKey})
		return nil, err
	}

	return UsersService.GetUserByEmail(user.Email)
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil || len(users) == 0 {
		log.Println(err)
		return nil, err
	}

	return &users[0], nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil || len(users) == 0 {
		log.Println(err)
		return nil, err
	}

	return &users[0], nil
}

func (s *userService) Login(userJustCredentials model.User) (string, error) {
	user, err := UsersService.GetUserByEmail(userJustCredentials.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		//FIXME
		return "", errors.New("no user with email " + userJustCredentials.Email)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userJustCredentials.Password)); err != nil {
		return "", err
	}

	claims := jwt.StandardClaims{
		Id:        strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
