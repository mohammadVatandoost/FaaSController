package Jobs

import (
	"controller/internal/messaging"

	"github.com/gin-gonic/gin"
)

// Routes ...
func (jobCon *Controller) Routes(group *gin.RouterGroup, messagingService *messaging.Service) {
    jobCon.Messaging = messagingService
	group.POST("/job", recordCon.jobHandler)
}
