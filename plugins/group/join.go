package group

import (
	"fmt"
	"inc/lib"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(join)",
		As:       []string{"join"},
		Tags:     "group",
		IsPrefix: true,
		IsOwner:  true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			gid, err := client.WA.JoinGroupWithLink(m.Querry)
			if err != nil {
				m.Reply("can't join the group.")
			} else {
				resp, _ := client.WA.GetGroupInfo(gid)
				m.Reply(fmt.Sprintf("successfully joined the group %s", resp.Name))
			}
		},
	})
}
