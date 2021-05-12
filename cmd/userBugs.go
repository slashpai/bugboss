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

// useremail
var userEmail string

// userBugsCmd represents the userBugs command
var userBugsCmd = &cobra.Command{
	Use:   "userBugs",
	Short: "Search bugs assigned to a user",
	Long:  `Give user email as input`,
	Run: func(cmd *cobra.Command, args []string) {
		GetUserBugs(userEmail)
	},
}

func init() {
	rootCmd.AddCommand(userBugsCmd)
	userBugsCmd.PersistentFlags().StringVarP(&userEmail, "email_id", "e", "", "Email ID")
}

// GetUserBugs ...
func GetUserBugs(userEmail string) {
	var bugs Bugs
	apiEndpoint := GetBugzillaEndpoint()

	url := fmt.Sprintf("%s%s%s", apiEndpoint, "/bug?assigned_to=", userEmail)
	// fmt.Println(url)

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
	formatUserBugs(bugs)
}

func formatUserBugs(bugs Bugs) {
	data := [][]string{}
	for i := 0; i < len(bugs.Bugs); i++ {
		bug := []string{strconv.Itoa(bugs.Bugs[i].BugID), bugs.Bugs[i].Summary, bugs.Bugs[i].Severity, bugs.Bugs[i].Priority, bugs.Bugs[i].AssignedTo, bugs.Bugs[i].CreatedBy, bugs.Bugs[i].QAContact}
		data = append(data, bug)
	}
	header := []string{"BugID", "Summary", "Severity", "Priority", "Assigned To", "Created By", "QA Contact"}
	RenderTable(data, header)
}
