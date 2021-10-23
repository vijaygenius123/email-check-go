package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, Has MX, Has SPF, SPF Record, Has DMARC, DMARC Record")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err()
	err != nil{
		log.Fatalf("Error: Could Not Read From Input %v \n", err)
	}

}

func checkDomain(domain string){

	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	if len(mxRecords) > 0{
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords , err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)



}
