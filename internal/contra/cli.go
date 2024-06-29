package contra

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "contra",
	Short: "Contra is a API contracts verification system client",
	Run: func(cmd *cobra.Command, args []string) {
		println("Contra is working!")
	},
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
