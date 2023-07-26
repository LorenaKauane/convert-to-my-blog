package cmd

import (
	"convert-to-my-blog/internal/application"
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Criar novo resumo",
	Long: `Criar novo resumo de conteudo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n==================== Bem vindo ao meu 1º Robo em GO 💜 🤖 ====================")
		fmt.Print("\nObjetivo: Automatizar a  criação de conteúdo para meu blog 🤗 ")
		fmt.Println("\nPara executar cerifique que voce tem as chaves do Notion e da OpenAi (.env)")
		fmt.Println("\nSaiba mais em lorenaporphirio.com")
		fmt.Println("\n===============================================================================")

		application.CreateNewResume()
	},
}

func init() {
	resumeCmd.AddCommand(newCmd)
}