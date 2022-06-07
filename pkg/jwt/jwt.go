package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtClaims 创建自己的Claims
type JwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	UID int64
}

var (
	//盐
	secret                 = []byte("wondersafebox")                   // 后续加密增加盐增加复杂度
	TokenExpired     error = errors.New("Token is expired")            // token失效
	TokenNotValidYet error = errors.New("Token not active yet")        // token未激活
	TokenMalformed   error = errors.New("That's not even a token")     // 非token格式
	TokenInvalid     error = errors.New("Couldn't handle this token:") // 未知错误
)

// CreateJwtToken 生成一个jwttoken
func CreateJwtToken(userid int64) (string, error) {

	// 定义过期时间,1天后过期
	expireToken := time.Now().Add(time.Hour * 24).Unix()

	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // token信息生效时间
			ExpiresAt: int64(expireToken),              // 过期时间
		},
		userid,
	}
	// 对自定义claims加密,jwt.SigningMethodHS256是加密算法得到第二部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 给这个token盐加密 第三部分,得到一个完整的三段的加密
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ParseJwtToken 解析token得到是自己创建的Claims
func ParseJwtToken(jwtToken string) (*JwtClaims, error) {
	var jwtclaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, jwtclaim, func(*jwt.Token) (interface{}, error) {
		//得到盐
		return secret, nil
	})
	if err != nil {
		// 细化token解析错误
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		return nil, err
	}
	return jwtclaim, nil
}
