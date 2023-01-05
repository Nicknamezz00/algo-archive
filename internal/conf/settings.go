package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

type Setting struct {
	vp *viper.Viper
}

type RedisSettingS struct {
	Host     string
	Password string
	DB       int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type FileLoggerSettingS struct {
	SavePath string
	FileName string
	FileExt  string
}

type LoggerSettingS struct {
	Level string
}

type DatabaseSettingS struct {
	TablePrefix string
	LogLevel    string
}

type MySQLSettingS struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type ServerSettingS struct {
	RunMode      string
	HTTPIp       string
	HTTPPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

func (s *MySQLSettingS) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		s.UserName,
		s.Password,
		s.Host,
		s.DBName,
		s.Charset,
		s.ParseTime,
	)
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigName("config")
	//vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}

func (s *Setting) Unmarshal(objects map[string]interface{}) error {
	for k, v := range objects {
		err := s.vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseSettingS) getLogLevel() logger.LogLevel {
	switch strings.ToLower(s.LogLevel) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Error
	}
}

func (s *LoggerSettingS) getLogLevel() logrus.Level {
	switch strings.ToLower(s.Level) {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}
