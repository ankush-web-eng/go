package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	textRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error %v\n", err)
	}

	for _, txtRecord := range textRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			hasSPF = true
			spfRecord = txtRecord
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%s, %t, %t, %s, %t, %s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
