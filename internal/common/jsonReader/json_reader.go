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
		log.Println("open file error: ", err)
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("file close error: ", err)
			panic(err)
		}
	}(file)

	if err = json.NewDecoder(file).Decode(&target); err != nil {
		log.Printf("error decoding JSON: %v", err)
		panic(err)
	}
}
