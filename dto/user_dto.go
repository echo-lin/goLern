package dto

import "golern/model"

type userDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) userDto {
	return userDto{
		Name:      user.Nickname,
		Telephone: user.Telephone,
	}
}
