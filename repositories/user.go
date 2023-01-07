package repositories

type User struct {
	BaseModel
	FirstName string
	LastName  string
	Email     string
	Password  string
	Active    bool      `gorm:"default:true"`
	Todos     []Todo    `gorm:"foreignKey:UserID"`
	Projects  []Project `gorm:"many2many:user_project"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(uint) (*User, error)
	GetByEmail(string) (*User, error)
	Create(User) (*User, error)
}
