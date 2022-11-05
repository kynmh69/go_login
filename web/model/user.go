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
	logger := logging.GetLogger()

	user := User{}
	row := Db.QueryRow("select * from user where user_id = ?", userId)
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

	result, err := Db.Exec("insert into user (id, user_id, password) values(?, ?, ?)", user.Id, user.UserId, user.Password)
	logger.Debug(result.LastInsertId())

	return &user, err
}
