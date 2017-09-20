package ip

import (
	"strconv"
	"strings"
)

func IsPrivateIp(src_ip string) bool {
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
	return !IsPrivateIp(src_ip)
}
