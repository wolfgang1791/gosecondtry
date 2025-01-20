package dto

// Estructura que representa el formato del JSON
type Source struct {
	Timestamp   string `json:"@timestamp"`
	Body        string `json:"Body"`
	Date        string `json:"Date"`
	From        string `json:"From"`
	MessageID   string `json:"Message-ID"`
	Subject     string `json:"Subject"`
	To          string `json:"To"`
}

type Hit struct {
	Index     string `json:"_index"`
	Type      string `json:"_type"`
	ID        string `json:"_id"`
	Score     int    `json:"_score"`
	Timestamp string `json:"@timestamp"`
	Source    Source `json:"_source"`
}

type Hits struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	MaxScore int `json:"max_score"`
	Hits     []Hit `json:"hits"`
}

type Response struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits Hits `json:"hits"`
}

// Función que retorna el JSON simulado
// func getSearchResponse() (string, error) {
// 	response := Response{
// 		Took:     171,
// 		TimedOut: false,
// 		Shards: struct {
// 			Total      int `json:"total"`
// 			Successful int `json:"successful"`
// 			Skipped    int `json:"skipped"`
// 			Failed     int `json:"failed"`
// 		}{
// 			Total:      6,
// 			Successful: 6,
// 			Skipped:    0,
// 			Failed:     0,
// 		},
// 		Hits: Hits{
// 			Total: struct {
// 				Value int `json:"value"`
// 			}{Value: 2157530},
// 			MaxScore: 1,
// 			Hits: []Hit{
// 				{
// 					Index: "_index",
// 					Type:  "_doc",
// 					ID:    "2eLWGGFXxLy",
// 					Score: 1,
// 					Timestamp: "2025-01-12T23:35:24.140067328Z",
// 					Source: Source{
// 						Timestamp: "2025-01-12T23:35:24.140067328Z",
// 						Body:      "As Discussed.",
// 						Date:      "Tue, 13 Jun 2000 04:55:00 -0700 (PDT)\r",
// 						From:      "john.lavorato@enron.com\r",
// 						MessageID: "<8056068.1075845533609.JavaMail.evans@thyme>\r",
// 						Subject:   "\r",
// 						To:        "laurie.m.pare@ca.pwcglobal.com\r",
// 					},
// 				},
// 				// Puedes agregar más hits aquí
// 			},
// 		},
// 	}

// 	// Convertir la respuesta a JSON
// 	jsonData, err := json.MarshalIndent(response, "", "  ")
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(jsonData), nil
// }

// func main() {
// 	// Llamar a la función que retorna el JSON
// 	jsonResponse, err := getSearchResponse()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println(jsonResponse)
// }
