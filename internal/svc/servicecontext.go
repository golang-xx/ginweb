package svc

import (
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

	db.AutoMigrate(&models.User{})

	return &ServiceContext{
		Config: c,
		DbEngin: db,
	}
}
