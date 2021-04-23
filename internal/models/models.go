package models

import (
	"errors"
	"ginweb02/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Mobile string `gorm:"index:mobile;type:varchar(13)"`
	Passwd string `gorm:"type:vachar(64)"`
}

func (u *User)BeforeCreate(db *gorm.DB) error  {

	if len(u.Passwd) < 6 {
		return errors.New("秘密太简单了")
	}

	u.Passwd = utils.Password(u.Passwd)
	return nil
}