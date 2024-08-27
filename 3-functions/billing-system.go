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