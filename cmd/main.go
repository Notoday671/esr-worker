package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Notoday671/esr-worker.git/internal/bashwork"
	"github.com/spf13/cobra"
)

type ConfigWorker struct {
	device      string
	gitBranch   string
	clean       bool
	pathJsonCfg string
}

type Config struct {
	Devices map[string][]string `json:"devices"`
}

func handlerUserInput(cfg *ConfigWorker) {
	var rootCmd = &cobra.Command{
		Use:   "esr-worker",
		Short: "Краткое описание",
		Long:  `Полное описание команды`,
		Run: func(cmd *cobra.Command, args []string) { //нахуй не нужно
			fmt.Printf("Path JSON Config: %s\n", cfg.pathJsonCfg)
			fmt.Printf("Device: %s\n", cfg.device)
			fmt.Printf("Git Branch: %s\n", cfg.gitBranch)
			fmt.Printf("Clean: %t\n", cfg.clean)

			// Парсим конфиг
			config, err := parseConfig(cfg.pathJsonCfg)
			if err != nil {
				fmt.Printf("Ошибка парсинга конфига: %v\n", err)
				return
			}

			// Получаем команды для устройства
			commands, exists := getCommandsForDevice(config, cfg.device)
			if !exists {
				fmt.Printf("Устройство '%s' не найдено в конфиге.\n", cfg.device)
				return
			}

			// испольняем команды
			bashwork.ExecuteCommands(commands)
		},
	}

	rootCmd.Flags().StringVarP(&cfg.pathJsonCfg, "path", "p", ".", "Путь к JSON конфигу")
	rootCmd.Flags().StringVarP(&cfg.device, "device", "d", "", "Устройство для установки")
	rootCmd.Flags().StringVarP(&cfg.gitBranch, "branch", "b", "", "Git ветка")
	rootCmd.Flags().BoolVarP(&cfg.clean, "clean", "c", false, "Очистить перед запуском")

	// Пометим флаг обязательным
	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("device")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func parseConfig(path string) (*Config, error) {
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

func getCommandsForDevice(config *Config, deviceName string) ([]string, bool) {
	commands, exists := config.Devices[deviceName]
	return commands, exists
}

func main() {
	var cfg ConfigWorker
	handlerUserInput(&cfg)
}
