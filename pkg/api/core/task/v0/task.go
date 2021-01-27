package v0

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/auth"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/task"
	dbTask "github.com/yoneyan/oit_kouzoka/pkg/api/store/task/v0"
	"log"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(database *gorm.DB) *TaskHandler {
	return &TaskHandler{db: database}
}

func (t *TaskHandler) Create(c *gin.Context) {
	var input task.Incoming
	err := c.BindJSON(&input)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if !auth.Authentication(input.Token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	mStore := dbTask.TaskStore{DB: t.db}
	if err := mStore.Create(&task.Task{Name: input.Name, Task: input.Task}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success !!"})
}

func (t *TaskHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not number"})
		return
	}

	var input task.Incoming
	err = c.BindJSON(&input)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if !auth.Authentication(input.Token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	mStore := dbTask.TaskStore{DB: t.db}
	if err := mStore.Delete(&task.Task{Model: gorm.Model{ID: uint(id)}}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success !!"})
}
func (t *TaskHandler) Get(c *gin.Context) {
	token := c.Query("token")
	if !auth.Authentication(token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	mStore := dbTask.TaskStore{DB: t.db}
	if err := mStore.GetAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, mStore.Tasks)
}
