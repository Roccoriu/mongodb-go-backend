package settings

import (
	"encoding/json"
	"log"
	"os"
)

type Settings struct {
	ConnectionString string `json:"connection_string"`
	DBName           string `json:"db_name"`
	ServerAddr       string `json:"server_addr"`
	TcpPort          string `json:"tcp_port"`
}

func GetSettings() Settings {
	file, err := os.Open("./settings/settings.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	settings := Settings{}

	if err := decoder.Decode(&settings); err != nil {
		log.Fatal(err)
	}

	return settings
}
