package group

import "inc/lib"

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:       "(setdesc|setdeskripsi|setdesk)",
		As:         []string{"setdesc"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsQuerry:   true,
		IsAdmin:    true,
		IsGroup:    true,
		IsBotAdmin: true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			err := client.WA.SetGroupTopic(m.From, "", "", m.Querry)
			if err != nil {
				m.Reply("Failed to change the group description")
				return
			}
			m.Reply("Successfully change the group description")
		},
	})
}
