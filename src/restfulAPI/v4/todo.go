package main

import (
	"time"
)

//Todo is a struct stored name,whether completed and when complete
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

//Todos is a list type of Todo
type Todos []Todo
