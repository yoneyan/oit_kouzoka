package task

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name   string `json:"name" xml:"name"`
	Task   string `json:"task" xml:"task"`
	Finish *bool  `json:"finish" xml:"finish"`
}

type Incoming struct {
	Token string `json:token`
	Name  string `json:"name" xml:"name"`
	Task  string `json:"task" xml:"task"`
}
