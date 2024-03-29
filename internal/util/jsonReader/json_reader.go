package jsonReader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadAndConvert(filePath string, target interface{}) {
	// 파일 열기
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Open file error: ", err)
		return
	}

	defer file.Close()

	// JSON 파일에서 target 구조체로 디코딩
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&target)
	if err != nil {
		fmt.Printf("decode error: %v\n", err)
		return
	}
}

func Write(filePath string, target interface{}) {
	// 파일 열기
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Create file error: ", err)
		return
	}

	defer file.Close()

	// target 구조체를 JSON 파일로 인코딩
	encoder := json.NewEncoder(file)
	err = encoder.Encode(target)
	if err != nil {
		fmt.Printf("encode error: %v\n", err)
		return
	}
}
