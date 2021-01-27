package store

import (
	"encoding/xml"
	"fmt"
	"github.com/yoneyan/oit_kouzoka/pkg/api/core/task"
	"github.com/yoneyan/oit_kouzoka/pkg/api/tool/config"
	_ "gorm.io/driver/sqlite"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func Read() (*task.Tasks, error) {
	var tasks task.Tasks
	data, _ := ioutil.ReadFile(config.Conf.XMLPath)
	err := xml.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func Delete(id uint) error {
	var tasks task.Tasks

	data, err := Read()
	if err != nil {
		return err
	}

	for _, tmp := range data.Task {
		if tmp.ID != id {
			tasks.Task = append(tasks.Task, tmp)
		}
	}
	b, _ := xml.MarshalIndent(&tasks, "", "\t")
	output, _ := os.Create(config.Conf.XMLPath)
	defer output.Close()

	output.Write(([]byte)(xml.Header))
	output.Write(b)

	return nil
}

func Save(input task.Task) error {
	data, _ := Read()

	var tasks task.Tasks

	var id uint = 1
	if data != nil {
		tasks.Task = data.Task
		for _, tmp := range data.Task {
			log.Println(tmp.ID, id)
			if tmp.ID >= id {
				id = tmp.ID + 1
				log.Println(id)
			}
		}
	}

	now := time.Now()
	timeNow := fmt.Sprintf("%s", now.Format("2006-01-02T15:04:05+09:00"))

	tasks.Task = append(tasks.Task, task.Task{ID: id, CreatedAt: timeNow, Name: input.Name, Task: input.Task})
	b, _ := xml.MarshalIndent(&tasks, "", "\t")
	output, _ := os.Create(config.Conf.XMLPath)
	defer output.Close()

	output.Write(([]byte)(xml.Header))
	output.Write(b)

	return nil
}
