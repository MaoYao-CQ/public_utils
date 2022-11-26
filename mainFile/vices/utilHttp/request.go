package utilHttp
import (
	"net"
	"net/http"
	"strconv"
	"strings"
)

func ClientIP(r *http.Request, debug ...bool) string {

	var ips string

	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		if len(debug) > 0 && debug[0] {
			ips += "4#" + ip + ";"
		} else {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		if len(debug) > 0 && debug[0] {
			ips += "5#" + ip + ";"
		} else {
			return ip
		}
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if len(debug) > 0 && debug[0] {
			ips += "6#" + ip + ";"
		} else {
			return ip
		}
	}

	if len(debug) > 0 && debug[0] {
		return ips
	} else {
		return ""
	}
}

func ClientPublicIP(r *http.Request, debug ...bool) string {
	var ip string
	var index int
	var ips string

	for index, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !HasLocalIPAddr(ip) {
			if len(debug) > 0 && debug[0] {
				ips += "1#" + strconv.Itoa(index) + "-" + ip + ";"
			} else {
				return ip
			}
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIPAddr(ip) {
		if len(debug) > 0 && debug[0] {
			ips += "2#" + ip + ";"
		} else {
			return ip
		}
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !HasLocalIPAddr(ip) {
			if len(debug) > 0 && debug[0] {
				ips += "3#" + ip + ";"
			} else {
				return ip
			}
		}
	}

	if len(debug) > 0 && debug[0] {
		return ips
	} else {
		return ""
	}
}

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
