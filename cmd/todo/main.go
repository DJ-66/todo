package main

import (
  "flag"
  "fmt"
  "os"

  "github.com/DJ-66/todo"
)

// Default file name
var todoFileName = ".todo.json"

func main() {
  // Parsing command line flags
  task := flag.String("task", "", "Task to be included in the ToDo list")
  list := flag.Bool("list", false, "List all tasks")
  complete := flag.Int("complete", 0, "Item to be completed")

  flag.Usage = func() {
    fmt.Fprintf(flag.CommandLine.Output(),
      "%s Go CLI App. Developed by Diy.Ngo - for 'DIY - GO Programer Certification' \n", os.Args[0])
    fmt.Fprintf(flag.CommandLine.Output(), "Copyleft 2024 - Because Knowledge, water and air should always be Free!\n")
    fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
    flag.PrintDefaults()
  }

  flag.Parse()


  // Define an items list
  l := &todo.List{}

  // Use the Get command to read to do items from file
  if err := l.Get(todoFileName); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  // Decide what to do based on the provided flags
  switch {
  case *list:
    // List current to do items
    fmt.Print(l)
  case *complete > 0:
    // Complete the given item
    if err := l.Complete(*complete); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }

    // Save the new list
    if err := l.Save(todoFileName); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }
  case *task != "":
    // Add the task
    l.Add(*task)

    // Save the new list
    if err := l.Save(todoFileName); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }
  default:
    // Invalid flag provided
    flag.Usage()
    os.Exit(1)
  }
}