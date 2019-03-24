package model

import "errors"

var (
	UserNotExist     = errors.New("user not exists")
	ErrInvalidPasswd = errors.New("Passwd or Username is error")
	ErrInvalidParams = errors.New("Invalid params")
)
