package group

import (
	"fmt"
	"inc/lib"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:       "(linkgroup|linkgrup|linkgc)",
		As:         []string{"linkgroup"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsGroup:    true,
		IsBotAdmin: true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			resp, err := client.WA.GetGroupInviteLink(m.From, false)
			if err != nil {
				m.Reply("Failed to get the group link.")
			} else {
				m.Reply(fmt.Sprintf("Link group: %s", resp))
			}
		},
	})
}
