/* 
We need a way to differentiate between standard and premium users. When a new user 
is created, they need a membership type, and that type will determine the message
character limit.
*/

package main

type User struct {
    Name string
    Membership

    // User embeds Membership (it inherits all of its data fields)
}

type Membership struct {
    Type             string
    MessageCharLimit int
}

func newUser(name string, membershipType string) User {
    membership := Membership{}
    membership.Type = membershipType

    if membershipType == "premium" {
        membership.MessageCharLimit = 1000
    } else {
        membership.Type = "standard"
        membership.MessageCharLimit = 100
    }
    return User{Name: name, Membership: membership}
}