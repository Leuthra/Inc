package plugins

import (
	"fmt"
	"inc/lib"
	"sort"
	"strings"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

type item struct {
	Name     []string
	IsPrefix bool
}

type tagSlice []string

func (t tagSlice) Len() int {
	return len(t)
}

func (t tagSlice) Less(i int, j int) bool {
	return t[i] < t[j]
}

func (t tagSlice) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

func menu(client *lib.Event, m *lib.IMessage) {
	var str string
	str += fmt.Sprintf("Hello %s\n\n⇒ Library: whatsmeow\n⇒ Language: Golang\n\n", m.PushName)
	var tags map[string][]item
	for _, list := range lib.GetList() {
		if tags == nil {
			tags = make(map[string][]item)
		}
		if _, ok := tags[list.Tags]; !ok {
			tags[list.Tags] = []item{}
		}
		tags[list.Tags] = append(tags[list.Tags], item{Name: list.As, IsPrefix: list.IsPrefix})
	}

	var keys tagSlice
	for key := range tags {
		keys = append(keys, key)
	}

	sort.Sort(keys)

	for _, key := range keys {
		str += fmt.Sprintf("「 *%s MENU* 」\n", strings.ToUpper(key))
		for _, e := range tags[key] {
			var prefix string
			if e.IsPrefix {
				prefix = m.Command[:1]
			} else {
				prefix = ""
			}
			for _, nm := range e.Name {
				str += fmt.Sprintf("ゝ *%s%s*\n", prefix, nm)
			}
		}
		str += "\n"
	}
	var isImage = waProto.ContextInfo_ExternalAdReplyInfo_IMAGE
	client.SendText(m.From, strings.TrimSpace(str), &waProto.ContextInfo{
		ExternalAdReply: &waProto.ContextInfo_ExternalAdReplyInfo{
			Title:                 proto.String("Inc Bot 2023"),
			Body:                  proto.String("Simple WhatsApps Bots"),
			MediaType:             &isImage,
			ThumbnailUrl:          proto.String("https://i.pinimg.com/564x/23/44/ff/2344ff68c03a2b12fc66f578ca986008.jpg"),
			MediaUrl:              proto.String("https://wa.me/stickerpack/inc.dev"),
			SourceUrl:             proto.String("https://id.pinterest.com/pin/1150669773530593386/"),
			ShowAdAttribution:     proto.Bool(true),
			RenderLargerThumbnail: proto.Bool(true),
		}})
}

func init() {
	lib.NewCommands(&lib.ICommand{
		Name:     "menu",
		As:       []string{"menu"},
		Tags:     "main",
		IsPrefix: true,
		Exec:     menu,
	})
}
