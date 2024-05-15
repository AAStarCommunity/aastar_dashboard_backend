package config

import (
	"aastar_dashboard_back/global_const"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	//viper.SetConfigFile("config.json")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log("Config loaded successfully")
	//t.Logf("Config: %v", viper.AllSettings())
	//t.Logf("DB Host: %v", viper.GetString("data_base.host"))
	timeObj := time.Now()

	res := timeObj.Format(global_const.TimeFormat)
	t.Logf(res)
}
