package config

import (
	log "../logger"

	"github.com/go-ini/ini"
)

type Database struct {
	User		string
	Password	string
	Host		string
	Name		string
}

var DatabaseSetting = &Database{}

type Server struct {
	RunMode      string
	HttpAddress	 string
	HttpPort     string
}

var ServerSetting = &Server{}


var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	log.InfoLogger.Println("Setting up configuration parameters...")

	var err error
	cfg, err = ini.Load("config/conf.ini")
	if err != nil {
		log.FatalLogger.Println("config.Setup, fail to parse 'config/conf.ini': %v", err)
	}
 
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)

	// DatabaseSetting.User = cfg.Section("database").Key("User").String()
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.FatalLogger.Println("config.MapTo %s err: %v", section, err)
	}
}
