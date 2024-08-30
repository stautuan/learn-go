/*
We offer a product that allows businesses to send pairs of messages to couples. 
It is mostly used by flower shops and movie theaters.

Write a function that sends 2 messages, first to the customer and then to the customer's spouse.
    1. Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
    2. Do the same for the msgToSpouse
    3. If both messages are sent successfully, return the total cost of the messages added together.
*/

package main

import (
    "fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {

    iCustomer, errCustomer := sendSMS(msgToCustomer)
    if errCustomer != nil {
        return 0, errCustomer
    }
    iSpouse, errSpouse := sendSMS(msgToSpouse)
    if errSpouse != nil {
        return 0, errSpouse
    }
    return (iCustomer + iSpouse), nil

}

func sendSMS(message string) (int, error) {
    const maxTextLen = 25
    const costPerChar = 2
    if len(message) > maxTextLen {
        return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
    }
    return costPerChar * len(message), nil
}