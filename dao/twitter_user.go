package dao

import (
	"fmt"
	"github.com/yushengguo557/twitter-space/global"
	"github.com/yushengguo557/twitter-space/models"
	"log"
)

func SaveTwitterUser(user models.TwitterUser) (err error) {
	var tus []models.TwitterUser
	resp := global.App.DB.Model(&models.TwitterUser{}).Where("id = ?", user.ID).Find(&tus)
	if err = resp.Error; err != nil {
		return err
	}
	if resp.RowsAffected > 0 {
		log.Println("Update User: ", user.ID)
		if err = global.App.DB.Model(&models.TwitterUser{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
			err = fmt.Errorf("update %v, err: %w", user, err)
			return err
		}
	} else {
		log.Println("Create User", user.ID)
		if err = global.App.DB.Model(&models.TwitterUser{}).Create(&user).Error; err != nil {
			err = fmt.Errorf("create %v, err: %w", user, err)
			return err
		}
	}
	return nil
}
