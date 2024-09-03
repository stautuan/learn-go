package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
    // get the length of each message, the length == cost
    // the cost accumulates with each message
    primaryCost := len(primary)
    secondaryCost := len(secondary) + primaryCost
    totalCost := len(tertiary) + secondaryCost

    message := [3]string{
        primary,
        secondary,
        tertiary,
    }

    cost := [3]int{
        primaryCost,
        secondaryCost,
        totalCost,
    }

    return message, cost
}