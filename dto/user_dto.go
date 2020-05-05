package dto

import "go_gin_second/model"

//UserDto 用户数据传输模型
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//ToUserDto 转为传输User信息
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
