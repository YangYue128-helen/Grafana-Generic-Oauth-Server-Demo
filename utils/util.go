package utils

import (
	"errors"
)

type User struct {
	ObjectID  string // 请替换成正确的用户属性
	NickName  string // 请替换成正确的用户属性
	Username  string // 请替换成正确的用户属性
}

func GetUserDefault() *User {
	user := &User{
		ObjectID: "123456",
		NickName: "test",
		Username: "test",
	}
	return user
}

func GetUserByID(uid string) (*User, error) {
	if (uid == "123456") {
		return GetUserDefault(), nil;
	}
	return nil, errors.New("User not found")
}

type OAuth2Token struct {
    AccessToken  string `json:"access_token"`  // 访问令牌
    TokenType    string `json:"token_type"`    // 令牌类型
    ExpiresIn    int    `json:"expires_in"`    // 令牌有效期（以秒为单位）
    RefreshToken string `json:"refresh_token"` // 刷新令牌
}