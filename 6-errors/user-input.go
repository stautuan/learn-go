/*
In Textio, users can set their profile status to communicate their current activity 
to those that choose to read it... However, there are some restrictions on what these 
statuses can contain. Your task is to implement a function that validates a user's 
status update. The status update cannot be empty and must not exceed 140 characters.
*/

package main

import "errors"

func validateStatus(status string) error {
    if len(status) == 0 {
        return errors.New("status cannot be empty")
    } else if len(status) > 140 {
        return errors.New("status exceeds 140 characters")
    } else {
        return nil
    }
}