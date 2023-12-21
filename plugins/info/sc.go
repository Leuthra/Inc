package info

import (
	"inc/lib"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(sc|source)",
		As:       []string{"sc"},
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			m.Reply("https://github.com/fckvania/MaoGo\n\n_Free Not For Sell_")
		},
	})
}
