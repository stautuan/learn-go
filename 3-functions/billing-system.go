/*
Textio is making some changes to how they bill users for bulk messages. 
The system should now calculate the bill based on:
    The number of messages sent
    The cost per message
    Thresholds for discounts

Discount rates:
    10% for more than 1000 messages.
    20% for more than 5000 messages.
    0% for anything less.
*/


package main


func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
    cost := calculateBaseBill(costPerMessage, numMessages)
    discountRate := calculateDiscountRate(numMessages)
    final := cost * discountRate
    return final
}

func calculateDiscountRate(messagesSent int) float64 {
    if messagesSent > 5000 {
        return 0.8
    }
    if messagesSent > 1000 {
        return 0.9
    }
    return 1.0
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
    return costPerMessage * float64(messagesSent)
}