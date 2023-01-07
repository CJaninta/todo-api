package configs

import (
	"fmt"
	"todo/repositories"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"))
	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	project := repositories.Project{}
	user := repositories.User{}
	todo := repositories.Todo{}

	db.AutoMigrate(project, user, todo)

	return db
}
