package download

import (
	"inc/lib"
	"inc/lib/api"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(tt|tiktok)",
		As:       []string{"tiktok"},
		Tags:     "downloader",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			url, err := api.GetTiktokVideo(m.Querry)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			bytes, err := client.GetBytes(url)
			if err != nil {
				m.Reply(err.Error())
				return
			}
			client.SendVideo(m.From, bytes, "", nil)
		},
	})
}