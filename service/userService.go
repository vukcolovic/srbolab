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
	UpdateUser(model.User) (*model.User, error)
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

	for _, u := range users {
		u.Password = ""
	}

	return users, nil
}

func (s *userService) CreateUser(user model.User) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		loger.ErrorLog.Println("Error creating user: ", err)
		return nil, err
	}
	_, err = database.Client.Exec(`INSERT INTO users (first_name, last_name, email, phone_number, contract_number, contract_type, jmbg, adress, started_work, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`,
		user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.ContractNumber, user.ContractType, user.JMBG, user.Adress, user.StartedWork, string(hashedPassword), time.Now(), time.Now())
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

	user := &users[0]
	user.Password = ""
	return user, nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil || len(users) == 0 {
		loger.ErrorLog.Println("Error getting user by email: ", err)
		return nil, err
	}

	user := &users[0]
	user.Password = ""
	return user, nil
}

func (s *userService) Login(userJustCredentials model.User) (*model.LoginResponse, error) {
	users := []model.User{}
	err := database.Client.Select(&users, `SELECT * FROM users WHERE email = $1`, userJustCredentials.Email)
	if err != nil || len(users) == 0 {
		loger.ErrorLog.Println("Error getting user by email: ", err)
		return nil, err
	}

	user := &users[0]

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

func (s *userService) UpdateUser(user model.User) (*model.User, error) {
	//TODO put all in one transaction
	_, err := database.Client.Exec(`UPDATE users SET first_name = $1, last_name = $2, email = $3, phone_number = $4, contract_number = $5, contract_type = $6, jmbg = $7, adress = $8, started_work = $9,  updated_at = $10 WHERE id = $11`,
		user.FirstName, user.LastName, user.Email, time.Now(), user.Id)
	if err != nil {
		loger.ErrorLog.Println("Error updating user: ", err)
		return nil, err
	}

	if user.Password != "" {
		users := []model.User{}
		err = database.Client.Select(&users, `SELECT * FROM users WHERE id = $1`, user.Id)
		if err != nil || len(users) == 0 {
			loger.ErrorLog.Println("Error getting user by id: ", err)
			return nil, err
		}
		oldUser := users[0]

		if err = bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(user.CurrentPassword)); err != nil {
			loger.ErrorLog.Println("error updating user, error comparing hashes: ", err)
			return nil, errors.New("pogresna trenutna sifra")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			loger.ErrorLog.Println("Error updating user, problem generate hash from a new password: ", err)
			return nil, err
		}

		_, err = database.Client.Exec(`UPDATE users SET password = $1, updated_at = $2 WHERE id = $3`,
			hashedPassword, time.Now(), user.Id)
		if err != nil {
			loger.ErrorLog.Println("Error updating user password: ", err)
			return nil, err
		}
	}

	//todo return that user if there is need for that
	return nil, err
}
