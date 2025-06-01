package utils

import (
	"blogs/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SecretKey string = "74a3324e4f440209aa9d0a9e5a69f7ebbe874ead45ae89a3fd9e251722618c85"

func GenerateToken(user *repository.User) (string, error) {
	// 创建token的声明数据
	claims := jwt.MapClaims{
		"userId":   user.ID,
		"userName": user.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置token过期时间，例如24小时后过期
	}

	// 使用指定的签名算法和密钥创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成编码后的token字符串
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(signedToken string) (*jwt.Token, error) {
	// 解析token并验证签名是否有效
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否是我们所期望的算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
