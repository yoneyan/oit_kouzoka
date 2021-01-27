package task

type Task struct {
	//gorm.Model
	ID        uint   `json:"id" xml:"id"`
	CreatedAt string `json:"createdAt" xml:"createdAt"`
	Name      string `json:"name" xml:"name"`
	Task      string `json:"task" xml:"task"`
}

type Tasks struct {
	Task []Task `json:"tasks" xml:"tasks"`
}

type Incoming struct {
	Token string `json:"token" xml:"token"`
	Name  string `json:"name" xml:"name"`
	Task  string `json:"task" xml:"task"`
}
