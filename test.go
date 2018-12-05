package main

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/Johnlovescoding/ENPM613/HOLMS/pkg/mongo"
)

func main() {
	byt := []byte(
		`{
		"course_records": {
			"jseasy": {
				"chapter1": false,
				"chapter2": false,
				"chapter3": false
			}
		}
	}`)

	//jsonData, _ := json.Marshall(str)

	userData := map[string]interface{}{}
	err := json.Unmarshal(byt, &userData)
	if err != nil {
		log.Println(err)
	}
	log.Println(userData)
	//log.Println(userData)
	log.Println(reflect.TypeOf(userData))
	log.Println(reflect.TypeOf(userData["course_records"]))

	student := mongo.Student{}

	//log.Println(reflect.TypeOf(userData["course_records"]))
	err = json.Unmarshal(byt, &student)
	if err != nil {
		log.Println(err)
	}

	log.Println(student)
}
