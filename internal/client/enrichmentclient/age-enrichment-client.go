package enrichmentclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AgeEnrichmentClient struct {
	AgeURL string
}

func NewAgeEnrichmentClient(AgeURL string) *AgeEnrichmentClient {
	return &AgeEnrichmentClient{AgeURL: AgeURL}
}

type AgeResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int8    `json:"age"`
}

func (c *AgeEnrichmentClient) FetchAge() (int8, error) {
	resp, err := http.Get(c.AgeURL)
	if err != nil {
		return 0, fmt.Errorf("GET request failed with error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var ageResponse AgeResponse
	if err := json.Unmarshal(body, &ageResponse); err != nil {
		return 0, err
	}

	return ageResponse.Age, nil
}

// func main() {
// 	ageURL := "http://example.com/api" // базовый URL API

// 	enrichmentClient := NewEnrichmentClient(ageURL)

// 	// Пример запроса к первому URL-у
// 	data1, err := enrichmentClient.FetchDataFromAPI("ageUrl1")
// 	if err != nil {
// 		fmt.Println("Error fetching data from ageUrl1:", err)
// 	} else {
// 		fmt.Println("Data from ageUrl1:", data1)
// 	}

// 	// Пример запроса к второму URL-у
// 	data2, err := enrichmentClient.FetchDataFromAPI("ageUrl2")
// 	if err != nil {
// 		fmt.Println("Error fetching data from ageUrl2:", err)
// 	} else {
// 		fmt.Println("Data from ageUrl2:", data2)
// 	}

// 	// Пример запроса к третьему URL-у
// 	data3, err := enrichmentClient.FetchDataFromAPI("ageUrl3")
// 	if err != nil {
// 		fmt.Println("Error fetching data from ageUrl3:", err)
// 	} else {
// 		fmt.Println("Data from ageUrl3:", data3)
// 	}
// }
