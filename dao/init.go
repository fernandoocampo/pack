package dao

import (
	"fmt"
	"os"

	"github.com/fernandoocampo/pack/util"
	"github.com/sirupsen/logrus"
)

var log *util.LogHandle

func init() {
	var err error
	log, err = util.NewLogger(util.Options{LogLevel: "Info", LogFormat: "text", LogFields: logrus.Fields{"pkg": "dao", "srv": "pack"}})
	if err != nil {
		fmt.Printf("cant load logger: %v", err)
		os.Exit(1)
	}
}
