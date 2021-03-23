package config

import (
	"log"
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
	HttpPort     int
}

var ServerSetting = &Server{}


var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("config.Setup, fail to parse 'conf/conf.ini': %v", err)
	}
 
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)

	// DatabaseSetting.User = cfg.Section("database").Key("User").String()
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("config.MapTo %s err: %v", section, err)
	}
}
