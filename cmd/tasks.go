package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/devoc09/gtodo/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/api/tasks/v1"
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "tasks in your TODO Lists.",
	Long: `tasks in your TODO Lists.
        currently only action show and create.
        `,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var showTasksCmd = &cobra.Command{
	Use:   "show",
	Short: "show tasks in your TODO Lists.",
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.ReadCredentials()
		client := getClient(config)
		srv, err := tasks.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve TODO clinet %v", err)
		}
		tasks, err := internal.GetTasks(srv, listId, ShowCompletedFlag)
		if err != nil {
			color.Red(err.Error())
			return
		}
		for i, t := range tasks {
			color.Green("[%d] %s\n", i+1, t.Title)
			fmt.Printf("  %s: %s\n", color.YellowString("Note"), t.Notes)
			fmt.Printf("  %s: %s\n", color.YellowString("Status"), t.Status)
			due, err := time.Parse(time.RFC3339, t.Due)
			if err != nil {
				fmt.Printf("  %s: %s\n", color.YellowString("Due"), color.BlueString("Date not set"))
			} else {
				fmt.Printf("  %s: %s\n", color.YellowString("Due"), due.Format("2006/1/2 15:04:05"))
			}
		}
	},
}

var addTasksCmd = &cobra.Command{
	Use:   "add",
	Short: "add task in your TODO List.",
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.ReadCredentials()
		client := getClient(config)
		srv, err := tasks.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve TODO clinet %v", err)
		}
		title := internal.GetInput("InputTitle:")
		note := internal.GetInput("InputNote(press enter skip):")
		due := internal.GetInput("InputDueDate(ex. 2021-04-01)(press enter skip):")
		due = due + "T00:00:00.00Z"
		task := internal.CreateTask(title, note, due)
		r, err := srv.Tasks.Insert(listId, task).Do()
		if err != nil {
			color.Red("Unable to create task: %v", err)
			return
		}
		color.Green("%v created", r.Title)
	},
}

var doneTasksCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark up task as done",
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.ReadCredentials()
		client := getClient(config)
		srv, err := tasks.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve TODO clinet %v", err)
		}
		taskls, err := internal.GetTasks(srv, listId, ShowCompletedFlag)
		if err != nil {
			color.Red(err.Error())
			return
		}
		taskmap := make(map[int]*tasks.Task)
		for i, t := range taskls {
			taskmap[i+1] = t
			color.Green("[%d] %s\n", i+1, t.Title)
			fmt.Printf("  %s: %s\n", color.YellowString("Note"), t.Notes)
			fmt.Printf("  %s: %s\n", color.YellowString("Status"), t.Status)
		}
		tasknum, err := internal.GetInputNum("Input Task Num:")
		if err != nil {
			return
		}
		task := taskmap[tasknum]
		task.Status = "completed"
		task.Hidden = true
		_, err = srv.Tasks.Patch(listId, task.Id, task).Do()
		if err != nil {
			color.Red("Unable to mark up task as completed: %v", err)
			return
		}
		color.Green("Mark up as complete: " + task.Title)
	},
}

var rmTasksCmd = &cobra.Command{
	Use:   "rm",
	Short: "delete Task",
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.ReadCredentials()
		client := getClient(config)
		srv, err := tasks.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve TODO clinet %v", err)
		}
		taskls, err := internal.GetTasks(srv, listId, ShowCompletedFlag)
		if err != nil {
			color.Red(err.Error())
			return
		}
		taskmap := make(map[int]*tasks.Task)
		for i, t := range taskls {
			taskmap[i+1] = t
			color.Green("[%d] %s\n", i+1, t.Title)
			fmt.Printf("  %s: %s\n", color.YellowString("Note"), t.Notes)
			fmt.Printf("  %s: %s\n", color.YellowString("Status"), t.Status)
		}
		tasknum, err := internal.GetInputNum("Input Task Num:")
		if err != nil {
			return
		}
		task := taskmap[tasknum]
		err = srv.Tasks.Delete(listId, task.Id).Do()
		if err != nil {
			color.Red("Unable to delete task: %v", err)
			return
		}
		fmt.Printf("%s: %s\n", color.GreenString("Deleted"), task.Title)
	},
}

var listId string = internal.LoadConfig().ListId
var ShowCompletedFlag bool = false

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.AddCommand(showTasksCmd, addTasksCmd, doneTasksCmd, rmTasksCmd)
}
