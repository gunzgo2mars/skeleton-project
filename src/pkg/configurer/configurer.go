package configurer

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	CONFIG_FILENAME_FORMAT = "%s.config.%s"

	CONFIG_ENV_FILENAME_FORMAT = ".%s.env"
)

func LoadConfig(conf interface{}, configPath, configName, configType, osenv string) error {
	currentEnvironment, ok := os.LookupEnv(osenv)
	if ok {
		configName = currentEnvironment + "." + configName
	}

	fmt.Printf("Debug config name: %s \n", configName)

	viper.SetDefault("config.path", configPath)
	if err := viper.BindEnv("config.path", "CONFIG_PATH"); err != nil {
		log.Printf("warning: %s \n", err)
	}
	viper.SetConfigName(fmt.Sprintf(CONFIG_FILENAME_FORMAT, currentEnvironment, configType))
	viper.SetConfigType(configType)
	viper.AddConfigPath(viper.GetString("config.path"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config error")
	}

	if err := viper.Unmarshal(conf); err != nil {
		return errors.Wrap(err, "unmarshal config to struct error")
	}
	return nil
}
