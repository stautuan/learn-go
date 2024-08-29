package main

import "fmt"

type expense interface {
    cost() int
}

type formatter interface {
    format() string
}

type email struct {
    isSubscribed bool
    body         string
}
func (e email) cost() int {
    lengthOfMessage := len(e.body)
    if e.isSubscribed == false {
        return lengthOfMessage * 5
    }
    return lengthOfMessage * 2
}
func (e email) format() string {
    subscribedText := "Subscribed"
    if e.isSubscribed == false {
        subscribedText = "Not Subscribed"
    }
    return fmt.Sprintf("'%s' | %s", e.body, subscribedText)
}