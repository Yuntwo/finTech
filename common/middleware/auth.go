package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
	// 实际用到的jwt工具
	"github.com/golang-jwt/jwt/v4"
	// 实际用到的jwt工具，比较旧的工具
	jwt2 "github.com/dgrijalva/jwt-go"
)

// AuthMiddleware 这种JWT验证方法直接调用jwt包的token相关处理方法，不需要自己实现JWT类型，通过JWT实例来调用方法，更简洁
// 鉴权核心参数其实只有一个secret
func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 token
		tokenString := c.GetHeader("Authorization")
		// strings.Index(标准库函数)returns the index of the first instance of the substring, or -1 if the substring is not present
		// 这里是判断token是否以Bearer开头，故与0比较
		if strings.Index(tokenString, "Bearer ") != 0 {
			// 这里定义type H是为了缩写map[string]interface{}
			// 写入响应错误码和信息并设置响应格式(Content-Type)为JSON。Q:不太确定后续是否可以多次写入并覆盖
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "failed to extract token",
			})
			// 不会中断当前handler的执行，但会中断当前handler之后的pending handler的执行
			c.Abort()
			return
		}

		// 解码
		// func(token *jwt.Token) (interface{}, error)是jwt.Parse要求传递的匿名函数，接收一个jwt.Token类型的参数，返回一个interface{}和error类型的值
		//  这个函数receive the parsed token and should return the cryptographic key for verifying the signature
		//  除了用业务定义的密钥验证签名之外，还要验证解析出的token所使用的加密算法(alg)和业务期望的一致
		// Parse方法的解析过程是隐式的，无需业务实现
		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			// 这里是类型转换
			return []byte(secret), nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		// 保存信息
		// Claims用来存储和访问JWT中的payload数据。MapClaims是jwt包自带的简单的key-value结构，没有自定义Claims的严格结构限制，适用于不需要复杂验证的场景
		// MapClaims是一个map[string]interface{}类型，string是key，interface{}(任意类型)是value的map
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["uid"])
			c.Set("payload", claims)
		}

		c.Next()
	}
}

const ErrMsgKey = "errMsg"

// JWTAuth JSON Web Token Auth
// 注意这里JWT和token字符串的区别。JWT只是一种工具(包含属性和方法的类型)，包含对token的各种操作，而token是实际流转的字符串，但不是JWT的属性
// TODO 这里主要做用户的身份验证「用户认证(Authentication)」，即判断是否为合法用户，比如是否过期、是否被篡改等
func JWTAuth() gin.HandlerFunc {
	// Q:gin.Context都有哪些内容，如何传递的？
	// func创建了一个匿名函数(也可以理解为闭包，因为访问/引用了外部变量，Java就不能完美支持闭包)，接收ctx *gin.Context，符合gin.HandlerFunc的签名，因此可以作为gin路由的中间件使用
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			// Q:这里JSON方法的效果是怎样的
			// Q:这里gin.H是什么类型，如何返回到JSON中
			ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Not Authorized."})
			// Q:这里Abort方法的效果是怎样的
			ctx.Abort()
			return
		}

		log.Print("get Authorization： ", token)

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: "Authorization has expired."})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusUnauthorized, gin.H{ErrMsgKey: err.Error()})
			ctx.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		ctx.Set("claims", claims)
	}
}

// JWT 自定义JWT结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	// TokenExpired error是内建的接口类型，errors.New是标准库函数，接受字符串，返回error
	TokenExpired     error = errors.New("token is expired")
	TokenNotValidYet error = errors.New("token not active yet")
	TokenMalformed   error = errors.New("that's not even a token")
	TokenInvalid     error = errors.New("couldn't handle this token")
	// SignKey 自定义用来加密的key
	SignKey string = "Our SecKill Secret Key"
	Issuer  string = "this is a issuer"
)

// CustomClaims 载荷，可以加一些自己需要的信息，用来生成token的实际重要内容
type CustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
	jwt2.StandardClaims
}

// NewJWT 根据SignKey新建一个jwt实例
// Q:这语法不太懂
// Q:这样每次新建的jwt实例不都是一样的吗
func NewJWT() *JWT {
	// 通过结构体字面量(直接指定字段的值)而非new然后赋值来创建实例
	return &JWT{
		// 类型转换，将string类型的SignKey转换为[]byte切片
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	// Q:不太懂这个方法
	token, err := jwt2.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt2.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	// Token异常情况
	if err != nil {
		if ve, ok := err.(*jwt2.ValidationError); ok {
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
	}
	// 这里的token是包下面的*jwt.Token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt2.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt2.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt2.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	// 这里是interface的assertion语法
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
