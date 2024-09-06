package main

func isValidPassword(password string) bool {
    length := len(password)
    hasUpper := false
    hasDigit := false

    if length > 5 && length < 12 {
        for i := 0; i < length; i++ {
            if password[i] >= 'A' && password[i] <= 'Z' {
                hasUpper = true
            }
            if password[i] >= '0' && password[i] <= '9' {
                hasDigit = true
            }
            if hasUpper && hasDigit {
                return true
            }
        }
    }
    return false
}