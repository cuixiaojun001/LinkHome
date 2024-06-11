package bootstrap

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/cuixiaojun001/LinkHome/common/cache"
	"github.com/cuixiaojun001/LinkHome/common/config"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/common/mysql"
	"github.com/robfig/cron/v3"
)

func SetUp(configFile string) error {
	// 初始化配置
	config.Init(configFile)

	// 初始化日志
	logger.SetUp()

	// 初始化DB
	mysql.SetUp()

	// 初始化缓存
	if err := cache.Init(config.GetStringMap("redis")); err != nil {
		return err
	}

	// 初始化热门房源定时任务
	HotHousesCron()

	// 初始化用户浏览记录定时任务
	UserViewCron()

	return nil
}

func HotHousesCron() {
	c := cron.New()
	c.AddFunc("@daily", func() {
		updateHotHouses()
	})
	c.Start()

	// 立即执行一次
	updateHotHouses()
}

func UserViewCron() {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		updateUserView()
	})
	c.Start()

	// 立即执行一次
	updateUserView()
}

func updateUserView() {
	client := cache.New("linkhome")
	db := mysql.GetGormDB(mysql.MasterDB)

	keys, err := client.Keys(context.Background(), "linkhome:user_views:*")
	if err != nil {
		logger.Errorw("Could not get cert", "err", err)
		return
	}

	for _, key := range keys {
		userID, err := strconv.Atoi(key[len("linkhome:user_views:"):])
		if err != nil {
			logger.Errorw("Invalid userID", "err", err)
			continue
		}
		key = "user_views:" + strconv.Itoa(userID)
		userViews, err := client.HGetAll(context.Background(), key)
		if err != nil {
			logger.Errorw("Could not get user views from Redis", "err", err)
			continue
		}

		tx := db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		for houseIDStr, countStr := range userViews {
			houseID, err := strconv.Atoi(houseIDStr)
			if err != nil {
				log.Fatalf("Invalid houseID: %v", err)
			}

			count, err := strconv.Atoi(countStr)
			if err != nil {
				log.Fatalf("Invalid count: %v", err)
			}

			tx = tx.Exec(`
                INSERT INTO user_views (user_id, house_id, view_count)
                VALUES (?, ?, ?)
                ON DUPLICATE KEY UPDATE view_count = view_count + VALUES(view_count)
            `, userID, houseID, count)
			if err != nil {
				tx.Rollback()
				logger.Errorw("Could not insert/update user_views:", "err", err)
				continue
			}
		}

		if err := tx.Commit().Error; err != nil {
			logger.Errorw("Could not commit transaction", "err", err)
			continue
		}

		client.Delete(context.Background(), key) // 删除已转移的数据
	}

}

func updateHotHouses() {
	client := cache.New("linkhome")
	db := mysql.GetGormDB(mysql.MasterDB)

	keys, err := client.Keys(context.Background(), "linkhome:hot_houses:*")
	if err != nil {
		logger.Errorw("Could not get cert", "err", err)
		return
	}
	for _, key := range keys {
		// 从key中提取城市名
		city := strings.TrimPrefix(key, "linkhome:hot_houses:")
		key = "hot_houses:" + city
		hotHouses, err := client.ZRevrRangeWithScores(context.Background(), key, 0, 9)
		if err != nil {
			logger.Errorw("ZRevrRange Could not get hot houses", "err", err)
			continue
		}

		tx := db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// 清空当前热门房源表
		if err := tx.Exec("DELETE FROM hot_houses WHERE city = ?", city).Error; err != nil {
			logger.Errorw("DELETE Could not clear hot houses", "err", err)
			continue
		}

		// 插入新的热门房源
		for _, house := range hotHouses {
			houseId, _ := strconv.Atoi(house.Member.(string))
			if err := tx.Exec("INSERT INTO hot_houses (city, house_id, score) VALUES (?, ?, ?)", city, houseId, house.Score).Error; err != nil {
				logger.Errorw("INSERT Could not insert hot house", "err", err)
				continue
			}
		}

		if err := tx.Commit().Error; err != nil {
			logger.Errorw("Could not commit transaction", "err", err)
			continue
		}
	}
}

// Destroy 项目销毁
func Destroy() {
	mysql.DestroyMySQL()
}
