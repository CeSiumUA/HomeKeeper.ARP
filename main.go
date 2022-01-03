package main

import (
	"fmt"
	"homekeeperarp/models"
	"net"
	"strings"
)

func main() {

}

func runScanner() {
	for {
		arpAddresses := ArpGetLocalAddresses()
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

		publishToApi(&scanResult)
	}
}

func getDns(address string) *models.DnsAddress {
	name, err := net.LookupAddr(address)
	dns := models.DnsAddress{}
	dns.Address = address
	if err != nil {
		fmt.Printf("Error resolving DNS: %s", err)
		dns.Name = ""
	} else {
		dns.Name = strings.Join(name, ",")
	}
	return &dns
}

func publishToApi(scanResult *models.ScanResult) {

}
