package usecase

import (
	"fmt"

	"github.com/ValeryBMSTU/web-11/internal/auth/provider"
)

func (u *Usecase) Authenticate(username, password string) (string, error) {
	exist, err := u.p.CheckUserByUsername(username)
	if !exist {
		return "", fmt.Errorf("user not found")
	}
	if err != nil {
		return "", err
	}

	if correct, _ := u.p.CheckPassword(username, password); !correct {
		return "", fmt.Errorf("invalid password")
	}

	return u.jp.GenerateToken(username)
}

func (u *Usecase) ValidateJWT(token string) (*provider.JWTClaims, error) {
	return u.jp.ValidateToken(token)
}

func (u *Usecase) Register(login, password string) error {
	exist, err := u.p.CheckUserByUsername(login)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("user already exists")
	}

	return u.p.CreateUser(login, password)
}

// func (u *Usecase) FetchHelloMessage() (string, error) {
// 	msg, err := u.p.SelectRandomHello()
// 	if err != nil {
// 		return "", err
// 	}

// 	if msg == "" {
// 		return u.defaultMsg, nil
// 	}

// 	return msg, nil
// }

// func (u *Usecase) SetHelloMessage(msg string) error {
// 	isExist, err := u.p.CheckHelloExitByMsg(msg)
// 	if err != nil {
// 		return err
// 	}

// 	if isExist {
// 		return nil
// 	}

// 	err = u.p.InsertHello(msg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
