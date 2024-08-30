/*
Textio now has a system to process different types of notifications: direct messages, 
group messages, and system alerts. Each notification type has a unique way of calculating 
its importance score based on user interaction and content.

1. Implement the importance method for each type, and return the importance score (int) for each type
2. processNotification function should identify the type and return different values for each type
*/

package main

type notification inferface {
    importance() int
}

type directMessage struct {
    senderUsername string
    messageContent string
    priorityLevel  int
    isUrgent       bool
}
func (dm directMessage) importance() int {
    if dm.isUrgent {
        return 50
    }
    return dm.priorityLevel
}

type groupMessage struct {
    groupName      string
    messageContent string
    priorityLevel  int
}
func (gm groupMessage) importance() int {
    return gm.priorityLevel
}

type systemAlert struct {
    alertCode      string
    messageContent string
}
func (s systemAlert) importance() int {
    return 100
}

func processNotification(n notification) (string, int) {
    switch v = n.(type) {
    case directMessage:
        return v.senderUsername, v.importance()
    case groupMessage:
        return v.groupName, v.importance()
    case systemAlert:
        return v.alertCode, v.importance()
    default:
        return "", 0
    }
}