package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"
)



type Email struct {
	MessageID string `json:"Message-ID"`
	Date      string `json:"Date"`
	From      string `json:"From"`
	To        string `json:"To"`
	Subject   string `json:"Subject"`
	Body      string `json:"Body"`
}

func parseEmail(emailContent string) Email {
	sections := strings.SplitN(emailContent, "\n\n", 2)
	headers := sections[0]
	body := ""
	if len(sections) > 1 {
		body = sections[1]
	}

	email := Email{}
	for _, line := range strings.Split(headers, "\n") {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			switch parts[0] {
			case "Message-ID":
				email.MessageID = parts[1]
			case "Date":
				email.Date = parts[1]
			case "From":
				email.From = parts[1]
			case "To":
				email.To = parts[1]
			case "Subject":
				email.Subject = parts[1]
			}
		}
	}
	email.Body = strings.TrimSpace(body)
	return email
}

func readFIle(path string) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	emails := strings.Split(string(fileContent), "\n\n\n")
	fmt.Print(emails)
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil)) // profiling?
	// }()
	for _, emailContent := range emails {
		email := parseEmail(emailContent)

		jsonData, err := json.Marshal(email)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}
		fmt.Print(string(jsonData))
		req, err := http.NewRequest("POST", "http://localhost:4080/api/myindex/_doc", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth("admin", "securepassword")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}



func main() {
	root := "./../../maildir" // Reemplaza con el directorio que deseas recorrer

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", path, err)
			return err
		}
		readFIle(path)
		fmt.Println("Visited:", path)
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}
