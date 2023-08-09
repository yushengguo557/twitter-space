package main

import (
	"context"
	"fmt"
	"github.com/yushengguo557/twitter-space/dao"
	"github.com/yushengguo557/twitter-space/global"
	"github.com/yushengguo557/twitter-space/models"
	"golang.org/x/time/rate"
	"log"
	"sync"
	"time"
)

// Corn 周期任务
func Corn() {
	go TimedLookupUser()
	go TimedUpdateSpace()
	go TimedSearchSpace()
	go TimedLookupCreator()
	//TimedUpdate()
}

// TimedUpdate 定时更新
func TimedUpdate() {
	var err error
	for range time.NewTicker(time.Minute * 15).C {
		// 1.清空数据库 Space 相关数据
		//err = global.App.DB.Model(&models.TwitterSpace{}).
		//	Where("data_status = ?", models.DataStatusEnable).
		//	Updates(map[string]any{
		//		"data_status": models.DataStatusDisable,
		//	}).Error
		//if err != nil {
		//	log.Println(err)
		//}

		// 2.搜索
		log.Println("开始搜索...")
		var spaces []models.TwitterSpace
		group := sync.WaitGroup{}
		limiter := rate.NewLimiter(rate.Every(15*time.Minute/25), 1)
		querys := []string{
			models.SpaceNFT,
			models.SpaceWEB3,
			models.SpaceMetaVerse,
			models.SpaceGame,
			models.SpaceDeFi,
			models.SpaceDAO,
			models.SpaceLayer2,
			models.SpaceInvestment,
			models.SpaceInvesting,
			models.SpaceTechnology,
		}
		for _, query := range querys {
			group.Add(1)
			go func(query string) {
				defer group.Done()
				limiter.Wait(context.Background())
				spaces, err = global.App.TwitterClient.SpaceSearch(query)
				if err != nil {
					log.Println(err)
					return
				}
				// 2.保存Space
				for _, space := range spaces {
					err = dao.SaveTwitterSpace(space)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}(query)
		}
		group.Wait()
		log.Println("搜索结束...")

		// 3.获取房间内用户
		// limiter2 := rate.NewLimiter(rate.Every(24*time.Hour/500), 1)
		log.Println("获取房间内的用户信息")
		for _, space := range spaces {
			group.Add(1)
			go func(space models.TwitterSpace) {
				defer group.Done()
				// 3.1 若Space状态为 live 获取房间内所有用户的信息
				if space.Status == "live" {
					var users []models.TwitterUser
					users, err = global.App.TwitterClient.SpaceUser(space.ID)
					if err != nil {
						log.Println(err)
						return
					}
					for _, user := range users {
						err = dao.SaveTwitterUser(user)
						if err != nil {
							log.Println(err)
							return
						}
					}
				} else {
					// 3.2 若Space状态为 其他 获取主持人信息
					var user models.TwitterUser
					resp := global.App.DB.Model(&models.TwitterUser{}).Where("id = ?", space.CreatorId).Find(&user)
					if resp.RowsAffected == 0 {
						userPtr, err := global.App.TwitterClient.RetrieveUserInfoByID(space.CreatorId)
						if err != nil {
							return
						}
						if err := dao.SaveTwitterUser(*userPtr); err != nil {
							log.Println(err)
							return
						}
					}
					if err = global.App.DB.Model(&models.TwitterUser{}).
						Where("id = ?", space.CreatorId).
						Update("space_id", space.ID).
						Error; err != nil {
						log.Println(err)
						return
					}
				}
			}(space)
		}
		group.Wait()
	}
}

// TimedUpdateSpace 定时更新 Space
func TimedUpdateSpace() {
	var err error
	var ids []string
	limiter := rate.NewLimiter(rate.Every(15*time.Minute/25), 1) // 创建限流器
	for range time.NewTicker(time.Minute * 5).C {
		log.Println("开始更新...")
		var total int64
		if err = global.App.DB.Model(&models.TwitterSpace{}).
			Where("`status` IN ? AND data_status = ?", []string{"live", "scheduled"}, models.DataStatusEnable).
			Count(&total).Error; err != nil {
			log.Println(err)
		}
		offset := 0
		limit := 64
		group := sync.WaitGroup{}
		for offset < int(total) {
			err = limiter.Wait(context.Background())
			if err != nil {
				// 请求被限流，执行相应的操作
				fmt.Println("-------请求被限流了-------")
			} else {
				// 请求未被限流，执行相应的操作
				fmt.Println("-------请求执行成功-------")
			}
			group.Add(1)
			go func(offset int) {
				defer group.Done()
				if err = global.App.DB.Model(&models.TwitterSpace{}).
					Select("id").
					Where("`status` IN ? AND data_status = ?", []string{"live", "scheduled"}, models.DataStatusEnable).
					Limit(limit).
					Offset(offset).
					Find(&ids).Error; err != nil {
					log.Println(err)
				}
				var spaces []models.TwitterSpace
				spaces, err = global.App.TwitterClient.SpaceLookup(ids)
				if err != nil {
					log.Println(err)
					return
				}
				for _, space := range spaces {
					err = dao.SaveTwitterSpace(space)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}(offset)

			offset += limit
		}
		group.Wait()
		log.Println("更新结束...")
	}
}

// TimedSearchSpace 定时搜索Space
func TimedSearchSpace() {
	var err error
	limiter := rate.NewLimiter(rate.Every(15*time.Minute/25), 1) // 创建限流器
	// 每隔 Period 搜索一次
	for range time.NewTicker(time.Minute * 30).C {
		// 1.搜索Space
		log.Println("开始搜索...")
		group := sync.WaitGroup{}
		querys := []string{
			models.SpaceNFT,
			models.SpaceWEB3,
			models.SpaceMetaVerse,
			models.SpaceGame,
			models.SpaceDeFi,
			models.SpaceDAO,
			models.SpaceLayer2,
			models.SpaceInvestment,
			models.SpaceInvesting,
			models.SpaceTechnology,
		}
		for _, query := range querys {
			err = limiter.Wait(context.Background())
			if err != nil {
				// 请求被限流，执行相应的操作
				fmt.Println("-------请求被限流了-------")
			} else {
				// 请求未被限流，执行相应的操作
				fmt.Println("-------请求执行成功-------")
			}
			group.Add(1)
			go func(query string) {
				defer group.Done()
				spaces, err := global.App.TwitterClient.SpaceSearch(query)
				if err != nil {
					log.Println(err)
					return
				}
				// 2.保存Space
				for _, space := range spaces {
					err = dao.SaveTwitterSpace(space)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}(query)
		}
		group.Wait()
		log.Println("搜索结束...")
	}
}

// TimedLookupUser 定时获取正在直播的Space中的用户信息
func TimedLookupUser() {
	var err error
	limiter := rate.NewLimiter(rate.Every(15*time.Minute/25), 1) // 创建限流器
	for range time.NewTicker(time.Minute * 10).C {
		log.Println("用户信息 - 开始...")
		group := sync.WaitGroup{}
		// 1.清空用户所属Space
		// if err = global.App.DB.Model(&models.TwitterUser{}).
		// 	Where("data_status = ?", models.DataStatusEnable).
		// 	Updates(map[string]any{"space_id": ""}).
		// 	Error; err != nil {
		// 	log.Println(err)
		// }

		// 2.更新用户所属Space
		var ids []string
		if err = global.App.DB.Model(&models.TwitterSpace{}).
			Select("id").
			Where("`status` IN ? AND data_status = ?", []string{"live"}, models.DataStatusEnable).
			Find(&ids).Error; err != nil {
			log.Println(err)
		}
		for _, id := range ids {
			err = limiter.Wait(context.Background())
			if err != nil {
				// 请求被限流，执行相应的操作
				log.Println("-------请求被限流了-------")
			} else {
				// 请求未被限流，执行相应的操作
				log.Println("-------请求执行成功-------")
			}
			group.Add(1)
			go func(id string) {
				defer group.Done()
				var users []models.TwitterUser
				users, err = global.App.TwitterClient.SpaceUser(id)
				if err != nil {
					log.Println(err)
					return
				}
				for _, user := range users {
					err = dao.SaveTwitterUser(user)
					if err != nil {
						log.Println(err)
						return
					}
				}
			}(id)
		}
		group.Wait()
		log.Println("用户信息 - 结束...")
	}
}

// TimedLookupCreator 定时更新主持人 Creator
func TimedLookupCreator() {
	var err error
	offset := 0
	limit := 64
	var total int64
	group := sync.WaitGroup{}
	var ids []string
	var user *models.TwitterUser
	limiter := rate.NewLimiter(rate.Every(24*time.Hour/500), 1) // 创建限流器 24小时500个请求
	for range time.NewTicker(time.Minute * 10).C {
		log.Println("开始更新主持人...")
		if err = global.App.DB.Model(&models.TwitterSpace{}).
			Where("`status` IN ? AND data_status = ?", []string{"ended", "canceled", "scheduled"}, models.DataStatusEnable).
			Count(&total).Error; err != nil {
			log.Println(err)
		}

		for offset < int(total) {
			group.Add(1)
			go func(offset int) {
				defer group.Done()
				if err = global.App.DB.Model(&models.TwitterSpace{}).
					Select("id").
					Where("`status` IN ? AND data_status = ?", []string{"ended", "canceled", "scheduled"}, models.DataStatusEnable).
					Limit(limit).
					Offset(offset).
					Find(&ids).Error; err != nil {
					log.Println(err)
				}
				for _, id := range ids {
					// 从数据库中 查询 id 主持人是否存在
					var users []models.TwitterUser
					resp := global.App.DB.Model(&models.TwitterUser{}).
						Where("id = ? AND data_status = ?", id, models.DataStatusEnable).
						Find(&users)
					if resp.Error != nil {
						log.Println(err)
						return
					}
					if resp.RowsAffected < 1 {
						err = limiter.Wait(context.Background())
						if err != nil {
							// 请求被限流，执行相应的操作
							fmt.Println("-------请求被限流了-------")
						} else {
							// 请求未被限流，执行相应的操作
							fmt.Println("-------请求执行成功-------")
						}
						user, err = global.App.TwitterClient.RetrieveUserInfoByID(id)
						if err != nil {
							log.Println(err)
							return
						}
						err = dao.SaveTwitterUser(*user)
						if err != nil {
							log.Println(err)
							return
						}
					}
				}
			}(offset)

			offset += limit
		}
		group.Wait()
		log.Println("更新结束主持人...")
	}
}
