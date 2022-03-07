package ipv4

import (
	"errors"
	"fmt"
	"net"
)

type IPv4 struct {
	A byte
	B byte
	C byte
	D byte
}

func ParseIPv4FromNetIP(i net.IP) (ip IPv4, err error) {
	if i = i.To4(); i == nil {
		return ip, errors.New("invalid IPv4 address")
	}
	return IPv4{
		A: i[0],
		B: i[1],
		C: i[2],
		D: i[3],
	}, nil
}

func ParseIPv4(s string) (ip IPv4, err error) {
	i := net.ParseIP(s)
	return ParseIPv4FromNetIP(i)
}

func (ip IPv4) Int() int32 {
	return int32(ip.A)<<24 + int32(ip.B)<<16 + int32(ip.C)<<8 + int32(ip.D)
}

func (ip IPv4) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.A, ip.B, ip.C, ip.D)
}

func (ip IPv4) Next() IPv4 {
	if ip.D += 1; ip.D == 0 {
		if ip.C += 1; ip.C == 0 {
			if ip.B += 1; ip.B == 0 {
				ip.A += 1
			}
		}
	}
	return ip
}

func (ip IPv4) Prev() IPv4 {
	if ip.D -= 1; ip.D == 255 {
		if ip.C -= 1; ip.C == 255 {
			if ip.B -= 1; ip.B == 255 {
				ip.A -= 1
			}
		}
	}
	return ip
}
