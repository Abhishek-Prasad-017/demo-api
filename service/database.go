package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)




func GetSoftwareDetails(name string, s *MongoService) Softwares {

	fmt.Println(" I am in database")
	var err error
	//fetching the session
	session := s.sess.Clone()
	if err != nil {
		log.Fatalln("Database connection could not be established", err)
	}

	var results Softwares
	defer session.Close()
	c := session.DB("").C(collection)

	// Printing the name of the software collected from the front-end
	log.Println("Name of software is : ",name)

	if err = c.Find(bson.M{"softwareName": name}).All(&results); err != nil {
		fmt.Println("Failed to read results:", err)
	}

	return results

}
