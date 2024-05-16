package config

import "C"
import (
	"aastar_dashboard_back/model"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var dscTemplate = "host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=%s"
var Environment *model.Env
var SystemConfigViper *viper.Viper

func Init(configPath string) {
	envName := model.DevEnv
	if len(os.Getenv(model.EnvKey)) > 0 {
		envName = os.Getenv(model.EnvKey)
	}
	Environment = &model.Env{
		Name: envName,
		Debugger: func() bool {
			return envName != model.ProdEnv
		}(),
	}
	SystemConfigViper = viper.New()
	SystemConfigViper.SetConfigFile(configPath)
	err := SystemConfigViper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func GetDsn() string {
	return fmt.Sprintf(dscTemplate,
		SystemConfigViper.GetString("DB.host"),
		SystemConfigViper.GetString("DB.port"),
		SystemConfigViper.GetString("DB.user"),
		SystemConfigViper.GetString("DB.password"),
		SystemConfigViper.GetString("DB.db_name"),
		SystemConfigViper.GetString("DB.tz"),
		SystemConfigViper.GetString("DB.ssl_mode"),
	)
}
