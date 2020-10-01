package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	ProjectPath = "/Users/fuhui/Code/src/github.com/think-next/lesson"
)

func main() {
	file, err := os.Open(ProjectPath + "/yaml/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	stream, _ := ioutil.ReadAll(file)

	/** parse two **/
	type yamlStruct struct {
		Title string // default yaml tag
		Class []string
		Age float64
	}
	resultStruct := &yamlStruct{}
	yaml.Unmarshal(stream, resultStruct)
	log.Printf("%v", resultStruct)
	/** parse end **/

	// config.ini -> map[string]interface{}
	result := make(map[string]interface{})
	if err := yaml.Unmarshal(stream, &result); err != nil {
		log.Fatal(err)
	}

	log.Printf("age real type %T, %v", result["age"], result["age"])
	// assert
	subjects := make([]string, 0)
	if class, isOk := result["class"]; isOk {
		realClass, isOk := class.([]interface{})
		if !isOk {
			log.Fatal("assert failure")
		}
		for _, v := range realClass {
			subjects = append(subjects, v.(string))
		}
		log.Println(subjects)
	}
}
