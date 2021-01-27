package api

import (
	"github.com/gin-gonic/gin"
	task "github.com/yoneyan/oit_kouzoka/pkg/api/core/task/v0"
	"github.com/yoneyan/oit_kouzoka/pkg/api/tool/config"
	"log"
	"net/http"
	"strconv"
)

func RestAPI() {
	log.Println("Server Port: " + strconv.Itoa(config.Conf.Port))

	router := gin.Default()
	router.Use(cors)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// Task
			v1.POST("/task", task.Create)
			v1.DELETE("/task/:id", task.Delete)
			v1.GET("/task", task.Get)
		}
	}
	err := http.ListenAndServe(":"+strconv.Itoa(config.Conf.Port), router)
	if err != nil {
		log.Println(err)
	}
}

func cors(c *gin.Context) {

	//c.Header("Access-Control-Allow-Headers", "Accept, Content-ID, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-ID", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
