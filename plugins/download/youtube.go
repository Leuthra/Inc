package download

import (
	"fmt"
	"inc/lib"
	"inc/lib/api"
	"regexp"
	"time"

	"github.com/Pauloo27/searchtube"
)

type method string

const (
	audio string = "Audio"
	video string = "Video"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "(ytmp4|ytmp3|play)",
		As:       []string{"ytmp4", "ytmp3", "play"},
		Tags:     "downloader",
		IsPrefix: true,
		IsQuery:  true,
		IsWaitt:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			var url string
			if api.IsYoutubeURL(m.Query) {
				url = m.Query
			} else {
				ser, _ := searchtube.Search(m.Query, 10)
				if len(ser) == 0 {
					m.Reply("Not Found")
					return
				}
				for _, v := range ser {
					d, _ := v.GetDuration()
					if d < 8*time.Minute {
						url = v.URL
						break
					}
				}
			}
			yt, err := api.YoutubeDL(url)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			caption := fmt.Sprintf("*Title*: %s\n*Author*: %s", yt.Info.Title, yt.Info.Author)

			if reg, _ := regexp.MatchString(`(ytmp3|play)`, m.Command); reg {
				build, err := yt.Link.Audio[0].Url()
				if err != nil {
					m.Reply(err.Error())
					return
				}

				bytes, err := client.GetBytes(build)
				if err != nil {
					m.Reply(err.Error())
					return
				}
				client.SendDocument(m.From, bytes, fmt.Sprintf("%s.mp3", yt.Info.Title), caption, m.ID)
			} else {
				build, err := yt.Link.Video[0].Url()
				if err != nil {
					m.Reply(err.Error())
					return
				}
				bytes, err := client.GetBytes(build)
				if err != nil {
					m.Reply(err.Error())
					return
				}
				client.SendVideo(m.From, bytes, caption, m.ID)
			}
		},
	})
}
