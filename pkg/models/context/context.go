package context

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Base struct {
	Context context.Context
	PID     int
}

func (ctx Base) Create(c *gin.Context) Base {
	ctx.Context = context.Background()
	ctx.PID = c.GetInt("pID")
	return ctx
}

func (ctx Base) SetTimeout(second int) {
	tm := time.Duration(second) * time.Second
	ctx.Context, _ = context.WithTimeout(ctx.Context, tm)
}
