package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

func (userService *UserService) LoginEmail(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.GVA_DB.Where("email = ?", u.Email).Preload("Authorities").Preload("Authority").First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_, err := userService.Register(*u)
		if err != nil {
			return &user, err
		}
		err = global.GVA_DB.Where("email = ?", u.Email).Preload("Authorities").Preload("Authority").First(&user).Error
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&user)
	return &user, err
}
