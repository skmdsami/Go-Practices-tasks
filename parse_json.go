package main

import (
	"errors"
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type User struct {
	ID int
	Name string
	Email string
	Age uint
}


func ParseAndValidateUser(json string) (*User, error) {
	// firstly remove the curly brackets
	json = strings.Replace(json, "{", "", 1)
	json = strings.Replace(json, "}", "", 1)

	properties := strings.Split(json, ",")
	hashmap := make(map[string]bool)
	hashmap["name"] = false
	hashmap["age"] = false
	hashmap["id"] = false
	hashmap["email"] = false
	user := User{}
	for _,property := range properties {
		if strings.Contains(property, ":") {
			index := strings.Index(property, ":")
			value := property[index+1: ]
			value = strings.TrimSpace(value)
			fmt.Println("value was", value)
			property = property[0: index]
			property = strings.TrimSpace(property)
			property = strings.Replace(property, `"`, ``, 2)
			fmt.Println("property was",property)
			if property == "name" {
				hashmap["name"] = true
				if len(value) > 2 {
					user.Name = value
				}else{
					return &User{}, errors.New("no name was given")
				}
			}else if property == "email"{
				hashmap["email"] = true
				if len(value) > 2 {
					user.Email = value
				}else{
					return &User{}, errors.New("no email was given")
				}
			}else if property == "age" {
				// we need to convert string to int
				age,_ := strconv.Atoi(value)
				hashmap["age"] = true
				if age > 18 {
					user.Age = uint(age)
				}else {
					return &User{}, errors.New("Below the age of 18")
				}
			}else if property == "id"{
				hashmap["id"] = true
				user.ID, _ = strconv.Atoi(value)
			}else{
				return &User{}, errors.New("Error Occured due to invalid field-name")
			}

		}
	}

	for value := range maps.Values(hashmap) {
		if value != true {
			return &User{}, errors.New("Error Occured")
		}
	}

	return &user, nil
}

func main() {
	json1 := `{"id": 1, "name": "Alice", "email": "alice@example.com", "age": 30}`
	fmt.Println("The passed User json was :",json1)
	user, error1 := ParseAndValidateUser(json1)
	fmt.Println("The Output was \n", )
	if error1 == nil {
		fmt.Println(user)
	}else{
		fmt.Println(error1)
	}
}

/*
validJSON := `{"id": 1, "name": "Alice", "email": "alice@example.com", "age": 30}`
invalidJSON := `{"id": 2, "name": "", "email": "bob@example.com", "age": 25}` // Invalid name
invalidAgeJSON := `{"id": 3, "name": "Charlie", "email": "charlie@example.com", "age": 17}` // Invalid age
malformedJSON := `{"id": 4, name: "David"}` // Malformed
*/