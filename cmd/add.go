package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition child-command",
	Long: `This sub-command of the root command - Calc
			It is used to add int or float`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("5 -  add.go is called to do the addition")
		// addInt(args)  // // IF no flag in this add - call only this function & close the func

		fstatus, _ := cmd.Flags().GetBool("float") // get the flag value, its default value is false
		if fstatus {                               // if status is true, call addFloat
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("float", "f", false, "This flag is to add float type numbers")
}

func addInt(args []string) {
	var sum int
	// Run for loop to iterate over the arguments;
	for _, ival := range args { // 1st return value is index of args, we can omit it using _

		itemp, err := strconv.Atoi(ival) // for string to int conversion Atio method is used, from strconv library

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func addFloat(args []string) { // This function is similar to addInt just converting string to float64
	var sum float64
	for _, fval := range args {
		// convert string to float64
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}
	fmt.Printf("Sum of floating numbers %s is %f", args, sum)
}
