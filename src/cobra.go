package main

import (
	"fmt"
	"savedata" //Ошибка

	"github.com/spf13/cobra"
)

func NewCobraCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Homam",
		Short: "",
		//Long:  "HomamLaunch",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var openCmd = &cobra.Command{ //Ошибка
		Use:   "FileOpener",
		Short: "FiOp",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileway := args[0]
			sf := savedata.NewSaveFile(fileway)
			err := sf.Open()
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
			} else {
				fmt.Printf("Opened file: %s\n", fileway)
			}
		},
	}

	return rootCmd
}
