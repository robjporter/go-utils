package network

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

func IsPrivateIP(src_ip string) bool {
	inet_network := func(ip string) uint32 {
		var (
			segments []string = strings.Split(ip, ".")
			ips      [4]uint64
			ret      uint64
		)
		for i := 0; i < 4; i++ {
			ips[i], _ = strconv.ParseUint(segments[i], 10, 64)
		}
		ret = ips[0]<<24 + ips[1]<<16 + ips[2]<<8 + ips[3]
		return uint32(ret)
	}

	ipa_beg := inet_network("10.0.0.0")
	ipa_end := inet_network("10.255.255.255")
	ipb_beg := inet_network("172.16.0.0")
	ipb_end := inet_network("172.31.255.255")
	ipc_beg := inet_network("192.168.0.0")
	ipc_end := inet_network("192.168.255.255")
	ip_seg := inet_network(src_ip)

	if (ip_seg >= ipa_beg && ip_seg <= ipa_end) || (ip_seg >= ipb_beg && ip_seg <= ipb_end) || (ip_seg >= ipc_beg && ip_seg <= ipc_end) {
		return true
	}
	return false
}

func IsPublicIP(src_ip string) bool {
	return !IsPrivateIP(src_ip)
}

func Atoi(addr string) (uint32, error) {
	var result uint32
	if len(addr) > 15 {
		return 0, errors.New("Invalid IP Address provided.")
	}
	splits := strings.Split(addr, ".")
	if len(splits) != 4 {
		return 0, errors.New("Invalid IP Address provided.")
	}
	for i := 0; i < len(splits); i++ {
		oct, err := strconv.ParseUint(splits[i], 10, 0)
		if err != nil {
			return result, errors.New("Invalid IP Address provided.")
		}
		result += uint32(oct) << uint32((4-1-i)*8)
	}
	return result, nil
}

func Itoa(addr uint32) (string, error) {
	var buf bytes.Buffer
	for i := 0; i < 4; i++ {
		oct := (addr >> uint32((4-1-i)*8)) & 0xff
		if oct < 0 {
			return "", errors.New("Invalid IP Address provided.")
		}
		buf.WriteString(strconv.FormatUint(uint64(oct), 10))
		if i < 3 {
			buf.WriteByte('.')
		}
	}
	return buf.String(), nil
}
