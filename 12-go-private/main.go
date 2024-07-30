package main

import (
	"fmt"

	"github.com/yodalis/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
