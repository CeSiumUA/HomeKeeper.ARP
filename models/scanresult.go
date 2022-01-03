package models

type ScanResult struct {
	Added   []DnsAddress
	Deleted []DnsAddress
}
