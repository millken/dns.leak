package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/utahta/go-cronowriter"

	"github.com/miekg/dns"
)

var lru *LRUCache
var csv *cronowriter.CronoWriter

func main() {

	csv = cronowriter.MustNew("dns.leak-%Y%m.csv", cronowriter.WithInit())
	log.SetFlags(log.LstdFlags)

	lru = NewLRUCache()
	name := []byte{108, 101, 97, 107, 46, 100, 110, 115, 111, 97, 46, 99, 111, 109}
	fqdn := dns.Fqdn(string(name[:]))

	dns.HandleFunc(fqdn, handleDNS)

	//go dnsServe(":53", "tcp")
	go dnsServe(":53", "udp")
	go httpServer(":59105")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	fmt.Printf("Signal (%v) received, stopping\n", s)
}
