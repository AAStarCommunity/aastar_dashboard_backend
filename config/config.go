package config

import "C"
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var dscTemplate = "host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=%s"
var systemConfigViper *viper.Viper
var (
	KeyJwtSecret = "JWT.Security"
	KeyJwtRealm  = "JWT.Realm"
)

func AllConfig() map[string]any {
	return systemConfigViper.AllSettings()
}

func Init(configPath string) {

	systemConfigViper = viper.New()
	systemConfigViper.SetConfigFile(configPath)
	err := systemConfigViper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	logrus.Infof("System Config: %v", AllConfig())

}
func GetSystemConfigByKey(key string) string {
	return systemConfigViper.GetString(key)
}
func GetSystemConfigInt64yKey(key string) int64 {
	return systemConfigViper.GetInt64(key)
}
func GetDsn() string {
	return fmt.Sprintf(dscTemplate,
		systemConfigViper.GetString("DB.host"),
		systemConfigViper.GetString("DB.port"),
		systemConfigViper.GetString("DB.user"),
		systemConfigViper.GetString("DB.password"),
		systemConfigViper.GetString("DB.db_name"),
		systemConfigViper.GetString("DB.tz"),
		systemConfigViper.GetString("DB.ssl_mode"),
	)
}
