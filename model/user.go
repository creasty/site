package model

type User struct {
	Base
	Identifiable

	Github string `json:"github"`
	Name   string `json:"name"`
}
