package commands

import (
	"github.com/spf13/cobra"
)

var ServerURL string

var RootCmd = &cobra.Command{
	Use:"CloudStorage"
	Short:"CloudStorage CLI – клиент для облачного хранилища"
}

func init(){
	RootCmd.PersistentFlags().StringVarP(&ServerURL, "server-url", "s", "http://localhost:8080", "Адрес сервера облачного хранилища (Base URL)")
	RootCmd.AddCommand(UploadCmd)
}