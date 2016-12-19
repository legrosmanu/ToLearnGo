package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User : just to try a document.
// The fields have to begin with a capital letter. In fact, mgo is an external package.
// Remember, only the fields which begin with a capital letter are exported.
type User struct {
	Firstname, Lastname string
}

func main() {

	url := "mongodb://localhost:27017/"

	session, err := mgo.Dial(url)
	defer session.Close()

	if err != nil {

		log.Println("Error to connect to MongoDB")
		panic(err)

	} else {

		log.Println("MongoDB is ready")

		c := session.DB("spot4musicians").C("users")

		newUser := User{Firstname: "Emmanuel", Lastname: "LEGROS"}

		info, err := c.Upsert(bson.M{"firstname": "Emmanuel"}, newUser)

		if err != nil {
			panic(err)
		} else {
			log.Println(info)
		}

		var results []User
		err = c.Find(bson.M{"firstname": "Emmanuel"}).All(&results)
		if err != nil {
			panic(err)
		}
		log.Println(results)

	}

}
