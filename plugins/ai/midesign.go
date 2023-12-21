package ai

import (
	"inc/lib"
	"inc/lib/api"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(miscrosoftdesigner|midesign)",
		As:       []string{"midesign"},
		Tags:     "ai",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			data, err := api.MicrosoftDesigner(m.Querry)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			buffer, err := client.GetBytes(data["image_urls_thumbnail"].([]interface{})[0].(map[string]interface{})["ImageUrl"].(string))
			if err != nil {
				m.Reply(err.Error())
				return
			}
			client.SendImage(m.From, buffer, m.Querry, m.ID)
		},
	})
}