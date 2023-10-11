package enrichmentclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GenderEnrichmentClient struct {
	GenderURL string
}

func NewGenderEnrichmentClient(GenderURL string) *GenderEnrichmentClient {
	return &GenderEnrichmentClient{GenderURL: GenderURL}
}

type GenderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func (c *GenderEnrichmentClient) FetchGenderFromAPI(endpoint string) (string, error) {
	url := fmt.Sprintf("%s/%s", c.GenderURL, endpoint)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var genderResponse GenderResponse
	if err := json.Unmarshal(body, &genderResponse); err != nil {
		return "", err
	}
	return genderResponse.Gender, nil 
}


// func main() {
// 	GenderURL := "http://example.com/api" // Замените на базовый URL вашего API

// 	enrichmentClient := NewEnrichmentClient(GenderURL)

// 	// Пример запроса к первому URL-у
// 	data1, err := enrichmentClient.FetchDataFromAPI("endpoint1")
// 	if err != nil {
// 		fmt.Println("Error fetching data from endpoint1:", err)
// 	} else {
// 		fmt.Println("Data from endpoint1:", data1)
// 	}

// 	// Пример запроса к второму URL-у
// 	data2, err := enrichmentClient.FetchDataFromAPI("endpoint2")
// 	if err != nil {
// 		fmt.Println("Error fetching data from endpoint2:", err)
// 	} else {
// 		fmt.Println("Data from endpoint2:", data2)
// 	}

// 	// Пример запроса к третьему URL-у
// 	data3, err := enrichmentClient.FetchDataFromAPI("endpoint3")
// 	if err != nil {
// 		fmt.Println("Error fetching data from endpoint3:", err)
// 	} else {
// 		fmt.Println("Data from endpoint3:", data3)
// 	}
// }
