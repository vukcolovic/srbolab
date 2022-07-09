package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"srbolabApp/database"
	"srbolabApp/loger"
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
	GetUserIDByToken(token string) (int, error)
	GetAllUsers(skip, take int) ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(model.User) (*model.User, error)
	//UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int) error
	GetUsersCount() (int, error)
	Login(model.User) (*model.LoginResponse, error)
}

func (s *userService) GetUserIDByToken(token string) (int, error) {
	claims := &jwt.StandardClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		loger.ErrorLog.Println("Error getting user by token, error parse claims: ", err)
		return 0, err
	}

	id, err := strconv.Atoi(claims.Id)
	if err != nil {
		loger.ErrorLog.Println("Error getting user by token: ", err)
		return 0, err
	}

	return id, nil
}

func (s *userService) GetAllUsers(skip, take int) ([]model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE deleted = false ORDER BY id desc OFFSET $1 LIMIT $2`, skip, take)
	if err != nil {
		loger.ErrorLog.Println("Error getting all users: ", err)
		return nil, err
	}

	return users, nil
}

func (s *userService) CreateUser(user model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		loger.ErrorLog.Println("Error creating user: ", err)
		return nil, err
	}
	_, err = database.Client.Exec(`INSERT INTO users (first_name, last_name, email, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`,
		user.FirstName, user.LastName, user.Email, string(hashedPassword), time.Now(), time.Now())
	if err != nil {
		loger.ErrorLog.Println("Error creating user: ", err)
		return nil, err
	}

	return UsersService.GetUserByEmail(user.Email)
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil || len(users) == 0 {
		loger.ErrorLog.Println("Error getting user by id: ", err)
		return nil, err
	}

	return &users[0], nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil || len(users) == 0 {
		loger.ErrorLog.Println("Error getting user by email: ", err)
		return nil, err
	}

	return &users[0], nil
}

func (s *userService) Login(userJustCredentials model.User) (*model.LoginResponse, error) {
	user, err := UsersService.GetUserByEmail(userJustCredentials.Email)
	if err != nil {
		loger.ErrorLog.Println("Error login user: ", err)
		return nil, err
	}
	if user == nil {
		loger.ErrorLog.Println("Error login user, no user with specified email")
		return nil, errors.New("no user with email " + userJustCredentials.Email)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userJustCredentials.Password)); err != nil {
		loger.ErrorLog.Println("Error login user, error comparing hashes: ", err)
		return nil, err
	}

	claims := jwt.StandardClaims{
		Id:        strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 100000).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		loger.ErrorLog.Println("Error login user, error signing token: ", err)
		return nil, err
	}

	return &model.LoginResponse{Token: token, FirstName: user.FirstName, LastName: user.LastName}, nil
}

func (s *userService) DeleteUser(userID int) error {
	_, err := database.Client.Exec(`UPDATE users SET deleted = true, updated_at = $1 WHERE id = $2`,
		time.Now(), userID)
	if err != nil {
		loger.ErrorLog.Println("Error deleting user: ", err)
		return err
	}

	return nil
}

func (s *userService) GetUsersCount() (int, error) {
	count := []int{}
	err := database.Client.Select(&count, `SELECT COUNT(id) FROM users WHERE deleted = false`)
	if err != nil || len(count) == 0 {
		loger.ErrorLog.Println("Error getting count of users: ", err)
		return 0, err
	}

	return count[0], nil
}
