package base

import "os/user"

type Question struct {
	Direction string   `json:"direction"`
	Images    []string `json:"images"`
}

type CreatedQuestion interface {
	GetCreator() *user.User
}
