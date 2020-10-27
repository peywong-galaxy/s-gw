package gin_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/peywong-galaxy/s-gw/log"
	"github.com/peywong-galaxy/s-gw/share"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func AddSequence() gin.HandlerFunc{
	return func(c *gin.Context) {
		log.Logger.Info("gin start to create unique sequence")

		guid := xid.New()
		c.Set(share.UniqueSequenceKey, guid.String())
		c.Set(share.UniqueTimestamp, guid.Time().Unix())
		c.Set(share.UniqueMachine, guid.Machine())
		c.Set(share.UniquePID, guid.Pid())
		c.Set(share.UniqueCounter, guid.Counter())

		log.Logger.Info("gin end in create unique sequence",
			zap.String("guid", guid.String()))

		c.Next()
	}
}
