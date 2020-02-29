package service

import (
	"flag"
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func foo() *MongoService {
	// *MyError satisfies error
	return nil;
}

func TestMongodb(t *testing.T) {
	path := flag.String("path", ".", "a string")
	file := flag.String("file", "config", "a string")
	flag.Parse()
	//Using Viper to read the config files
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName(*file)
	v.AddConfigPath(*path)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	mongo, err := LoadMongo(v)
	if err!=nil {
		//log.Println(err)
	}
	log.Println(mongo)
	mongoError := foo()
	assert.Equal(t, mongo, mongoError)
}
