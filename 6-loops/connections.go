/*
Write a function that takes an integer groupSize representing the number
of people in the group chat and returns an integer representing the number
of connections between them. For each additional person in the group, the 
number of new connections continues to grow. Use a "for" loop to accumulate
the number of connections instead of directly using a mathematical formula.
*/

package main

func countConnections(groupSize int) int {
    connections := 0
    for i := 0; i <= groupSize; i++ {
        connections += 1
    }
    return connections
}