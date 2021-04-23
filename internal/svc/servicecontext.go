package svc

import (
	"fmt"
	"ginweb02/internal/config"
	"ginweb02/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	db,err := gorm.Open(mysql.Open(c.DataSourceName),&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "tech_",
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	err2 := db.AutoMigrate(&models.User{})
	if err2 != nil {
		fmt.Println(err2)
	}
	return &ServiceContext{
		Config: c,
		DbEngin: db,
	}
}
