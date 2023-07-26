package dao

import (
	"fmt"
	"github.com/yushengguo557/twitter-space/global"
	"github.com/yushengguo557/twitter-space/models"
	"log"
)

// SaveTwitterSpace 保存Space 若存在则更新 若不存在则创建
func SaveTwitterSpace(space models.TwitterSpace) (err error) {
	var tss []models.TwitterSpace
	resp := global.App.DB.Model(&models.TwitterSpace{}).
		Where("id = ?", space.ID).
		Find(&tss)
	if err = resp.Error; err != nil {
		return err
	}
	if resp.RowsAffected > 0 {
		log.Println("Update Space: ", space.ID)
		if err = global.App.DB.Model(&models.TwitterSpace{}).
			Where("id = ?", space.ID).
			Updates(map[string]any{
				"status":            space.Status,
				"participant_count": space.ParticipantCount,
				"ended_at":          space.EndedAt,
			}).Error; err != nil {
			err = fmt.Errorf("update %v, err: %w", space, err)
			return err
		}
	} else {
		log.Println("Create Space: ", space.ID)
		log.Println(space.Lang)
		if err = global.App.DB.Model(&models.TwitterSpace{}).
			Create(&space).
			Error; err != nil {
			err = fmt.Errorf("create space: %v, err: %w", space, err)
			return err
		}
	}
	return nil
}
