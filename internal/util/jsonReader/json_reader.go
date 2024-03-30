package jsonReader

import (
	"encoding/json"
	"log"
	"os"
)

func ReadAndConvert(filePath string, target interface{}) {
	// 파일 열기
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Open file error: ", err)
		panic(err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&target); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		panic(err)
	}
}

func Write(filePath string, target interface{}) {
	// 파일 열기
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Create file error: ", err)
		panic(err)
	}

	defer file.Close()

	if err = json.NewEncoder(file).Encode(&target); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		panic(err)
	}
}
