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
			fmt.Printf("Path JSON Config: %s\n", cfg.pathJsonCfg)
			fmt.Printf("Device: %s\n", cfg.device)
			fmt.Printf("Git Branch: %s\n", cfg.gitBranch)
			fmt.Printf("Clean: %t\n", cfg.clean)
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
