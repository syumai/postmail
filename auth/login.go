// MIT license (c) andelf 2013

package auth

import (
	"errors"
	"net/smtp"
)

type loginAuth struct {
	userName, password string
}

func LoginAuth(userName, password string) smtp.Auth {
	return &loginAuth{userName, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.userName), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

