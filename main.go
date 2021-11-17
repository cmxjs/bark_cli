package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/cmxjs/bark_cli/src/bark"
	"github.com/spf13/viper"
)

var (
	host              string
	key               string
	isArchive         int // 1 indicate auto archive
	automaticallyCopy int // 1 indicate automatically copy
)

func init() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	configFile := path.Join(userConfigDir, "bark_cli.yml")
	if _, err := os.Stat(configFile); err != nil {
		// create config file
		viper.SetDefault("host", "https://api.day.app")
		viper.Set("key", "")
		viper.SetDefault("isArchive", 1)
		viper.SetDefault("automaticallyCopy", 1)
		viper.SetConfigFile(configFile)
		viper.SetConfigType("yml")
		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
		fmt.Printf("please special key in config file(%v)\n", configFile)
		os.Exit(1)
	}
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	host = viper.GetString("host")
	if host == "" {
		fmt.Printf("please special host in config file(%v)\n", configFile)
		os.Exit(1)
	}

	key = viper.GetString("key")
	if key == "" {
		fmt.Printf("please special key in config file(%v)\n", configFile)
		os.Exit(1)
	}

	isArchive = viper.GetInt("isArchive")
	automaticallyCopy = viper.GetInt("automaticallyCopy")
}

var Body, Title, Group string

func init() {
	flag.StringVar(&Body, "b", "", "body")
	flag.StringVar(&Title, "t", "", "title")
	flag.StringVar(&Group, "g", "", "group")
	flag.Parse()
}

func main() {
	b := bark.Bark{
		Host:              host,
		Key:               key,
		IsArchive:         isArchive,
		AutomaticallyCopy: automaticallyCopy,
	}
	code, err := b.Send(Body, Title, Group)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(code)
}
