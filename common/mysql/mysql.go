package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cuixiaojun001/LinkHome/common/config"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

var (
	MasterDB = "master" // MasterDB mysql主库名称
	SlaveDB  = "slave"  // SlaveDB mysql从库名称

	PlatformMaster = "platformMaster" // PlatformMaster 控制台主库
	PlatformSlave  = "platformMaster" // PlatformSlave 控制台从库
)

var MySQLPool map[string]*gorm.DB

// Config 数据库连接配置
type Config struct {
	Dsn             string // Dsn Mysql连接DSN，比如：root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=true
	MaxIdleConns    uint   // MaxIdleConns 最大空闲连接数
	MaxOpenConns    uint   // MaxOpenConns 最大连接数
	ConnMaxLifetime uint   // ConnMaxLifetime 连接最长保持时间
	IsEnableLog     bool   // IsEnableLog 是否打印SQL日志
}

func init() {
	MySQLPool = make(map[string]*gorm.DB)
}

func SetUp() {
	mySqlConfMap := map[string]string{
		"master": "mysql.storage.master",
		"slave":  "mysql.storage.slave",
	}
	SetUpMySQL(mySqlConfMap)
}

func SetUpMySQL(nameKeyMap map[string]string) {
	for alias, key := range nameKeyMap {
		var conf, err = getMySQLConfig(key)
		if err != nil {
			logger.Errorw("Get MySQL Config error", "key", key, "err", err)
			continue
		}

		gormDB, err := Open(conf)
		if err != nil || gormDB == nil {
			logger.Errorw("Connect to MySQL error", "key", key, "config", conf, "err", err)
			continue
		}
		logger.Debugw("Connected to MySQL[" + alias + "]!")

		MySQLPool[alias] = gormDB
	}
}

func getMySQLConfig(key string) (*Config, error) {
	var conf = Config{}

	if !config.IsSet(key) {
		return nil, errors.New("config key is not existed")
	}
	if err := config.Sub(key).Unmarshal(&conf); err != nil {
		return nil, err
	}

	if conf.Dsn == "" {
		return nil, errors.New("dsn is not define")
	}
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 200
	}
	if conf.ConnMaxLifetime == 0 {
		conf.ConnMaxLifetime = 450
	}

	return &conf, nil
}

// Open 打开数据库连接
func Open(dbConf *Config) (*gorm.DB, error) {
	// init sql.DB
	sqlDB, err := sql.Open("mysql", dbConf.Dsn)
	if err != nil || sqlDB.Ping() != nil {
		return nil, err
	}

	// init connection pool
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(int(dbConf.MaxIdleConns))
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(int(dbConf.MaxOpenConns))
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(dbConf.ConnMaxLifetime) * time.Second)

	// init gorm
	cfg := &gorm.Config{}
	if dbConf.IsEnableLog {
		cfg.Logger = ormlog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容)
			ormlog.Config{ // 日志配置，打印慢 SQL 和错误
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  ormlog.Info, // 日志级别
				IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 禁用彩色打印
			},
		)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm: %w", err)
	}
	return db, nil
}

// GetGormDB 获取DB实例
func GetGormDB(name string) *gorm.DB {
	if ret, ok := MySQLPool[name]; ok {
		return ret
	}
	return nil
}

func DestroyMySQL() {
	for k, v := range MySQLPool {
		if sqlDB, err := v.DB(); err == nil {
			_ = sqlDB.Close()
		}
		delete(MySQLPool, k)
	}
}
