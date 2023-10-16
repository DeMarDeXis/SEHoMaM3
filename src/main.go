package main

/*  import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)  */

func main() {
	cmd := NewCobraCommand()
	cmd.Execute()
}
