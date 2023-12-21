package group

import (
	"fmt"
	"inc/lib"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(leave)",
		As:       []string{"leave"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsOwner:  true,
		IsGroup:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			err := client.WA.LeaveGroup(m.From)
			if err != nil {
				m.Reply("failed to get out of this group.")
				fmt.Println(err.Error())
			}
		},
	})
}
