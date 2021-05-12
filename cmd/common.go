package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Bugs ...
type Bugs struct {
	Bugs []Bug `json:"bugs"`
}

// Bug ...
type Bug struct {
	BugID        int       `json:"id"`
	Product      string    `json:"product"`
	Component    []string  `json:"component"`
	Summary      string    `json:"summary"`
	Keywords     []string  `json:"keywords"`
	Priority     string    `json:"priority"`
	Severity     string    `json:"severity"`
	Status       string    `json:"status"`
	QAContact    string    `json:"qa_contact"`
	AssignedTo   string    `json:"assigned_to"`
	CreatedBy    string    `json:"creator"`
	Blocks       []int64   `json:"blocks"`
	CloneOf      []int64   `json:"cf_clone_of"`
	CreationTime time.Time `json:"creation_time"`
	LastUpdated  time.Time `json:"last_change_time"`
}

// GetBugzillaEndpoint ...
func GetBugzillaEndpoint() string {
	return fmt.Sprintf("%s%s%s", "https://", bugzillaURL, "/rest")
}

// RenderTable ...
func RenderTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
