package owner

import (
	"inc/lib"
	"inc/lib/helpers"
	"strconv"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(setmode|mode)",
		As:       []string{"setmode"},
		Tags:     "owner",
		IsPrefix: true,
		IsOwner:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			if m.Querry == "public" {
				helpers.Public = true
				m.Reply("Public Mode: " + strconv.FormatBool(helpers.Public))
			} else {
				helpers.Public = false
				m.Reply("Public Mode: " + strconv.FormatBool(helpers.Public))
			}
		},
	})
}
