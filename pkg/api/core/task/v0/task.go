package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/auth"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/task"
	"github.com/yoneyan/oit_kouzoka/pkg/api/store"
	"log"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	var input task.Incoming
	err := c.BindXML(&input)
	if err != nil {
		log.Println(err)
		c.XML(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if !auth.Authentication(input.Token) {
		c.XML(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	err = store.Save(task.Task{Name: input.Name, Task: input.Task})
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.XML(http.StatusOK, gin.H{"status": "Success !!"})
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{"error": "id is not number"})
		return
	}

	var input task.Incoming
	err = c.BindXML(&input)
	if err != nil {
		log.Println(err)
		c.XML(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if !auth.Authentication(input.Token) {
		c.XML(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	err = store.Delete(uint(id))
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.XML(http.StatusOK, gin.H{"status": "Success !!"})
}
func Get(c *gin.Context) {
	token := c.Query("token")
	if !auth.Authentication(token) {
		c.XML(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	tasks, err := store.Read()
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.XML(http.StatusOK, tasks)
}
