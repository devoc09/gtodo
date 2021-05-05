package internal

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"google.golang.org/api/tasks/v1"
)

func GetTasks(srv *tasks.Service, id string, showCompleted bool) ([]*tasks.Task, error) {
	r, err := srv.Tasks.List(id).ShowHidden(showCompleted).Do()
	if err != nil {
		log.Fatalf("unable to retrieve Tasks in your default TODO List. %v", err)
	}
	if len(r.Items) == 0 {
		return nil, errors.New("No Tasks in your default TODO List.")
	}
	return r.Items, nil
}

func GetInput(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}
	s, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		s := "ERROR!"
		return s
	}
	// s = strings.TrimRight(s, "\n")
	return s
}

func GetInputNum(label string) (int, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	s, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Not Number!")
		return -1, err
	}
	return i, err
}

func CreateTask(title string, note string, due string) *tasks.Task {
	if note != "" {
		if due != "" {
			task := &tasks.Task{Title: title, Notes: note, Due: due}
			return task
		} else {
			task := &tasks.Task{Title: title, Notes: note}
			return task
		}
	} else {
		if due != "" {
			task := &tasks.Task{Title: title, Due: due}
			return task
		}
	}
	task := &tasks.Task{Title: title}
	return task
}

func IsPastDue(due string) bool {
	duedate, err := time.Parse(time.RFC3339, due)
	if err != nil {
		return false
	}
	if duedate.Before(time.Now()) {
		return true
	} else {
		return false
	}
}
