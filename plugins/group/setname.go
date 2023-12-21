package group

import "inc/lib"

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:       "(setname|setnamegc|setnamegrup|setnamegroup)",
		As:         []string{"setnamegroup"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsQuerry:   true,
		IsAdmin:    true,
		IsGroup:    true,
		IsBotAdmin: true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			err := client.WA.SetGroupName(m.From, m.Querry)
			if err != nil {
				m.Reply("Failed to change the group name")
				return
			}
			m.Reply("Berhasil mengubah nama group")
		},
	})
}
