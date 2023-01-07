package repositories

type Todo struct {
	BaseModel
	Title     string
	Detail    string
	Step      string `gorm:"default:OPEN"`
	UserID    uint
	ProjectID uint
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetById(int) (*Todo, error)
	GetByProjectId(int) ([]Todo, error)
	Create(Todo) (*Todo, error)
	UpdateStep(int, string) error
}
