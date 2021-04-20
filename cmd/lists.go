package cmd

import (
	"fmt"
	"log"

	"github.com/devoc09/gtodo/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/api/tasks/v1"
)

// listsCmd represents the lists command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Show and Create tasklists for currently singed-in account.",
	Long: `
Show and Create TODO Lists for currently singed-in accont.

Show Task Lists:
  gtodo lists show
        `,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// shwoListCmd represents lists's subcommand to show takslist
var showListsCmd = &cobra.Command{
	Use:   "show",
	Short: "show TODO Lists",
	Long:  `show TODO Lists for the google account currently signed in.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.ReadCredentials()
		client := getClient(config)
		srv, err := tasks.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve TODO Client. %v", err)
		}

		r, err := srv.Tasklists.List().MaxResults(10).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve TODO lists. %v", err)
		}
		fmt.Println("TODO Lists:")
		if len(r.Items) > 0 {
			for _, i := range r.Items {
				fmt.Printf("%s (%s)\n", i.Title, i.Id)
			}
		} else {
			fmt.Println("No TODO Lists found.")
		}
	},
}

var createListCmd = &cobra.Command{
	Use:   "create",
	Short: "create TODO List",
	Long:  `create TODO List for the google account currently signed in.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.ReadCredentials()
		clinet := getClient(config)
		srv, err := tasks.New(clinet)
		if err != nil {
			log.Fatalf("Unable to retrieve TODO Client %v", err)
		}
		if title == "" {
			fmt.Println("Title is empty. set any <Title> you want.")
			return
		}
		t := &tasks.TaskList{Title: title}
		r, err := srv.Tasklists.Insert(t).Do()
		if err != nil {
			log.Fatalf("Unable to create TODO List. %v", err)
		}
		title = ""
		fmt.Println("Created TODO List!! " + color.GreenString(r.Title))
	},
}

var title string

func init() {
	createListCmd.Flags().StringVarP(&title, "title", "t", "", "title of TODO List (required)")
	rootCmd.AddCommand(listsCmd)
	listsCmd.AddCommand(showListsCmd, createListCmd)
}
