package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "convert-to-my-blog",
	Short: "Resuma seu roteiro para post em blog com IA",
	Long: `Resuma seu roteiro feito no notion para post em blog com IA, traduzidas para EN and PTBR`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


