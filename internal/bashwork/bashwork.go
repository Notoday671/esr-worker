package bashwork

import (
	"fmt"
	"os/exec"
)

func ExecuteCommands(commands []string) {
	for _, command := range commands {
		cmd := exec.Command("sh", "-c", command)
		// Выполнение команды и получение её вывода
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Ошибка выполнения команды: %v\n", err)
			return
		}
		fmt.Println("Вывод команды", command)
		fmt.Println(string(output))
	}
}
