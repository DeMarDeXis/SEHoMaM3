package main

import (
	"fmt"
	"strings"

	"log"

	"github.com/spf13/cobra"
)

// Reigistr command
var (
	persistRootFlag bool

	rootCmd = &cobra.Command{
		Use:   "Root Command",
		Short: "This is Short Root Command",
		Long:  "Hello! I am Short Root Command and ...",
		/* 	Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("O.pruvet! Dobre odpoledne!")
		}, */
	}

	echoCmd = &cobra.Command{
		Use:   "Echo command",
		Short: "Just say Hello",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag") //Whats the flag?
	rootCmd.AddCommand(echoCmd)
}

// Error
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
