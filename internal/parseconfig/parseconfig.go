package parseconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Devices map[string][]string `json:"devices"`
}

func ParseConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return &config, nil
}

func GetCommandsForDevice(config *Config, deviceName string) ([]string, bool) {
	commands, exists := config.Devices[deviceName]
	return commands, exists
}

/*
timofey.iuzyak@timofeyIuzyak:~/esr-base$ git remote -v
origin  git@gitlab2.eltex.loc:esr-routers/esr-base.git (fetch)
*/

func ParseESRbaseConfig() {

}
