package jwt

import (
	"errors"
	"git.legchelife.ru/root/template/pkg/models/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func SecurityJWT(section int) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenBearer := c.GetHeader("Authorization")
		if tokenBearer == "" {
			err := errors.New("authorization required")
			c.AbortWithStatusJSON(401, response.Base{Data: nil, Error: err.Error()})
			return
		}

		//очищаем токен от Bearer и получаем сам токен
		// но сначала проверяем, что он начинается с Bearer
		if len(tokenBearer) < len("Bearer ") {
			err := errors.New("authorization required")
			c.AbortWithStatusJSON(401, response.Base{Data: nil, Error: err.Error()})
			return
		}
		t := tokenBearer[len("Bearer "):]

		token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New("unexpected signing method" + token.Header["alg"].(string))
				return nil, err
			}
			secret, ok := os.LookupEnv("SECRET")
			if !ok {
				err := errors.New("SECRET not found")
				return nil, err
			}
			return []byte(secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, response.Base{Data: nil, Error: err.Error()})
			return
		}

		if token.Valid {
			//добавляем sub в контекст
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

				if (int(claims["uab"].(float64)) & (1 << section)) == 0 {
					err := errors.New("нет доступа к разделу")
					c.AbortWithStatusJSON(401, response.Base{Data: nil, Error: err.Error()})
					return
				}

				payload := Payload{
					ID:            int(claims["id"].(float64)),
					UserAccessBit: int(claims["uab"].(float64)),
					CompanyID:     int(claims["cid"].(float64)),
					DateStart:     int(claims["nbf"].(float64)),
					DateEnd:       int(claims["exp"].(float64)),
				}

				c.Set("payload", payload)
			} else {
				err = errors.New("не удалось обработать токен")
				c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
				return
			}

		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			err = errors.New("ошибка формата токена")
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			err = errors.New("неверная подпись токена")
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			err = errors.New("токен просрочен или еще не активен")
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		} else {
			err = errors.New("не удалось обработать токен")
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}
