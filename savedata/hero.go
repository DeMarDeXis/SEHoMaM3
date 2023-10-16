package savedata

import (
	"fmt"

	"github.com/spf13/cobra"
	//"your_project_path/savedata"
)

func NewCobraCommand() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "SEHoMaM3"}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "open",
		Short: "Open a file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			sf := savedata.NewSaveFile(filename) //err
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
