package info

import (
	"fmt"
	"inc/lib"
	"regexp"
	"strings"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(channelinfo|ci)",
		As:       []string{"ci"},
		Tags:     "info",
		IsPrefix: true,
		IsQuery:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			pattern := regexp.MustCompile(`https?://whatsapp.com/channel/`)
			if !pattern.MatchString(m.Query) {
				m.Reply("Url Invalid")
				return
			}

			key, err := client.WA.GetNewsletterInfoWithInvite(strings.Split(m.Query, "/")[4])
			if err != nil {
				m.Reply("i don't know")
				return
			}

			m.Reply(fmt.Sprintf("*Channel Information*\n*Link:* %s\n*ID:* %s\n*Name:* %v\n*Followers:* %v\n\n*Description:* %v\n*Create At:* %v", m.Query, key.ID, key.ThreadMeta.Name.Text, key.ThreadMeta.SubscriberCount, key.ThreadMeta.Description.Text, key.ThreadMeta.CreationTime))
		},
	})
}
