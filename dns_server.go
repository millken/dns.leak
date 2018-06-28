package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/beevik/guid"
	"github.com/miekg/dns"
)

func removeDuplicates(elements []string) []string {
	result := []string{}

	for i := 0; i < len(elements); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if elements[v] == elements[i] {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, elements[i])
		}
	}
	return result
}

func dnsServe(addr, net string) {
	log.Printf("Starting up DNS Server, LISTEN : %s(%s)", addr, net)
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

	if len(name) <= 36 {
		w.WriteMsg(m)
		return
	}
	uuid := name[len(name)-36:]
	if !guid.IsGuid(uuid) {
		log.Printf("[ERROR] Token [%s] is not guid v4 format", uuid)
		w.WriteMsg(m)
		return
	}
	obj := lru.Get(uuid)
	var hosts []string
	if obj == nil {
		hosts = append(hosts, host)
	} else {
		hosts = append(obj.([]string), host)
	}
	lru.Put(uuid, removeDuplicates(hosts))
	csv.Write([]byte(fmt.Sprintf("%s,%s\n", time.Now().Format("2006-01-02"), host)))
	//log.Printf("%s %s", host, q.Name)
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
