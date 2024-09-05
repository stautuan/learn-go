/*
Textio is introducing a feature that allows users to filter their messages
based on specific criteria. For this feature, messages are categorized into 
three types: `TextMessage`, `MediaMessage`, and `LinkMessage`. Users can
filter their messages to view only the types they are interested in.

Implement a function that filters a slice of messages based on the message type.

It should take a slice of `Message` interfaces and a string indicating the
desired type ("text", "media", or "link"). It should return a new slice of
`Message` interfaces containing only messages of the specified type.
*/

/* 
Category Type:
TextMessage
MediaMessage
LinkMessage
*/

package main

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
func (nm MediaMessage) Type() string {
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

// function takes a slice of Message interfaces and a string indicating a desired type
// function returns a new slice of Message interfaces
func filterMessages(messages []Message, filterType string) []Message {
    // create an empty slice to hold our filtered Messages
    filteredMessages := []Message{}
    // iterate through the messages
    for _, message := range messages {
        // check if the type of the message matches the desired type
        if message.Type() == filterType {
            // append the message if it matches
            filteredMessages = append(filteredMessages, message)
        }
    }
    return filteredMessages
}