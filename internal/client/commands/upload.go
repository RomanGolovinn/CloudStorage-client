package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"mycloud-client/internal/client/api"
)

var UploadCmd = &cobra.Command{
	Use:   "upload [local_path] [remote_path]",
	Short: "Загрузить локальный файл на облачное хранилище",
	Args:  cobra.ExactArgs(2),
	Run:   runUpload,
}

func runUpload(cmd *cobra.Command, args []string) {
	localPath := args[0]
	remotePath := args[1]

	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		fmt.Printf("Ошибка: Локальный файл не найден по пути: %s\n", localPath)
		return
	}

	cloudAPI := api.NewClient(ServerURL)

	token := "" 

	fmt.Printf("Загрузка '%s' на '%s'...\n", localPath, remotePath)
	
	fileInfo, err := cloudAPI.UploadFile(localPath, remotePath, token)

	if err != nil {
		fmt.Printf("Загрузка не удалась: %v\n", err)
		return
	}

	fmt.Printf("Файл успешно загружен.\n")
	fmt.Printf("ID: %s\n", fileInfo.ID)
	fmt.Printf("Путь: %s\n", fileInfo.Path)
	fmt.Printf("Размер: %d байт\n", fileInfo.Size)
}