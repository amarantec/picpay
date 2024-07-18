package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthorizationResponse struct {
	Authorized bool `json:"authorized"`
}

func CheckExternService() (bool, error) {
	resp, err := http.Get("https://util.devi.tools/api/v2/authorize")
	if err != nil {
		return false, fmt.Errorf("error when querying authorizing service: %w", err)
	}
	defer resp.Body.Close()

	var authResp AuthorizationResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return false, fmt.Errorf("error when decode answer from authorizing service: %w", err)
	}
	return authResp.Authorized, nil
}
