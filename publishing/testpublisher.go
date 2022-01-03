package publishing

import (
	"fmt"
	"homekeeperarp/models"
)

type CliPublisher struct {
}

func (publicher *CliPublisher) Publish(scanResult *models.ScanResult) error {
	fmt.Println("Added: ")
	for _, addr := range scanResult.Added {
		fmt.Printf("\t%s", addr)
	}
	fmt.Println("Deleted: ")
	for _, addr := range scanResult.Deleted {
		fmt.Printf("\t%s", addr)
	}
	return nil
}

func CreateCliPublisher() Publisher {
	return &CliPublisher{}
}