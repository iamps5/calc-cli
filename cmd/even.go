package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "Brief description of your command",
	Long:  `Long  description of your command`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			}
		}
		fmt.Printf("The even addition of %s is %d", args, evenSum)
	},
}

func init() {
	// rootCmd.AddCommand(evenCmd)
	addCmd.AddCommand(evenCmd) // Now this evenCmd is Sub-child of Add instead of root

	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")
	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
