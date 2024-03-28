package common

import "github.com/gin-gonic/gin"

type Common struct {
}

func NewCommon() *Common {
	return &Common{}
}

func (cm Common) On(c *gin.Context) {
	c.JSON(200, gin.H{"message": "On"})
}

func (cm Common) Off(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Off"})
}
