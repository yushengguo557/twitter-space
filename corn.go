package main

import (
	"log"
	"time"
	"twitter-space/dao"
	"twitter-space/global"
	"twitter-space/models"
)

// Corn 周期任务
func Corn() {
	go TimedSearchSpace()
	go TimedLookupUser()
	go TimedUpdateSpace()
}

// TimedUpdateSpace 定时更新 Space
func TimedUpdateSpace() {
	var err error
	var ids []string
	for range time.NewTicker(time.Minute * 2).C {
		var total int64
		if err = global.DB.Model(&models.TwitterSpace{}).
			Where("`status` IN ? AND data_status = ?", []string{"live", "scheduled"}, models.DataStatusEnable).
			Count(&total).Error; err != nil {
			log.Println(err)
		}
		offset := 0
		for offset < int(total) {
			if err = global.DB.Model(&models.TwitterSpace{}).
				Select("id").
				Where("`status` IN ? AND data_status = ?", []string{"live", "scheduled"}, models.DataStatusEnable).
				Limit(100).
				Offset(offset).
				Find(&ids).Error; err != nil {
				log.Println(err)
			}
			var spaces []models.TwitterSpace
			spaces, err = global.Client.SpaceLookup(ids)
			if err != nil {
				log.Println(err)
			}
			for _, space := range spaces {
				err = dao.SaveTwitterSpace(space)
				if err != nil {
					log.Println(err)
				}
			}
			offset += 100
		}
	}
}

// TimedSearchSpace 定时搜索Space
func TimedSearchSpace() {
	// 每隔 Period 搜索一次
	for range time.NewTicker(Period).C {
		// TODO: 搜索Space
		querys := []string{models.SpaceNFT, models.SpaceWEB3, models.SpaceMetaVerse, models.SpaceGame, models.SpaceDeFi}
		for _, query := range querys {
			spaces, err := global.Client.SpaceSearch(query)
			if err != nil {
				log.Println(err)
			}
			// TODO: 保存Space
			for _, space := range spaces {
				err = dao.SaveTwitterSpace(space)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

// TimedLookupUser 定时获取正在直播的Space中的用户信息
func TimedLookupUser() {
	var err error
	for range time.NewTicker(time.Minute * 2).C {
		// TODO: 清空用户所属Space
		if err = global.DB.Model(&models.TwitterUser{}).
			Where("data_status = ?", models.DataStatusEnable).
			Updates(map[string]any{"space_id": nil}).
			Error; err != nil {
			log.Println(err)
		}

		// TODO: 更新用户所属Space
		var ids []string
		if err = global.DB.Model(&models.TwitterSpace{}).
			Select("id").
			Where("`status` IN ? AND data_status = ?", []string{"live"}, models.DataStatusEnable).
			Find(&ids).Error; err != nil {
			log.Println(err)
		}
		for _, id := range ids {
			var users []models.TwitterUser
			users, err = global.Client.SpaceUser(id)
			for _, user := range users {
				err = dao.SaveTwitterUser(user)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
