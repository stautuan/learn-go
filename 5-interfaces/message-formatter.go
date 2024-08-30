/*
As Textio evolves, the team has decided to introduce a new feature for custom message 
formats. Depending on the user's preferences, messages can be sent in different formats, 
such as plain text, markdown, code, or even encrypted text. To efficiently manage this, 
you'll implement a system using interfaces.
*/

package main

import "fmt"

type Formatter interface {
    Format() string
}

type PlainText struct {
    message string
}
func (p PlainText) Format() string {
    return p.message
}

type Bold struct {
    message string
}
func (b Bold) Format() string {
    return fmt.Sprintf("**%s**", b.message)
}

type Code struct {
    message string
}
func (c Code) Format() string {
    return fmt.Sprintf("`%s`", c.message)
}

func SendMessage(formatter Formatter) string {
    return formatter.Format() // Adjusted to call Format without an argument
}