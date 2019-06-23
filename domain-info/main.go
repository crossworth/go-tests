package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/goware/urlx"
)

func main() {
	domain := flag.String("domain", "crossworth.com.br", "the domain to check")
	flag.Parse()

	domainUrl, err := urlx.Parse(*domain)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Host %v\n", domainUrl.Host)

	ip, err := urlx.Resolve(domainUrl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("IP %v\n", ip)

	names, err := net.LookupAddr(ip.String())
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		fmt.Printf("NAME %s\n", name)
	}

	cname, err := net.LookupCNAME(domainUrl.Host)
	if err != nil {
		panic(err)
	}

	fmt.Printf("CNAME %+v\n", cname)

	txts, err := net.LookupTXT(domainUrl.Host)
	if err != nil {
		panic(err)
	}

	for _, txt := range txts {
		fmt.Printf("TXT %+v\n", txt)
	}

	nss, err := net.LookupNS(domainUrl.Host)
	if err != nil {
		panic(err)
	}

	for _, ns := range nss {
		fmt.Printf("NS %+v\n", ns)
	}

	mxs, err := net.LookupMX(domainUrl.Host)
	if err != nil {
		panic(err)
	}

	for _, mx := range mxs {
		fmt.Printf("MX %+v\n", mx)
	}

	//_, srvs, err := net.LookupSRV("", "tcp", domainUrl.Host)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, srv := range srvs {
	//	fmt.Printf("SRV %+v\n", srv)
	//}
}
