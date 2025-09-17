package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}

	return dm.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch val := n.(type) {
	case directMessage:
		return val.senderUsername, val.importance()
	case groupMessage:
		return val.groupName, val.importance()
	case systemAlert:
		return val.alertCode, val.importance()
	default:
		return "", 0
	}
}

func main() {
	dm1 := directMessage{
		senderUsername: "New User 1",
		messageContent: "Message Cont 1",
		priorityLevel:  30,
		isUrgent:       false,
	}

	fmt.Println(processNotification(dm1))

	gm1 := groupMessage{
		groupName:      "New Group 1",
		messageContent: "Message Cont 2",
		priorityLevel:  70,
	}

	fmt.Println(processNotification(gm1))

	sa1 := systemAlert{
		alertCode:      "Red",
		messageContent: "Message Cont 3",
	}

	fmt.Println(processNotification(sa1))
}
