/*
Write a function that calculates how many messages we can send in a given batch
given a costMultiplier and a maxCostInPennies

The first message costs a penny, and each message after that costs the same as the
previous message multiplied by the costMultiplier
*/

package main

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
    actualCostInPennies := 1.0
    maxMessagesToSend := 1
    balance := float64(maxCostInPennies) - actualCostInPennies

    for balance >= 0 {
        actualCostInPennies *= costMultiplier
        balance -= actualCostInPennies
        maxMessagesToSend++
    }
    if balance <= 0  {
        maxMessagesToSend--
    }
    return maxMessagesToSend
}
