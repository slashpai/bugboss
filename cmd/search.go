package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var bugID string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a bugzilla id",
	Long:  `Search a given bugzila id`,
	Run: func(cmd *cobra.Command, args []string) {
		searchBug()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().StringVarP(&bugID, "id", "n", "", "Bug ID")
}

func searchBug() {
	var bugs Bugs
	apiEndpoint := GetBugzillaEndpoint()
	url := fmt.Sprintf("%s%s%s", apiEndpoint, "/bug/", bugID)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(body), &bugs)
	formatBug(bugs)
}

func formatBug(bugs Bugs) {
	data := [][]string{}
	for i := 0; i < len(bugs.Bugs); i++ {
		bug := []string{strconv.Itoa(bugs.Bugs[i].BugID), bugs.Bugs[i].Summary, bugs.Bugs[i].Severity, bugs.Bugs[i].Priority, bugs.Bugs[i].AssignedTo, bugs.Bugs[i].CreatedBy, bugs.Bugs[i].QAContact, bugs.Bugs[i].Status, bugs.Bugs[i].LastUpdated.String()}
		data = append(data, bug)
	}
	header := []string{"BugID", "Summary", "Severity", "Priority", "Assigned To", "Created By", "QA Contact", "Status", "Last Updated"}
	RenderTable(data, header)
}
