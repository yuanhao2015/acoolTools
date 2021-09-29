package test

import (
	"github.com/yuanhao2015/acoolTools"
	"net/http"
	"testing"
)

func TestGetIp(t *testing.T) {
	for _, v := range []struct {
		remoteAddr string
		expected   string
	}{
		{"101.1.0.4:100", "101.1.0.4"},
		{"101.1.0.4:", "101.1.0.4"},
		{"192.168.1.1", "192.168.1.1"},
		{":100", ""},
	} {
		if actual := acoolTools.ClientIPUtils.GetClientIP(&http.Request{RemoteAddr: v.remoteAddr}); actual != v.expected {
			t.Errorf("RemoteAddr:%s actual: %s, expected %s", v.remoteAddr, actual, v.expected)
		}
	}

	r := &http.Request{Header: http.Header{}}
	r.Header.Set("X-Real-Ip", " 60.60.60.60 ")
	r.Header.Set("X-Forwarded-For", "  20.20.20.20, 30.30.30.30")
	r.RemoteAddr = "  40.40.40.40:42123 "
	if ip := acoolTools.ClientIPUtils.GetClientIP(r); ip != "20.20.20.20" {
		t.Errorf("actual: 20.20.20.20, expected:%s", ip)
	}
	r.Header.Del("X-Forwarded-For")
	r.Header.Set("X-Real-Ip", " 50.50.50.50 ")
	if ip := acoolTools.ClientIPUtils.GetClientIP(r); ip != "50.50.50.50" {
		t.Errorf("actual: 50.50.50.50, expected:%s", ip)
	}
	r.Header.Del("X-Forwarded-For")
	r.Header.Del("X-Real-Ip")
	if ip := acoolTools.ClientIPUtils.GetClientIP(r); ip != "40.40.40.40" {
		t.Errorf("actual: 40.40.40.40, expected:%s", ip)
	}
	r.Header.Set("X-Real-Ip", " 192.168.1.1 ")
	if ip := acoolTools.ClientIPUtils.GetClientIP(r); ip != "40.40.40.40" {
		t.Errorf("actual: 40.40.40.40, expected:%s", ip)
	}

	t.Log(acoolTools.ClientIPUtils.GetCityByIp("175.0.118.197"))
}
