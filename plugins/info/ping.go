package info

import (
	"fmt"
	"inc/lib"
	"time"
)

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "ping",
		As:       []string{"ping"},
		Tags:     "info",
		IsPrefix: true,
		Exec: func(client *lib.Event, m *lib.IMessage) {
			now := time.Now()
				mdate := time.Unix(m.Info.Timestamp.Unix(), 0)
				mtime := now.Sub(mdate)
				ms := mtime.Seconds()
				txt := fmt.Sprintf("%.3f seconds", ms)
				m.Reply(txt)
		},
	})
}
