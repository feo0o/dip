package cmd

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/feo0o/dip/toolkit/ipv4"
	"github.com/urfave/cli/v2"
)

var calculateCmd *cli.Command

func init() {
	calculateCmd = &cli.Command{
		Name:    "calculate",
		Aliases: []string{"calc"},
		// Usage:   "calculate CIDR detail",
		Action: calculate,
		Subcommands: []*cli.Command{
			&cli.Command{
				Name:   "cidr",
				Usage:  "calculate CIDR detail",
				Action: calculateCIDR,
			},
		},
	}
}

func calculate(c *cli.Context) error {
	// todo
	fmt.Println("todo")
	return nil
}

func calculateCIDR(c *cli.Context) error {
	cidr := c.Args().Get(0)
	if cidr == "" {
		log.Fatal("need a cidr address")
	}
	// parse address for CIDR formate: x.x.x.x/y
	s := strings.Split(cidr, "/")
	if len(s) != 2 {
		log.Fatal("invalid CIDR formate")
	}
	mask, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal("invalid CIDR netmask")
	}
	if mask < 0 || mask > 32 {
		log.Fatal("invalid CIDR netmask")
	}
	ip := net.ParseIP(s[0])
	if ip = ip.To4(); ip == nil {
		log.Fatal("only IPv4 supported for now")
	}

	// ip's ipMask is the network
	ipMask := net.CIDRMask(mask, 32)
	if len(ipMask) != 4 {
		log.Fatal("only IPv4 supported for now")
	}
	count := 0
	if ^ipMask[0] != 0 {
		count += int(ipMask[0]) << 24
	}
	if ^ipMask[1] != 0 {
		count += int(ipMask[1]) << 16
	}
	if ^ipMask[2] != 0 {
		count += int(^ipMask[2]) << 8
	}
	if ^ipMask[3] != 0 {
		count += int(^ipMask[3])
	}

	// network ip
	nw := ip.Mask(ipMask)
	ni, err := ipv4.ParseIPv4FromNetIP(nw)
	if err != nil {
		return err
	}

	// broadcast
	bc := net.IP(make([]byte, 4))
	for i := range ip {
		bc[i] = ip[i] | ^ipMask[i]
	}

	// broadcast ip
	bi, err := ipv4.ParseIPv4FromNetIP(bc)
	if err != nil {
		return err
	}

	fmt.Printf("IP count:       %d\n", count+1)
	fmt.Printf("Network:        %s\n", nw.String())
	fmt.Printf("Netmask:        %s\n", net.IP(ipMask).String())
	fmt.Printf("First usable:   %s\n", ni.Next().String())
	fmt.Printf("Last usable:    %s\n", bi.Prev().String())
	fmt.Printf("Broadcast:      %s\n", bc.String())
	return nil
}
