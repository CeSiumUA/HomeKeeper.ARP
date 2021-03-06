package main

import (
	"flag"
	"fmt"
	"homekeeperarp/models"
	"homekeeperarp/publishing"
	"net"
	"strings"
	"time"
)

func main() {
	publisher := createPublisher()
	runScanner(publisher)
}

func runScanner(publisher *publishing.Publisher) {
	for {
		arpAddresses := PingAddresses()
		created, deleted := GetDifference(arpAddresses)
		createdDns := make([]models.DnsAddress, 0)
		deletedDns := make([]models.DnsAddress, 0)

		for _, address := range *created {
			nameAddress := getDns(address)
			if nameAddress.Name != "" {
				createdDns = append(createdDns, *nameAddress)
			}
		}

		for _, address := range *deleted {
			nameAddress := getDns(address)
			if nameAddress.Name != "" {
				deletedDns = append(deletedDns, *nameAddress)
			}
		}

		scanResult := models.ScanResult{
			Added:   createdDns,
			Deleted: deletedDns,
		}

		(*publisher).Publish(&scanResult)
		time.Sleep(10 * time.Second)
	}
}

func getDns(address string) *models.DnsAddress {
	name, err := net.LookupAddr(address)
	dns := models.DnsAddress{}
	dns.Address = address
	if err != nil {
		dns.Name = ""
	} else {
		dns.Name = strings.Join(name, ",")
	}
	return &dns
}

func createPublisher() *publishing.Publisher {
	publisherVariant := flag.String("p", "cli", "cli - to set command line as output, http for http server sending")
	if *publisherVariant == "http" {
		publisher, err := createHttpPublisher()
		if err != nil {
			fmt.Println(err)
			publisher, _ = createCliPublisher()
		}
		return publisher

	} else {
		publisher, _ := createCliPublisher()
		return publisher
	}
}

func createHttpPublisher() (*publishing.Publisher, error) {
	httpUrl := flag.String("u", "", "enter endpoint api url")
	if *httpUrl == "" {
		return nil, fmt.Errorf("api endpoint not specified")
	}
	publisher := publishing.CreateHttpPublisher(*httpUrl)
	return &publisher, nil
}

func createCliPublisher() (*publishing.Publisher, error) {
	publisher := publishing.CreateCliPublisher()
	return &publisher, nil
}
