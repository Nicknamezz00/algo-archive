package conf

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"sync"
	"time"
)

var (
	db    *gorm.DB
	Redis *redis.Client
	once  sync.Once
)

func setupDBEngine() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisSetting.Host,
		Password: redisSetting.Password,
		DB:       redisSetting.DB,
	})
}

func InitGormDB() *gorm.DB {
	once.Do(func() {
		var err error
		if db, err = newDBEngine(); err != nil {
			logrus.Fatalf("new gorm database failed: %s", err)
		}
	})
	return db
}

func newDBEngine() (*gorm.DB, error) {
	newLogger := logger.New(
		logrus.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  DBSetting.getLogLevel(),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	config := &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DBSetting.TablePrefix,
			SingularTable: true,
		},
	}

	plugin := dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(MysqlSetting.MaxIdleConns).
		SetMaxOpenConns(MysqlSetting.MaxOpenConns)

	var (
		db  *gorm.DB
		err error
	)
	logrus.Debugln("use MySQL as database")
	if db, err = gorm.Open(mysql.Open(MysqlSetting.Dsn()), config); err != nil {
		_ = db.Use(plugin)
	}

	return db, err
}
