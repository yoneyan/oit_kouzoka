package v0

import (
	"github.com/jinzhu/gorm"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/task"
)

type TaskStore struct {
	DB    *gorm.DB
	Tasks []task.Task
}

func (h *TaskStore) Create(t *task.Task) error {
	return h.DB.Create(&t).Error
}

func (h *TaskStore) Delete(t *task.Task) error {
	return h.DB.Delete(t).Error
}

func (h *TaskStore) GetAll() error {
	var tasks []task.Task
	err := h.DB.Find(&tasks).Error
	if err != nil {
		return err
	}

	h.Tasks = tasks
	return nil
}
