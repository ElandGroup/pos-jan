package main

import (
	"bytes"
	"flag"
	"fmt"

	"os"

	"github.com/spf13/viper"
)

var (
	fruitEnv = flag.String("fruit-env", os.Getenv("FRUIT_ENV"), "fruit env")
)

func init() {
	flag.Parse()

	config := ReadConfig(*fruitEnv)

	//3.
	fmt.Println(config.Fruit.Connection)
}

func main() {

	//1. read byte
	ReadByte()
	fmt.Printf("%T,%v", viper.Get("name"), viper.Get("name"))

	//2.read config file
	//Note that the need to add a space after the colon
	//err:name:xiao
	//right:name: xiao
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Printf("%T,%v", viper.Get("name"), viper.Get("name"))

	// go run main.go config.go -fruit-env staging

}

func ReadByte() {
	config := []byte(`
Hacker: true
name: steve
hobbies:
- snowboarding
- go
clothing:
  jacket: leather
eyes : brown
`)
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))
}
