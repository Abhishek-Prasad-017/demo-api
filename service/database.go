package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)




func GetSoftwareDetails(name string, s *MongoService) Softwares {

	fmt.Println(" I am in database")
	var results Softwares
	var err error
	if s==nil {
		return results
	}
	//fetching the session
	session := s.sess.Clone()
	if err != nil {
		log.Fatalln("Database connection could not be established", err)
		//return err
	}
	defer session.Close()
	c1 := session.DB("").C(collection1)
	// Printing the name of the software collected from the front-end
	//log.Println("Name of software is : ",name)

	if err = c1.Find(bson.M{}).All(&results); err != nil {
		fmt.Println("Failed to read results:", err)
	}

	log.Println(results)
	return results

}

func UpdateCount(name string, s *MongoService) Softwares {

	fmt.Println(" I am in database")
	var err error
	//var err1 error
	//fetching the session
	session := s.sess.Clone()
	if err != nil {
		log.Fatalln("Database connection could not be established", err)
	}

	var results Softwares
	defer session.Close()
	c1 := session.DB("").C(collection1)
	// Printing the name of the software collected from the front-end
	//log.Println("Name of software is : ",name)

	if err = c1.Find(bson.M{}).All(&results); err != nil {
		fmt.Println("Failed to read results:", err)
	}


	log.Println(results)
	return results

}

