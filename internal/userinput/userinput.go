package userinput

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ConfigWorker struct {
	Device      string
	GitBranch   string
	Clean       bool
	PathJsonCfg string
}

func HandlerUserInput(cfg *ConfigWorker) {
	var rootCmd = &cobra.Command{
		Use:   "esr-worker",
		Short: "Краткое описание",
		Long:  `Полное описание команды`,
		Run: func(cmd *cobra.Command, args []string) { //нахуй не нужно
			fmt.Printf("Device: %s\n", cfg.Device)
			fmt.Printf("Git Branch: %s\n", cfg.GitBranch)
			fmt.Printf("Clean: %t\n", cfg.Clean)
		},
	}

	rootCmd.Flags().StringVarP(&cfg.Device, "device", "d", "", "Устройство для установки")
	rootCmd.Flags().StringVarP(&cfg.GitBranch, "branch", "b", "", "Git ветка")
	rootCmd.Flags().BoolVarP(&cfg.Clean, "clean", "c", false, "Очистить перед запуском")

	// Пометим флаг обязательным
	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("device")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
