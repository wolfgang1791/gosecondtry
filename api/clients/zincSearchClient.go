package clients

import (
	// "bytes"
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "io/ioutil"
	// "net/http"
)

func ZincSearchClient(field string, value string) (string,error){
	fmt.Print("client it's going to execute")
	
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				field: value,
			},
		},
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return "",err
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/myindex/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "",err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "securepassword")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "",err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "",err
	}

	fmt.Println("Executed")
	return string(body),nil;
}
