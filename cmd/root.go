package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{ // "rootCmd" struct var -- declared of type *cobra.Command
	Use:   "calc",
	Short: "Use this app for addition", //"A brief description of your application",
	Long: `This app can add:
			Int
			Float
			int or float even number only`,

	Run: func(cmd *cobra.Command, args []string) { fmt.Println("4 - Running Base/root command only") }, // Base cmd will print the "4 - ..", Instead of above Info (use, short & Long)
	// RUN is the function which would be called in case it's comment. Descrption will get print
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("1 - This is init func in root.go before initConfig func \n")

	cobra.OnInitialize(initConfig)

	fmt.Println("2 - This is initconfig Function. This run Before rootCmd \n ")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".calc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".calc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
