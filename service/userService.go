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
	GetAllUsers(skip, take int) ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(model.User) (*model.User, error)
	//UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int) error
	GetUsersCount() (int, error)
	Login(model.User) (string, error)
}

func (s *userService) GetAllUsers(skip, take int) ([]model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE deleted = false ORDER BY id desc OFFSET $1 LIMIT $2`, skip, take)
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
		ExpiresAt: time.Now().Add(time.Hour * 100000).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) DeleteUser(userID int) error {
	_, err := database.Client.Exec(`UPDATE users SET deleted = true, updated_at = $1 WHERE id = $2`,
		time.Now(), userID)
	if err != nil {
		//ErrorLog("error inserting format", loger.AdditionalFields{"Error": err, "DbKey": formatToAdd.DbKey})
		return err
	}

	return nil
}

func (s *userService) GetUsersCount() (int, error) {
	count := []int{}
	err := database.Client.Select(&count, `SELECT COUNT(id) FROM users WHERE deleted = false`)
	if err != nil || len(count) == 0 {
		//ErrorLog("error inserting format", loger.AdditionalFields{"Error": err, "DbKey": formatToAdd.DbKey})
		return 0, err
	}

	return count[0], nil
}
