/*
Write a function using the `append()` function. It accepts a slice of `cost` structs 
and returns a slice of `float64`s, where `float64` represents the total cost for a day.

The length of the returned slice should be inferred from the highest numbered `day`
field. Some days may have multiple costs, while others may have none. Include all
actual and implied totals in the returned slice, starting with day '0'. Use the 
`append()` function to dynamically increase the size of the returned slice when needed.
*/

package main

type cost struct {
    day   int
    value float64
}

func getCostsByDay(costs []cost) []float64 {

    costsByDay := []float64{} 
    for i := 0; i < len(costs); i++ {
        cost := costs[i]
        for cost.day <= len(costsByDay) {
            costsByDay = append(costsByDay, 0.0)
        }
        costsByDay[cost.day] += cost.value
    }
    return costsByDay

}