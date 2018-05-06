package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fernandoocampo/pack/controller"
	"github.com/fernandoocampo/pack/dao"
	"github.com/fernandoocampo/pack/service"
	"github.com/fernandoocampo/pack/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// config file location
const (
	extfile = "/etc/pack/conf/conf.toml"
	intfile = "conf/conf.toml"
)

var log *util.LogHandle

func main() {
	// close first connection when server will go down.
	defer dao.CloseMgoSession()
	// start http server
	initHTTPServer()
}

func init() {
	// initialize parameters
	initConf()
	// initialize logger
	initLogger()
	// initialize database connection
	initDb()
	// initialize inversion of control
	initIoC()
}

// initConf initializes configuration file
func initConf() {
	var confFile *string
	// check if external file exists
	if _, err := os.Stat(extfile); os.IsNotExist(err) {
		confFile = flag.String("file", intfile, "service configuration file")
	} else {
		confFile = flag.String("file", extfile, "service configuration file")
	}
	flag.Parse()

	if *confFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	viper.SetConfigFile(*confFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

// initIoC initializes the dao and service used service and controller.
func initIoC() {
	mongodao := new(dao.MongoDAO)
	basicpack := new(service.BasicPack)
	healthservice := new(service.PackHealth)
	service.SetPackDAO(mongodao)
	controller.SetService(basicpack)
	controller.SetHealthService(healthservice)
}

// initLogger Initialize logger
func initLogger() {
	fmt.Println("... starting pack service logger")
	options := util.Options{
		LogLevel:  viper.GetString("service.app.logLevel"),
		LogFormat: viper.GetString("service.app.logFormat"),
		LogFields: logrus.Fields{"pkg": "main", "srv": "pack"},
	}
	fmt.Printf("%+v", options)
	var err error
	// load log for dao package
	log, err = util.NewLogger(options)
	if err != nil {
		fmt.Printf("cant load logger: %v", err)
		os.Exit(1)
	}
}

// initdb initialize mongo connection and returns session
// session must be closed when application finish
func initDb() {
	log.Info("Initializing mongo connection...")
	dao.SetDBname(viper.GetString("service.mongo.dbName"))
	dao.SetMongoAddrs(viper.GetStringSlice("service.mongo.hosts"))
	dao.SetUserdb(viper.GetString("service.mongo.userName"))
	dao.SetPwddb(viper.GetString("service.mongo.password"))
	dao.SetTimeout(viper.GetInt("service.mongo.timeout"))
	// Initialize and store a Mongo session for every requests.
	dao.InitMgoSession()
	log.Info("...Mongo session is ready")
}

// initHTTPServer start webserver on the configuration parameter host.
func initHTTPServer() {
	log.Println("Starting pack service")
	port := viper.GetString("service.app.port")
	log.Println("Starting application on ", port)
	controller.StartWebServer(port)
}
