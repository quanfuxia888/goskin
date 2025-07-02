package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"

	"quanfuxia/pkg/config"
)

var DB *gorm.DB

func Init() {
	dsn := config.Cfg.MySQL.DSN
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			logWriter{},
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		),
	})
	if err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		os.Exit(1)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("✅ MySQL 已连接")
}

// logWriter 用于重定向日志到 zap
type logWriter struct{}

func (logWriter) Printf(format string, args ...interface{}) {
	fmt.Printf("[GORM] "+format+"\n", args...)
}
