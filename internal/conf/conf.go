package conf

import (
	"log"
	"sync"
	"time"
)

var (
	fileLoggerSetting *FileLoggerSettingS
	loggerSetting     *LoggerSettingS
	DBSetting         *DatabaseSettingS
	MysqlSetting      *MySQLSettingS
	ServerSetting     *ServerSettingS
	JWTSetting        *JWTSettingS
	Mutex             *sync.Mutex
)

func Initialize() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting error: %v", err)
	}

	setupLogger()
	//setupDBEngine()
}

func setupSetting() error {
	s, err := NewSetting()
	if err != nil {
		return err
	}

	mp := map[string]interface{}{
		"Database":   &DBSetting,
		"MySQL":      &MysqlSetting,
		"Logger":     &loggerSetting,
		"FileLogger": &fileLoggerSetting,
		"Server":     &ServerSetting,
		"JWT":        &JWTSetting,
	}
	if err = s.Unmarshal(mp); err != nil {
		return err
	}

	ServerSetting.ReadTimeOut *= time.Second
	ServerSetting.WriteTimeOut *= time.Second

	Mutex = &sync.Mutex{}
	return nil
}
