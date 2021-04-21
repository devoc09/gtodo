# gtodo
A command line tool for using Google Tasks basic actions.

## Installation
- Please get it from the [Releases](https://github.com/devoc09/gtodo/releases)

if you have a Go env
- Use Makefile or self build

Author's env
`go version go1.15.3 darwin/amd64`

## Set up
`credentials.json`, `token.json` and `config.yaml` are required<br>
1. make and download `credentials.json` from [Google Tasks API Go quickstart](https://developers.google.com/tasks/quickstart/go)(Details on how to do exactly that can be found there.)
1. make dir `~/.config/gtodo` and move `credentials.json` to `~/.config/gtodo/`
1. `$ gtodo login` or other command download `token.json` automatically to `~/.config/gtodo`
1. make `config.yaml` into `~/.config/gtodo/`<br>
```yaml
# config.yaml
listid: foobarbuz # <- write your default TODO Task List ID.
```
you can check the TODO List ID with the `$ gtodo lists show`

## Usage
#### Show TODO Lists
```
$ gtodo lists show
TODO Lists:
ListsName (your list id)
```
#### Create TODO List
```
$ gtodo lists create -t new_todo_list
Created TODO List!! new_todo_list

$ gtodo lists show
TODO Lists:
Main (your list id1)
new_todo_list (your list id2)
```
#### Show Tasks (you can only see one list of tasks)
```
$ gtodo tasks show
[1] task1
  Note:
  Status: needsAction
  Due: Date not set
[2] task2
  Note: sample description
  Status: needsAction
  Due: Date not set
[3] task3
  Note:
  Status: needsAction
  Due: 2021/4/20 00:00:00
```
#### Create Task
```
$ gtodo tasks add
InputTitle:: task4
InputNote(press enter skip)::
InputDueDate(ex. 2021-04-01)(press enter skip)::
task4 created
$ gtodo tasks show
[1] task4
  Note:
  Status: needsAction
  Due: Date not set
[2] task1
  Note:
  Status: needsAction
  Due: Date not set
[3] task2
  Note: sample description
  Status: needsAction
  Due: Date not set
[4] task3
  Note:
  Status: needsAction
  Due: 2021/4/20 00:00:00
```
#### Delete Task
```
$ gtodo tasks rm
[1] task4
  Note:
  Status: needsAction
  Due: Date not set
[2] task1
  Note:
  Status: needsAction
  Due: Date not set
[3] task2
  Note: sample description
  Status: needsAction
  Due: Date not set
[4] task3
  Note:
  Status: needsAction
  Due: 2021/4/20 00:00:00
✔ Input Task Num:: 1
Deleted: task4
$ gtodo tasks show
[1] task1
  Note:
  Status: needsAction
  Due: Date not set
[2] task2
  Note: sample description
  Status: needsAction
  Due: Date not set
[3] task3
  Note:
  Status: needsAction
  Due: 2021/4/20 00:00:00
```
#### Mark up as Done
```
$ gtodo tasks done
[1] task4
  Note:
  Status: needsAction
[2] task1
  Note:
  Status: needsAction
  Due: Date not set
[3] task2
  Note: sample description
  Status: needsAction
  Due: Date not set
[4] task3
  Note:
  Status: needsAction
  Due: 2021/4/20 00:00:00
Input Task Num:: 1
Mark up as complete: task4
$ gtodo tasks show
[1] task1
  Note:
  Status: needsAction
  Due: Date not set
[2] task2
  Note: sample description
  Status: needsAction
  Due: Date not set
[3] task3
  Note:
  Status: needsAction
  Due: 2021/4/20 00:00:00
```
#### Get token.json automatically
`$ gtodo login`

## Thanks
- Inspired by [BRO3886/gtasks](https://github.com/BRO3886/gtasks) and [mattn/todo](https://github.com/mattn/todo)
- I was able to create this app thanks to a great set of libraries.
