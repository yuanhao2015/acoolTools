/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: get_ip
* @Date: 2021-9-26 11:56
 */
package get_ip

import (
	"errors"
	"math"
	"net"
	"net/http"
	"strings"
)

type GetIpUtil struct {
}

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func (g *GetIpUtil) HasLocalIPAddr(ip string) bool {
	return g.hasLocalIP(net.ParseIP(ip))
}

// 获取公网ip
func (g *GetIpUtil) GetClientIP(r *http.Request) string {
	// var r *http.Request
	ip := strings.TrimSpace(g.clientPublicIP(r))
	if ip == "" {
		ip = g.clientIP(r)
	}
	return ip
}

// HasLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高，详见：https://github.com/thinkeridea/go-extend/issues/2
func (g *GetIpUtil) hasLocalIP(ip net.IP) bool {
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

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func (g *GetIpUtil) clientIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if ip != "" && err == nil {
		return ip
	}
	long, err := g.IPString2Long(strings.TrimSpace(r.RemoteAddr))
	if err == nil && long != 0 {
		return strings.TrimSpace(r.RemoteAddr)
	}
	return ""
}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func (g *GetIpUtil) clientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !g.HasLocalIPAddr(ip) {
			return ip
		}
	}
	if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !g.HasLocalIPAddr(ip) {
		return ip
	}
	if ip = g.RemoteIP(r); ip != "" && !g.HasLocalIPAddr(ip) {
		return ip
	}
	return ""
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func (g *GetIpUtil) RemoteIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// IPString2Long 把ip字符串转为数值
func (g *GetIpUtil) IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// Long2IPString 把数值转为ip字符串
func (g *GetIpUtil) Long2IPString(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}

// IP2Long 把net.IP转为数值
func (g *GetIpUtil) IP2Long(ip net.IP) (uint, error) {
	b := ip.To4()
	if b == nil {
		return 0, errors.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// Long2IP 把数值转为net.IP
func (g *GetIpUtil) Long2IP(i uint) (net.IP, error) {
	if i > math.MaxUint32 {
		return nil, errors.New("beyond the scope of ipv4")
	}
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip, nil
}
