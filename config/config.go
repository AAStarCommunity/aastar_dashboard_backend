package config

import "C"
import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/rpc_server/controller/oauth"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var dscTemplate = "host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s sslmode=%s"
var Environment *model.Env
var systemConfigViper *viper.Viper

func AllConfig() map[string]any {
	return systemConfigViper.AllSettings()
}

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
	systemConfigViper = viper.New()
	systemConfigViper.SetConfigFile(configPath)
	err := systemConfigViper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	githubClientId := systemConfigViper.GetString("OAuth.Github.ClientId")
	githubClientSecret := systemConfigViper.GetString("OAuth.Github.ClientSecret")
	oauth.SetGithubOAuthAppInfo(&githubClientId, &githubClientSecret)
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
