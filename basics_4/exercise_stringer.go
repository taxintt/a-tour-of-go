package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
// fmt.Sprintf - cast all type of variables to string and print it
// caution: cast (byte -> string)
func (ip IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v",ip[0], ip[1], ip[2], ip[3])
}

func main(){
	hosts := map[string]IPAddr{
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}

	for name, ip := range hosts{
		fmt.Printf("%v: %v\n", name, ip)
	}
}