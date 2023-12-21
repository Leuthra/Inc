package group

import (
	"fmt"
	"inc/lib"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:       "(revoke|resetlink)",
		As:         []string{"revoke"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsAdmin:    true,
		IsBotAdmin: true,
		IsGroup:    true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			resp, err := client.WA.GetGroupInviteLink(m.From, true)
			if err != nil {
				m.Reply("Failed to reset the link group.")
			} else {
				m.Reply(fmt.Sprintf("New group link: %s", resp))
			}
		},
	})
}
