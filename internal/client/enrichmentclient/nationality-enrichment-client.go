package enrichmentclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NationalityEnrichmentClient struct {
	NationalityURL string
}

func NewNationalityEnrichmentClient(NationalityURL string) *NationalityEnrichmentClient {
	return &NationalityEnrichmentClient{NationalityURL: NationalityURL}
}

type NationalityResponse struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

func (c *NationalityEnrichmentClient) FetchNationalityFromAPI(endpoint string) (string, error) {
	url := fmt.Sprintf("%s/%s", c.NationalityURL, endpoint)

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

	var nationalityResponse NationalityResponse
	if err := json.Unmarshal(body, &nationalityResponse); err != nil {
		return "", err
	}

	return chooseCountryByBrobability(nationalityResponse.Country), nil
}

func chooseCountryByBrobability(countries []Country)string {
	maxProbability := -1.0
	var countryWithMaxProbability string // Переменная для хранения CountryID с максимальным Probability

	for _, country := range countries {
		if country.Probability > maxProbability {
			maxProbability = country.Probability
			countryWithMaxProbability = country.CountryID
		}
	}

	return countryWithMaxProbability
}




// func main() {
// 	ageURL := "http://example.com/api" // Замените на базовый URL вашего API

// 	enrichmentClient := NewEnrichmentClient(ageURL)

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
