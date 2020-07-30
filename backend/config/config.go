package config

import (
	"fmt"
	"strings"

	"github.com/gadavy/lhw/zap"
	"github.com/spf13/viper"

	"github.com/lissteron/cloudsuny/pkg/server"
)

func Init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs/")

	_ = viper.ReadInConfig()
}

func LoggerConfig() *zap.Config {
	return &zap.Config{
		Level:         viper.GetString("logger.level"),
		CollectorURL:  viper.GetString("logger.collector.url"),
		Namespace:     viper.GetString("logger.namespace"),
		Source:        "github.com/lissteron/cloudsuny",
	}
}

func ServerConfig() *server.Config {
	return &server.Config{
		Addr:         fmt.Sprintf("0.0.0.0:%s", viper.GetString("server.http.port")),
		ReadTimeout:  viper.GetDuration("server.read.timeout"),
		WriteTimeout: viper.GetDuration("server.write.timeout"),
		IdleTimeout:  viper.GetDuration("server.idle.timeout"),
		TLSCertFile:  viper.GetString("server.tls.cert"),
		TLSKeyFile:   viper.GetString("server.tls.key"),
	}
}
