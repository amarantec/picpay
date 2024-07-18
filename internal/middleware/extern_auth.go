package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthorizationResponse struct {
	Authorization bool `json:"authorization"`
}

func CheckExternService() (bool, error) {
	resp, err := http.Get("https://run.mocky.io/v3/4362cf11-ee75-4a9c-b4c1-dfab7e7b5c6e")
	//resp, err := http.Get("https://util.devi.tools/api/v2/authorize")
	if err != nil {
		return false, fmt.Errorf("error when querying authorizing service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("error: server answer = false")
	}

	var authResp AuthorizationResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return false, fmt.Errorf("error when decode answer from authorizing service: %w", err)
	}
	log.Printf("Answer from authorizing service: %v", authResp)
	return authResp.Authorization, nil
}
