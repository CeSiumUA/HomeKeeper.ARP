package main

import "github.com/mostlygeek/arp"

func ArpGetLocalAddresses() *[]string {
	addresses := make([]string, 0)
	for ip := range arp.Table() {
		addresses = append(addresses, ip)
	}
	return &addresses
}
