package info

import (
	"fmt"
	"inc/lib"
	"time"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "ping",
		As:       []string{"ping"},
		Tags:     "info",
		IsPrefix: true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			start := time.Now()
			fmt.Println("pong")
			m.Reply("Speed: " + time.Since(start).String())
		},
	})
}
