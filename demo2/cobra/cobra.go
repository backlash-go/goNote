package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "myapp",
		Short: "A simple example Cobra application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("this Run in cobra")
		},

	}


    cmd.Flags().String("global","","")
    cmd.Flags().String("lo","","")
    cmd.Flags().String("lo","","")
	cmd.Flags().String("global","","")

	cmd.Execute()
	fmt.Println("first output")


}
