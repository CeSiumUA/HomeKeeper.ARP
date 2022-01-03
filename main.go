package main

import (
	"fmt"
	"homekeeperarp/models"
	"homekeeperarp/publishing"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	publisher, err := createHttpPublisher()
	//publisher, err := createCliPublisher()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	runScanner(publisher)
}

func runScanner(publisher *publishing.Publisher) {
	for {
		arpAddresses := PingAddresses()
		created, deleted := GetDifference(arpAddresses)
		createdDns := make([]models.DnsAddress, len(*created))
		deletedDns := make([]models.DnsAddress, len(*deleted))

		for index, address := range *created {
			createdDns[index] = *(getDns(address))
		}

		for index, address := range *deleted {
			deletedDns[index] = *(getDns(address))
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

func createHttpPublisher() (*publishing.Publisher, error) {
	arguments := os.Args
	if len(arguments) == 1 {
		return nil, fmt.Errorf("Api endpoint not specified!")
	}
	clearArguments := arguments[1:]
	publisher := publishing.CreateHttpPublisher(clearArguments[0])
	return &publisher, nil
}

func createCliPublisher() (*publishing.Publisher, error) {
	publisher := publishing.CreateCliPublisher()
	return &publisher, nil
}
