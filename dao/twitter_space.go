package dao

import (
	"fmt"
	"log"
	"twitter-space/global"
	"twitter-space/models"
)

// SaveTwitterSpace 保存Space 若存在则更新 若不存在则创建
func SaveTwitterSpace(space models.TwitterSpace) (err error) {
	var sps []models.TwitterSpace
	resp := global.DB.Model(&models.TwitterSpace{}).Where("id = ?", space.ID).Find(&sps)
	if err = resp.Error; err != nil {
		return err
	}
	if resp.RowsAffected > 0 {
		log.Println("Update Space: ", space.ID)
		if err = global.DB.Model(&models.TwitterSpace{}).Where("id = ?", space.ID).Updates(&space).Error; err != nil {
			err = fmt.Errorf("update %v, err: %w", space, err)
			return err
		}
	} else {
		log.Println("Create Space: ", space.ID)
		log.Println(space.Lang)
		if err = global.DB.Model(&models.TwitterSpace{}).Create(&space).Error; err != nil {
			err = fmt.Errorf("create space: %v, err: %w", space, err)
			return err
		}
	}
	return nil
}
