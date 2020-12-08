package Jobs

import (
	"controller/internal/messaging"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller ...
type Controller struct {
	Messaging *messaging.Service
}

type CacheItem {

}

type JobReq struct {
	ID        string   `json:"id"`
	Variables []string `json:"variables"`
}

func (jobCon Controller) jobHandler(c *gin.Context) {
	var jobReq JobReq
	err := c.BindJSON(&jobReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

    


	c.JSON(http.StatusOK, "")
}
