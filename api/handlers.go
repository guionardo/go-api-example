package api

import "github.com/gin-gonic/gin"

type HTTPHandlers interface {
	GetFeira(c *gin.Context)
	GetFeiras(c *gin.Context)
	CreateFeira(c *gin.Context)
	DeleteFeira(c *gin.Context)
	UpdateFeira(c *gin.Context)
}
