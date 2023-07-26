package cmd

import (
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resumo",
	Long: `Resumo`,
}

func init() {
	rootCmd.AddCommand(resumeCmd)
}
