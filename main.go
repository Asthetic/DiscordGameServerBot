package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/Asthetic/DiscordGameServerBot/network"
	"github.com/Asthetic/DiscordGameServerBot/storage"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Session is our global discord session
var Session, _ = discordgo.New()

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	done := make(chan bool)

	for {
		getIP()

		select {
		case <-done:
			return
		case s := <-sig:
			log.Infof("Got signal to kill program: %v", s)
			return
		case <-ticker.C:
			getIP()
		}
	}
}

func getIP() {
	currentIP, err := storage.GetIP()
	if err != nil {
		log.WithError(err).Errorf("Error fetching file: %v", err)
	}

	ip, err := network.GetPublicIP()
	if err != nil {
		log.WithError(err).Errorf("Error getting IP address: %v")
	}

	if currentIP != ip {
		currentIP = ip
		err = storage.WriteIP(network.Network{IP: currentIP})
		if err != nil {
			log.WithError(err).Errorf("Error writing IP to local storage")
		}
	}

	log.Infof("%s", currentIP)
}
