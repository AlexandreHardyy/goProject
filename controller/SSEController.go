package controller

import (
	"io"

	"github.com/AlexandreHardyy/goProject/broadcaster"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Stream godoc
// @Summary Stream payments
// @Description Stream payments
// @Tags payments
// @Accept  json
// @Produce  json
// @Success 200 {object} Payment
// @Router /payments/stream [get]
func Stream(c *gin.Context) {
	listener := make(chan interface{})
	broadcaster := broadcaster.GetBroadcaster()

	broadcaster.Register(listener)
	defer broadcaster.Unregister(listener)

	clientGone := c.Request.Context().Done()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case payment := <-listener:
			c.SSEvent("payment", payment)
			return true
		}
	})
}
