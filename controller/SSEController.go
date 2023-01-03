package controller

import (
	"goProject/broadcaster"
	"io"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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
