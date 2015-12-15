package main

import (
	"github.com/HackerLoop/rotonde/shared"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"

	log "github.com/Sirupsen/logrus"

	_ "github.com/kidoman/embd/host/rpi"
)

const i2cAddr = 0x40

const servoPWMFreq = 50

func main() {
	client := rotonde.NewClient("ws://rotonde:4224/")

	bus := embd.NewI2CBus(1)
	pca = &pca9685.PCA9685{pca9685.New(bus, i2cAddr)}
	pca.Freq = servoPWMFreq

	client.OnNamedAction("PCA9685_CHANNEL", func(m interface{}) bool {
		action := m.(rotonde.Action)

		channel, ok := action.Data["channel"].(int)
		if !ok {
			log.Error("channel should be a number")
			return
		}

		value, ok := action.Data["onTime"].(int)
		if !ok {
			log.Error("onTime should be a number")
			return
		}

		value, ok := action.Data["offTime"].(int)
		if !ok {
			log.Error("offTime should be a number")
			return
		}

		pca.SetPwm(channel, onTime, offTime)
	})

}
