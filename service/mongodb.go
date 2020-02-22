package service

import (

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
)


type MongoService struct {
	sess *mgo.Session
}


func (ms *MongoService) Close() {
	ms.sess.Close()
}

// LoadMongoFromEnv populates the values in this object with those loaded from the environmental variables
func LoadMongo(viper *viper.Viper) (*MongoService, error) {
	log.Println("I am in mongodb.go")

	// Using variable "m" to store the database details so as to connect to a session
	m := &mgo.DialInfo{}

	m.Addrs = []string{viper.GetString("server.address") + ":" +viper.GetString("server.port") }

	m.Mechanism = "SCRAM-SHA-1"

	m.Username = viper.GetString("database.dbuser")
	m.Password = viper.GetString("database.dbpassword")
	m.Database = viper.GetString("database.dbname")

	sess1, err := mgo.DialWithInfo(m)

	if err != nil {
		return nil, err
	}

	ms := MongoService{
		sess: sess1,
	}

	log.Println("I am in mongodb.go. Connection established successfully!!!")
	return &ms, nil
}




