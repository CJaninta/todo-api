package repositories

type Project struct {
	BaseModel
	Name   string
	Detail string
	Todos  []Todo `gorm:"foreignKey:ProjectID"`
	Users  []User `gorm:"many2many:user_project"`
}

type ProjectRepository interface {
	GetAll() ([]Project, error)
	GetById(uint) (*Project, error)
	GetByName(string) (*Project, error)
	GetProjectByUser(int) ([]Project, error)
	Create(Project) (*Project, error)
}
