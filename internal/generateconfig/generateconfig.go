package generateconfig

import (
	"os"

	"github.com/Notoday671/esr-worker.git/internal/parseconfig"
)

func GenerateConfig(cfg parseconfig.Config) {
	file, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
}
