package gotutorial

import (
	"fmt"
	"net"
	"testing"
)

func TestDNSRecord(t *testing.T) {
	iprecords, _ := net.LookupIP("facebook.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	iprecords, _ = net.LookupIP("baidu.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	cname, _ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)

	cname, _ = net.LookupCNAME("m.baidu.com")
	fmt.Println(cname)
}

func TestPTR(t *testing.T) {
	ptr, _ := net.LookupAddr("6.8.8.8")
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}
}

func TestSRV(t *testing.T) {
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ncname: %s \n\n", cname)

	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

}
