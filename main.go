package vault

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Write(token string, url string, data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Create request with JSON data and header
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("X-Vault-Token", token)
	if err != nil {
		panic(err)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
