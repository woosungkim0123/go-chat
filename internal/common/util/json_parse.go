package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetBodyData(r *http.Request, target interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		log.Printf("failed json decode: %v", err)
		return err
	}
	return nil
}
