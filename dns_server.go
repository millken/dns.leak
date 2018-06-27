package main

import (
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func dnsServe(addr, net string) {
	if err := newDnsServer(addr, net).ListenAndServe(); err != nil {
		fmt.Printf("Failed to setup the %q server: %s\n", net, err.Error())
	}
}

func newDnsServer(addr, net string) *dns.Server {
	return &dns.Server{
		Addr:       addr,
		Net:        net,
		TsigSecret: nil,
	}
}

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
	m := &dns.Msg{}
	m.SetReply(r)

	if len(r.Question) == 0 {
		return
	}
	q := r.Question[0]
	host, _, err := net.SplitHostPort(w.RemoteAddr().String())
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return
	}
	name := q.Name[0 : len(q.Name)-16]

	lru.Put(name, host)
	fmt.Println(host, name)
	/*
		switch q.Qtype {

		case dns.TypeAAAA, dns.TypeA:
			rr := &dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    300,
				},
				A: net.ParseIP("127.0.0.1").To4(),
			}

			m.Answer = append(m.Answer, rr)
		}
	*/
	w.WriteMsg(m)
}
