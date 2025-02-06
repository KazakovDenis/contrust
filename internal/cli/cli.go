package contrust

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "contrust",
	Short: "Contrust is an API validation system client",
	Run: func(cmd *cobra.Command, args []string) {
		println("Contrust is working!")
	},
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
