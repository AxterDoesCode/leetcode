package main

//This solution uses a map as a lookup table.
//If it is the opening bracket then append it to the stack
//If it isn't and opening bracket check if the stack is empty or if the previous item in the stack
//Isn't the matching pair then break and return false immediately
//The last else statement is if it is the matching pair then remove the last entry from the top of the stack

func isValid(s string) bool {
    if len(s) == 0 || len(s)%2 == 1 {
        return false
    }

    stack := []rune{}

    pairs := map[rune]rune {
        '(':')',
        '{':'}',
        '[':']',
    }

    for _, r := range s {
        if _, ok := pairs[r]; ok {
            stack = append(stack, r)
        } else if len(stack) == 0 || pairs[stack[len(stack) - 1]] != r {
            return false
        } else {
            stack = stack[:len(stack) - 1]
        }
    }
    return len(stack) == 0
}
