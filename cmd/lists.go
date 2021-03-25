package cmd

import (
	"fmt"
	"log"

	"github.com/devoc09/gtodo/internal"
	"github.com/spf13/cobra"
	"google.golang.org/api/tasks/v1"
)

// listsCmd represents the lists command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lists called")
	},
}

// listsShowCmd represents lists's subcommand to show takslist
var listsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show TODO Lists",
	Long: `show TODO Lists for the google account currently signed in.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show subcommand called")
                config := internal.ReadCredentials()
                client := getClient(config)
                srv, err := tasks.New(client)
                if err != nil {
                    log.Fatalf("Unable to retrieve TODO lists. %v", err)
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

func init() {
	rootCmd.AddCommand(listsCmd)
        listsCmd.AddCommand(listsShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
