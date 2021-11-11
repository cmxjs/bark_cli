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

// package main
//

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// )
//
// type Bark struct {
//
// }
//
//
// var (
// 	config_file string = "bark_config.json"
// 	env         bool
// 	base_url    string
//
// 	title             string
// 	body              string
// 	isArchive         int // 1 indicate auto archive
// 	automaticallyCopy int // 1 indicate automatically copy
// 	copy              string
// 	_url              string
// 	group             string
// )
//
// var data = []byte(`{
//   "host": "https://api.day.app",
//   "port": null,
//   "key": "JWXbq2fbwJxJQxVdB9EsLJ"
// }`)
//
// type Config struct {
// 	Host string `json:"host"`
// 	Port int    `json:"port"`
// 	Key  string `json:"key"`
// }
//
// func init() {
// 	flag.BoolVar(&env, "config", false, "print global config file path")
// 	flag.StringVar(&title, "t", "", "title")
// 	flag.StringVar(&body, "b", "", "body")
// 	flag.IntVar(&isArchive, "isArchive", 1, "auto archive")
// 	flag.IntVar(&automaticallyCopy, "automaticallyCopy", 0, "automaticallyCopy")
// 	flag.StringVar(&copy, "copy", "", "copy string, default is \"\"")
// 	flag.StringVar(&_url, "url", "", "url address, default is \"\"")
// 	flag.StringVar(&group, "group", "", "message group name \"\"")
// }
//
// func send() (status int) {
// 	if title != "" {
// 		base_url = fmt.Sprintf("%s/%s", strings.TrimSuffix(base_url, "/"), url.QueryEscape(title))
// 	}
// 	if body != "" {
// 		base_url = fmt.Sprintf("%s/%s", strings.TrimSuffix(base_url, "/"), url.QueryEscape(body))
// 	}
//
// 	p := url.Values{}
// 	p.Add("isArchive", strconv.Itoa(isArchive))
// 	p.Add("automaticallyCopy", strconv.Itoa((automaticallyCopy)))
// 	if copy != "" {
// 		p.Add("copy", copy)
// 	}
// 	if _url != "" {
// 		p.Add("url", _url)
// 	}
// 	if group != "" {
// 		p.Add("group", group)
// 	}
//
// 	base_url = fmt.Sprintf("%s?%s", base_url, p.Encode())
// 	fmt.Println(base_url)
//
// 	resp, err := http.Get(base_url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return resp.StatusCode
// }
//
// func create_config_file() (err error) {
// 	f, err := os.Create(config_file)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
//
// 	if _, err = f.Write(data); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func parse_json(file string, c *Config) (err error) {
// 	data, err := os.ReadFile(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return json.Unmarshal(data, c)
// }
//
// func main() {
// 	flag.Parse()
//
// 	ConfigDir, err := os.UserConfigDir()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	config_file = filepath.Join(ConfigDir, config_file)
// 	if _, err := os.Stat(config_file); err != nil {
// 		if err := create_config_file(); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("create default config file success, path:%s\n", config_file)
// 	}
//
// 	if env {
// 		fmt.Printf("Global_Config:\n  %s\n\n", config_file)
// 		os.Exit(0)
// 	}
//
// 	config := Config{}
// 	if err = parse_json(config_file, &config); err != nil {
// 		log.Fatal(err)
// 	}
// 	if config.Port != 0 {
// 		base_url = fmt.Sprintf("%s:%d/%s", strings.Trim(config.Host, "/"), config.Port, strings.Trim(config.Key, "/"))
// 	} else {
// 		base_url = fmt.Sprintf("%s/%s", strings.Trim(config.Host, "/"), strings.Trim(config.Key, "/"))
// 	}
// 	if body == "" {
// 		flag.Usage()
// 		log.Fatal("must input body, see help info")
// 	}
// 	status := send()
// 	fmt.Println(status)
// }
//
