package publishing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"homekeeperarp/models"
	"net/http"
)

type HttpApiPublisher struct {
	ApiEndpoint string
}

func (publisher *HttpApiPublisher) Publish(scanResult *models.ScanResult) error {
	if publisher.ApiEndpoint == "" {
		return errors.New("empty API endpoint")
	}
	client := http.Client{}
	buff := bytes.NewBuffer(make([]byte, 0))
	encoder := json.NewEncoder(buff)
	encoder.Encode(*scanResult)
	request, err := http.NewRequest("POST", publisher.GetEndpoint(), buff)
	request.Header.Add("content-type", "application/json")
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	fmt.Printf("Response status from API: %s\n", response.Status)
	return err
}

func (publisher *HttpApiPublisher) SetApiEndpoint(url string) {
	publisher.ApiEndpoint = url
}

func (publisher *HttpApiPublisher) GetEndpoint() string {
	return fmt.Sprintf("%s/api/network", publisher.ApiEndpoint)
}

func CreateHttpPublisher(url string) Publisher {
	publisher := HttpApiPublisher{}
	publisher.SetApiEndpoint(url)
	return &publisher
}
