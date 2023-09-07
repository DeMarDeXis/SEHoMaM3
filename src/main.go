package main

import (
	"fmt"
	//"strings"

	//"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "myapp"}

	var name string

	var cmdSayHello = &cobra.Command{
		Use:   "sayhello",
		Short: "Say hello to someone",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	cmdSayHello.Flags().StringVarP(&name, "name", "n", "World", "Name to say hello to")

	rootCmd.AddCommand(cmdSayHello)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
