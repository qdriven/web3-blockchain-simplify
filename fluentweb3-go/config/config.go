package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type WorkSpace struct {
	ShHome     string `mapstructure:"sh_home"`
	WorkingDir string `mapstructure:"working_dir"`
}

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type AppConfig struct {
	Ws WorkSpace      `mapstructure:"ws"`
	Db DatabaseConfig `mapstructure:"db"`
}

func init() {
	viper.SetConfigFile("app.toml")
}
func NewConfig() *AppConfig {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	config := AppConfig{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s", err))
	}
	fmt.Println(config.Ws.WorkingDir, config.Ws.ShHome)
	watchConfigChanges(&config)
	return &config
}

func watchConfigChanges(config *AppConfig) {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		// Use viper to update the struct with the new configuration data
		if err := viper.Unmarshal(&config); err != nil {
			panic(fmt.Errorf("unable to decode into struct: %s", err))
		}

		// Print out the updated configuration data
		fmt.Printf("worksapce: %s\n", config.Ws)
	})
}
