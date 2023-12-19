package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	TableDefaultCharset = "DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci"
)

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// Configs ...
var configs *ViperConfig

func Get() *ViperConfig {
	return configs
}

func init() {
	configs = readConfig(map[string]interface{}{
		"port": 7890,
	})
}

func readConfig(defaults map[string]interface{}) *ViperConfig {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.AddConfigPath(".././conf")
	v.AddConfigPath("./conf")

	v.AutomaticEnv()
	v.SetConfigName(".env.dev")

	err := v.ReadInConfig()
	if err != nil {
		zap.S().Errorw("fail to read configs", "err", err)
		return nil
	}

	return &ViperConfig{
		Viper: v,
	}
}
