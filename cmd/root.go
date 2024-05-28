package cmd

import (
	"fmt"
	"os"
	"strings"

	"fastjump/jumper"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "fj_cli",
}

var jumpCmd = &cobra.Command{
	Use: "jump",
	Run: func(cmd *cobra.Command, args []string) {
		j := jumper.NewJumper("~/.fj/db")
		matched := j.Jump(args)
		if len(matched) == 0 {
			os.Exit(1)
		}
		fmt.Println(matched)
	},
}

var hintCmd = &cobra.Command{
	Use: "hint",
	Run: func(cmd *cobra.Command, args []string) {
		j := jumper.NewJumper("~/.fj/db")
		hints := j.Hint(args)
		fmt.Println(strings.Join(hints, "\n"))
	},
}

var cleanCmd = &cobra.Command{
	Use: "clean",
	Run: func(cmd *cobra.Command, args []string) {
		j := jumper.NewJumper("~/.fj/db")
		j.Clean()
	},
}

func init() {
	rootCmd.AddCommand(jumpCmd)
	rootCmd.AddCommand(hintCmd)
	rootCmd.AddCommand(cleanCmd)
}

func Execute() {
	rootCmd.Execute()
}
