package main

import (
	"fmt"
	"net"
	"time"

	"github.com/mostlygeek/arp"
	"github.com/tatsushid/go-fastping"
)

func ArpGetLocalAddresses() *[]string {
	addresses := make([]string, 0)
	for ip := range arp.Table() {
		addresses = append(addresses, ip)
	}
	return &addresses
}

func PingAddresses() *[]string {
	baseAddress := "192.168.0."
	baseNumber := 2
	pinger := fastping.NewPinger()
	defer pinger.Stop()
	for {
		if baseNumber > 255 {
			break
		}
		pingAddress := fmt.Sprintf("%s%d", baseAddress, baseNumber)
		remoteAddress, err := net.ResolveIPAddr("ip4:icmp", pingAddress)
		if err != nil {
			continue
		}
		pinger.AddIPAddr(remoteAddress)
		baseNumber++
	}
	addresses := make([]string, 0)
	pinger.OnRecv = func(i *net.IPAddr, d time.Duration) {
		addresses = append(addresses, i.IP.String())
	}
	err := pinger.Run()
	if err != nil {
		fmt.Println(err)
	}
	return &addresses
}
