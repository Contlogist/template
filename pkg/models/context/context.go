package context

import (
	"context"
	"fmt"
	"git.legchelife.ru/root/template/pkg/security/jwt"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

type Base struct {
	Context context.Context
	Cancel  context.CancelFunc
	PID     int         `json:"pID"`
	Payload jwt.Payload `json:"payload"`
}

// Create создает кастомный контекст.
// Контектс содержит:
// - контекст (базовый контекст)
// - PID (ID пользователя)
// - Payload (полезная нагрузка user.Tokens)
func (ctx Base) Create(c *gin.Context) Base {
	ctx.Context = context.Background()

	if c != nil {
		ctx.PID = c.GetInt("pID")
		payload, ok := c.Get("payload")
		if !ok {
			fmt.Println("Warning: 'payload' not found in context")
			ctx.Payload = jwt.Payload{} // или установите ваш дефолтный payload
		} else {
			// Проверяем тип с использованием reflect
			if reflect.TypeOf(payload) == reflect.TypeOf(jwt.Payload{}) {
				ctx.Payload = payload.(jwt.Payload)
			} else {
				fmt.Println("Warning: 'payload' has an unexpected type in context")
				ctx.Payload = jwt.Payload{} // или установите ваш дефолтный payload
			}
		}
	}

	return ctx
}

// SetTimeout устанавливает таймаут для контекста, если он еще не установлен.
// Контекст может быть установлен:
// - в юзкейсе (для всех запросов в юзкейсе)
// - в репозитории (для конкретного запроса)
func (ctx *Base) SetTimeout(second float32) {
	_, ok := ctx.Context.Deadline()
	if !ok {
		tm := time.Duration(second) * time.Second
		ctx.Context, ctx.Cancel = context.WithTimeout(ctx.Context, tm)
	}
}
