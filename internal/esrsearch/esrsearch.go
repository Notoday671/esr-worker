package esrsearch

import (
	"fmt"
	"os"
	"path/filepath"
)

// написать функцию для поиска проекта
func SearchProject() {

}

func findFolder(rootDir, targetName string) (string, error) {
	var result string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверяем, является ли текущий элемент директорией и совпадает ли имя
		if info.IsDir() && info.Name() == targetName {
			result = path
			return filepath.SkipDir // Останавливаем обход текущей директории
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if result == "" {
		return "", fmt.Errorf("папка с именем %q не найдена", targetName)
	}

	return result, nil
}
