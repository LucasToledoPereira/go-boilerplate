package config

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	Server struct {
		Host string
		Port string
	}
	Filestore struct {
		Domain     string
		Region     string
		BucketName string
		KeyId      string
		Secret     string
	}
	Authorization struct {
		Realm       string
		Secret      string
		IdentityKey string
	}
	Settings struct {
		MultiTenant bool
		Audit       bool
	}
}

var C config

func ReadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
	return err
}
