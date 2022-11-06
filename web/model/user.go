package model

import (
	"errors"
	"go_login/logging"
	"go_login/utils"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string
	UserId    string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SignUp(userId, password string) (*User, error) {
	db := ConnectDb()
	logger := logging.GetLogger()

	user := User{}
	row := db.QueryRow("select * from user where user_id = ?", userId)
	row.Scan(&user.Id, &user.UserId, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if user.UserId != "" {
		logger.Error("Already exists user id.", user.UserId)
		err := errors.New("既にユーザネームは使用されています。")
		return nil, err
	}

	encryptPw, err := utils.PasswordEncrypt(password)
	uuidObj := uuid.New()
	if err != nil {
		logger.Fatalln("not password encrypt.")
	}
	user = User{
		Id:       uuidObj.String(),
		UserId:   userId,
		Password: encryptPw,
	}
	stmt, _ := db.Prepare("insert into user (id, user_id, password) values(?, ?, ?)")
	result, err := stmt.Exec(user.Id, user.UserId, user.Password)
	logger.Debug(result.LastInsertId())

	defer db.Close()

	return &user, err
}

func Login(userId, password string) (*User, error) {
	logger := logging.GetLogger()
	db := ConnectDb()
	user := User{}

	row := db.QueryRow("select user_id, password from user where user_id = ?", userId)

	err := row.Scan(&user.UserId, &user.Password)
	if err != nil {
		logger.Warn("User not found.", err.Error())
		errorMsg := errors.New("ユーザが存在しません")
		return nil, errorMsg
	}

	err = utils.CompareHashAndPassword(user.Password, password)

	if err != nil {
		logger.Warn("wrong passsword.")
		return nil, err
	}

	defer db.Close()

	return &user, nil
}

func GetOneUser(userId string) (*User, error) {
	logger := logging.GetLogger()
	db := ConnectDb()
	user := User{}

	logger.Info("get User.", userId)

	row := db.QueryRow("select user_id, password from user where user_id = ?", userId)

	err := row.Scan(&user.UserId, &user.Password)

	if err != nil {
		logger.Error("not found user:", userId)
		return nil, err
	}
	return &user, nil
}
