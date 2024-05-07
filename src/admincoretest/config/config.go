package config

import (
	"admincoretest/xorm/models"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ziipin-server/niuhe"
	"gopkg.in/yaml.v2"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var Config struct {
	ServerAddr string
	Debug      bool
	LogLevel   string
	DB         struct {
		ShowSQL bool
		Sync    bool   // 是否同步数据库
		Debug   bool   // 是否开启 debug 模式
		Main    string // 数据库链接信息
	}
	// TODO 在这下面添加其他配置内容
}

var MainDB *xorm.Engine

func LoadConfig(path string) error {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(configBytes, &Config); err != nil {
		return err
	}
	niuhe.SetLogLevel(niuhe.MustParseLogLevelName(Config.LogLevel))

	if err := initDatabases(); err != nil {
		return err
	}
	return nil
}

func initDatabases() error {
	var err error
	// 初始化主库
	if MainDB, err = xorm.NewEngine("mysql", Config.DB.Main); err != nil {
		return err
	}
	if Config.DB.ShowSQL {
		niuhe.LogInfo("xorm: main db show sql")
		MainDB.ShowSQL(true)
	}
	if Config.DB.Debug {
		niuhe.LogInfo("xorm: main db log debug")
		MainDB.Logger().SetLevel(log.LOG_DEBUG)
	}
	if Config.DB.Sync {
		err := MainDB.Sync2(models.GetSyncModels()...)
		niuhe.LogInfo("xorm: main db sync, %v", err)
	}

	return nil
}

func init() {

}
