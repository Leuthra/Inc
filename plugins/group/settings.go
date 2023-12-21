package group

import (
	"fmt"
	"inc/lib"
	"strings"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(opengc|opengroup|closegc|closegroup)",
		As:       []string{"opengc", "closegc"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsAdmin:  true,
		IsGroup:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			err := client.WA.SetGroupLocked(m.From, strings.Contains(m.Command, "close"))
			if err != nil {
				m.Reply("Error")
				fmt.Println(err.Error())
			}
		},
	})
}
