package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BugId
var BugID string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a bugzilla id",
	Long:  `Search a given bugzila id`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		//searchBug()
		//GetBug()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().StringVarP(&BugID, "id", "n", "", "Bug ID")
}

func searchBug() {
	fmt.Println("BugID: ", BugID)
}
