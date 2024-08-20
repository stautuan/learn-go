package main

import "fmt"

func main() {
    senderName := "Syl"
    recipient := "Kaladin"
    message1 := "The Words, Kaladin. You have to speak the Words!"

    fmt.Printf("%s to %s: %s\n", senderName, recipient, message1)

    fname := "Dalinar"
    lname := "Kholinar"
    age := 45
    messageRate := 0.5
    isSubscribed := false
    message2 := "Sometimes a hypocrite is nothing more than a man in the process of changing."

    userLog := fmt.Sprintf(
        "Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s",
        fname,
        lname,
        age,
        messageRate,
        isSubscribed,
        message2,
    )
    
    fmt.Println(userLog)
}