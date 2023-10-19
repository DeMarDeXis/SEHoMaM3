package main

import (
    "github.com/spf13/cobra"
    "fmt"
    "" // Замените на путь к вашему пакету savedata
)

func NewCobraCommand() *cobra.Command {
    var rootCmd = &cobra.Command{Use: "SEHoMaM3"}

    rootCmd.AddCommand(&cobra.Command{
        Use:   "open",
        Short: "Open a file",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            filename := args[0]
            sf := savedata.NewSaveFile(filename)
            err := sf.Open()
            if err != nil {
                fmt.Printf("Error opening file: %v\n", err)
            } else {
                fmt.Printf("Opened file: %s\n", filename)
            }
        },
    })

    return rootCmd
}
