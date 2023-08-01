package util

// 使用jwt进行token的生成和解析
import (
	"SkyLine/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims 定义token里面添加的信息
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(user entity.SQLUser) (string, error) {
	//设置token过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	//设置token里面添加的信息
	claims := Claims{
		Username: user.UserName,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			//设置token过期时间
			ExpiresAt: expireTime.Unix(),
			//设置token发放者
			Issuer: "SkyLine",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte("SkyLine"))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("SkyLine"), nil
	})
	if err != nil {
		return nil, err
	}
	//判断token是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
