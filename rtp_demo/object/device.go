package object

import (
	"github.com/avct/uasurfer"
	"net/http"
	"strings"
)

type Device struct {
	Ua          string        	`json:"ua"`
	Ip          string        	`json:"ip,omitempty"`
	Os          string        	`json:"os,omitempty"`
	_uaParse *uasurfer.UserAgent
}
func (d *Device) InitDevice(r *http.Request) {
	d.setUa(r)
	d.setIp(r)
	d.setOs()
}
func (d *Device) setIp(r *http.Request) {
	cfConnectingIp := r.Header.Get("Cf-Connecting-Ip")
	if len(cfConnectingIp)>0{
		d.Ip = cfConnectingIp
		return
	}
	forwarded := r.Header.Get("Forwarded")
	if len(forwarded) > 0 {
		forwarded = strings.Replace(forwarded,"for=","",1)
		d.Ip = forwarded
		return
	}
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if len(xForwardedFor) > 0 {
		xForwardedForArr := strings.Split(xForwardedFor, ",")
		if len(xForwardedForArr) > 0 {
			d.Ip = xForwardedForArr[0]
			return
		}
	}
	xRealIp := r.Header.Get("X-Real-Ip")
	if len(xRealIp) > 0 {
		d.Ip = xRealIp
	}
}
func (d *Device) setUa(r *http.Request) {
	d.Ua = r.Header.Get("User-Agent")
	d._uaParse = uasurfer.Parse(d.Ua)
}
func (d *Device) setOs() {
	d.Os = d._uaParse.OS.Name.String()
}
