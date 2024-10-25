package bitcask

import (
	"fmt"
	"os"

	"hiteshwadhwani/bitcask-datstore.git/pkg"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bitcask",
	Short: "Bitcask Fast key/value datastore",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to bitcask",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.RunInteractiveMode()
	},
}

func Execute() {
	rootCmd.AddCommand(connectCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
