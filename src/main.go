package texteditor

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "texteditor"}

	var fileObj string

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new file.txt",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Create(fileObj)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()
			fmt.Println("Created file:", fileObj)
		},
	}

	var readCmd = &cobra.Command{
		Use:   "read",
		Short: "Read a text file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.Open(fileObj)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			data := make([]byte, 1024)
			_, err = file.Read(data)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			fmt.Println(string(data))
		},
	}

	// var writeCmd = &cobra.Command{
	// 	Use: "write",
	// 	Short: "Write smth on text file",
	// 	Args: Cobra.ExactArgs(1),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		file,err := os.Open
	// 	},
	// }

	rootCmd.AddCommand(createCmd, readCmd)
	rootCmd.PersistentFlags().StringVarP(&fileObj, "file", "f", "", "File path")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
