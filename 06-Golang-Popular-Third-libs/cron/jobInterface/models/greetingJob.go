package models

import "fmt"

type GreetingJob struct {
	Name  string
	JodID string
}

func (g GreetingJob) Run() {
	fmt.Println("Running Job:", g.JodID, "->", g.Name)
}
