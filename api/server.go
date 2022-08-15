package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guionardo/go-api-example/docs"
	"github.com/guionardo/go-api-example/infra"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	port   int
	host   string
}

func NewServer(config *infra.Config) *Server {
	server := &Server{
		router: gin.Default(),
		port:   config.HttpPort,
		host:   config.HttpHost,
	}
	server.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	return server
}

func (server *Server) Router() *gin.Engine {
	return server.router
}

func (server *Server) RegisterRoutes(t HTTPHandlers) {
	docs.SwaggerInfo.Title = "Go API Example"

	group := server.router.Group("/feiras")
	group.GET("/", t.GetFeiras)
	group.GET("/:registro", t.GetFeira)
	group.POST("/", t.CreateFeira)
	group.PUT("/", t.UpdateFeira)
	group.DELETE("/:registro", t.DeleteFeira)

	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL(fmt.Sprintf("%s:%d/swagger/doc.json", server.host, server.port)),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

}

func (server *Server) Start() {
	log.Printf("Server is running on port %d", server.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.port), server.router))
}
