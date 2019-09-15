package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	str := ""
	for index, v := range i {
		str += fmt.Sprint(v)
		if index != len(i) - 1 {
			str += "."
		}
	}
	return str
}

func main() {
	hosts := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Println("%v: %v\n", name, ip)
	}
}