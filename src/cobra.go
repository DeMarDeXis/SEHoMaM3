package main

import (
	"fmt"

	"github.com/DeMarDeXis/SEHoMaM3/tree/cobra/savedata"

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

	var openCmd = &cobra.Command{
		Use:   "FileOpener",
		Short: "FiOp",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileway := args[0]
			sf, err := savedata.NewSaveFile(fileway) // изменено
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
			} else {
				fmt.Printf("Opened file: %s\n", fileway)
			}
		},
	}

	return rootCmd
}
