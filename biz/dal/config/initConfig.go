package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"

	"kcers-survey/biz/pkg/consts"
)

// InitConfig 初始化配置
func InitConfig() {

	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)

	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed:%s", err.Error())
	}
	if err := v.Unmarshal(&GlobalServerConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	hlog.Info("config Info: %v", GlobalServerConfig)
}
