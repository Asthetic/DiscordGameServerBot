package main

import (
	"fmt"

	"github.com/Asthetic/DiscordGameServerBot/network"
	"github.com/Asthetic/DiscordGameServerBot/storage"
)

func main() {
	currentIP, err := storage.GetIP()
	if err != nil {
		fmt.Printf("error fetching file: %v", err)
	}

	ip, err := network.GetPublicIP()
	if err != nil {
		fmt.Printf("error getting IP address: %v", err)
	}

	if currentIP != ip {
		currentIP = ip
		storage.WriteIP(network.Network{IP: currentIP})
	}

	fmt.Printf("%s", currentIP)
}
