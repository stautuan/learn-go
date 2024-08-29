package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
    // msg is an instance of message
    messageText := msg.getMessage()
    cost := len(messageText) * 3 
    return messageText, cost
}

type message interface {
    getMessage() string
}

type birthdayMessage struct {
    birthdayTime  time.Time
    recipientName string
}
func (bm birthdayMessage) getMessage() string {
    return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
    reportName    string
    numberOfsends int
}
func (sr sendingReport) getMessage() string {
    return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfsends)
}
