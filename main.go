package vault

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// write and update data
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

// Read data from vault KV 2
func Read(token string, Url string) map[string]interface{} {
	url := Url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		panic(err)
	}

	req.Header.Set("X-Vault-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		panic(err)
	}
	return data

}

// to delete
func Delete(token string, Url string) {
	req, err := http.NewRequest("DELETE", Url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Vault-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()

}
