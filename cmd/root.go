package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var bugzillaURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bugboss",
	Short: "A bugzilla cli written in go",
	Long: `Buzilla cli to help to interact with bugzilla.
You can quickly search a bugzilla issue instead of waiting for web UI to load`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bugboss.yaml)")
	rootCmd.PersistentFlags().StringVar(&bugzillaURL, "bugzilla-url", "", "bugzilla Url")

}

// initConfig reads in config file
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".bugboss" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bugboss")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// If not passed through option take default from config file
	if bugzillaURL == "" {
		if viper.GetString("bugzilla.url") == "" {
			panic("Bugzilla URL neither provided through options nor in config file, check help option and try again!")
		} else {
			bugzillaURL = viper.GetString("bugzilla.url")
		}
	}
}
