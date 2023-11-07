package utils

import (
	"errors"
)
type Role string;

const (
	Admin     Role = "admin"
	Developer Role = "developer"
	Viewer    Role = "viewer"
)

type User struct {
	ObjectID  string // 请替换成正确的用户属性
	NickName  string // 请替换成正确的用户属性
	Username  string // 请替换成正确的用户属性
	Roles     []Role 
}

func NewUser(id string, username string, nickName string, roles []Role) *User {
    user := &User{
        ObjectID: id,
		NickName: nickName,
		Username: username,
        Roles:    roles,
    }
	return user
}

func GetUserByID(uid string) (*User, error) {
	if (uid == "123456") {
		return NewUser("123456", "test", "test", []Role{Admin}), nil
	} else {
		return NewUser("234567", "test1", "test1", []Role{Developer, Viewer}), nil
	}
	return nil, errors.New("User not found")
}

type OAuth2Token struct {
    AccessToken  string `json:"access_token"`  // 访问令牌
    TokenType    string `json:"token_type"`    // 令牌类型
    ExpiresIn    int    `json:"expires_in"`    // 令牌有效期（以秒为单位）
    RefreshToken string `json:"refresh_token"` // 刷新令牌
}