package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func Read(token string, Url string) map[string]interface{} {
	url := Url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	req.Header.Set("X-Vault-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		panic(err)
	}
	return data

}
