package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCobraCommand() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "SEHoMaM3"}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "hello",
		Short: "Prints 'Hello, SEHoMaM3'",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, SEHoMaM3")
		},
	})

	return rootCmd
}
