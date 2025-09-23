package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// IMPORTANT: There is a gotcha here. If you use "for i, message := range messages {" here, you will be getting a logical error since this loop will create a copy of each "message"
	// instead of modifying the original "messages". Hence, if you want to modify each "message" for "messages", you will need to use index
	for i := range messages {
		tags := tagger(messages[i])
		messages[i].tags = append(messages[i].tags, tags...)
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}

	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

func main() {
	testData := []sms{
		{id: "001", content: "Urgent, please respond!"},
		{id: "002", content: "Big sale on all items!"},
		{id: "003", content: "Enjoy your day"},
		{id: "004", content: "Sale! Don't miss out on these urgent promotions!"},
		{id: "005", content: "i nEEd URgEnt help, my FROZEN FLAME was used"},
		{id: "006", content: "wAnt to saLE 200x heavy leather"},
	}

	result := tagMessages(testData, tagger)
	fmt.Println(result)

}
