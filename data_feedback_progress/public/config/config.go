package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	Debug bool

	FpMysqlDsn     string
	FpMysqlMaxConn int
)

func loadConfig() {
	Debug = viper.GetBool("debug")

	FpMysqlDsn = viper.GetString("feedback_progress_mysql_dsn")
	FpMysqlMaxConn = viper.GetInt("feedback_progress_mysql_max_conn")
}

func updateConfig() {
	ticker := time.Tick(60 * time.Second)

	select {
	case <-ticker:
		err := viper.ReadInConfig()
		if err == nil {
			loadConfig()
		}
	}
}

func configSetting(path string) {
	viper.SetConfigName(`config`) // 默认配置文件是同一目录下config.yaml
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func init() {
	var path string
	if p := os.Getenv("CONFIG_PATH"); p != "" {
		path = p
	} else {
		path = fmt.Sprintf("%s/conf", App().BasePath)
	}

	fmt.Println(path)
	configSetting(path)
	loadConfig()
	//go updateConfig()
}
