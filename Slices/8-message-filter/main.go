package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	filteredMsg := []Message{} // empty slice of type Message struct

	for _, message := range messages {
		if message.Type() == filterType {
			filteredMsg = append(filteredMsg, message)
		}
	}

	return filteredMsg
}

func main() {
	msgs := []Message{
		TextMessage{
			Sender:  "Text1",
			Content: "Content1",
		},
		MediaMessage{
			Sender:    "Text4",
			MediaType: "MediaType2",
			Content:   "Content4",
		},
		LinkMessage{
			Sender:  "Text6",
			URL:     "URL2",
			Content: "Content6",
		},
		TextMessage{
			Sender:  "Text2",
			Content: "Content2",
		},
		MediaMessage{
			Sender:    "Text3",
			MediaType: "MediaType1",
			Content:   "Content3",
		},
		LinkMessage{
			Sender:  "Text5",
			URL:     "URL1",
			Content: "Content5",
		},
	}

	filteredMsgs := filterMessages(msgs, "link")
	fmt.Println(filteredMsgs)

	return
}
