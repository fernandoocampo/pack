package controller

import (
	"fmt"
	"os"

	"github.com/fernandoocampo/pack/util"
	"github.com/sirupsen/logrus"
)

var log *util.LogHandle

func init() {
	var err error
	log, err = util.NewLogger(util.Options{LogLevel: "Warn", LogFormat: "text", LogFields: logrus.Fields{"pkg": "controller", "srv": "pack"}})
	if err != nil {
		fmt.Printf("cant load logger: %v", err)
		os.Exit(1)
	}
}
