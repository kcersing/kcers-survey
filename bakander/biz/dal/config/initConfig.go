package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"

	"kcers-survey/biz/pkg/consts"
	"kcers-survey/biz/pkg/utils"
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

	if GlobalServerConfig.Host == "" {
		address, err := utils.GetLocalIPv4Address()
		if err != nil {
			hlog.Fatalf("get localIpv4Addr failed:%s", err.Error())
		} else {
			GlobalServerConfig.Host = address
		}
	}
}
