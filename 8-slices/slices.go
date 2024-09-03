/*
Textio's free users only get 1 retry message, while pro members get
an unlimited amount.

Write a function that takes a `plan` variable as input as well as an
array of 3 messages. You've been provided with constants representing
the plan types at the top of the file.

If the plan is a pro plan, return all the strings from messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a `nil` slice and an error that says `unsupported plan`.
*/

package main

import "errors"

const (
    planFree = "free"
    planPro = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
    if plan == planPro {
        return messages[:], nil
    }
    if plan == planFree {
        return messages[0:2], nil
    }
    return nil, errors.New("unsupported plan")
}