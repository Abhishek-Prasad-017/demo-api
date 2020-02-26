package main
//go mod init github.com/demo-api
import (
	"flag"
	"fmt"
	//"github.com/demo-api/config"
	"github.com/demo-api/server"
	"github.com/spf13/viper"
	"log"
)

func main() {

	//Read secrets from config.yaml file


	//environment := flag.String("environment", "dev", "a string")
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

	// Commented code below is used to print the values from the config file

	/*
	var configuration config.Configurations

	err := v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)
	ex := v.Get("server.port")
	fmt.Println("Server is\t", ex)
	*/

	//Server address exposed
	serverPath := "0.0.0.0:8080"
	log.Println("Starting server at " + serverPath)
	log.Fatal(server.StartServer(serverPath, v))
	log.Println("Stopping server")


}
