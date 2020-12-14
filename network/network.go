package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type Network struct {
	IP string `json:"ip,omitempty"`
}

var endpoints = map[string]bool{
	"https://api.ipify.org?format=text": true,
	"https://myexternalip.com/text":     true,
	"https://v4.ident.me/":              true,
}

func GetPublicIP() (string, error) {
	for endpoint, enabled := range endpoints {
		if enabled {
			resp, err := http.Get(endpoint)
			if err != nil {
				endpoints[endpoint] = false
			}

			body, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				return "", err
			}

			ip := fmt.Sprintf("%s", body)
			if net.ParseIP(ip) == nil {
				return "", fmt.Errorf("incorrect format return for IPv4")
			}

			return ip, nil
		}
	}

	return "", fmt.Errorf("unable to get IP address, ensure machine is connected to the internet")
}
