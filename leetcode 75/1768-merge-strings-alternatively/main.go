package main

import "fmt"

// mergeAlternately takes two strings and merges them by alternating characters
func mergeAlternately(word1 string, word2 string) string {
    m, n := len(word1), len(word2) // lengths of the two strings
    i, j := 0, 0                   // pointers for each string
    a := []byte{}                 // result byte slice

    // Loop until both strings are completely traversed
    for i < m || j < n {
        if i < m {
            a = append(a, word1[i]) // append character from word1
            i++
        }
        if j < n {
            a = append(a, word2[j]) // append character from word2
            j++
        }
    }

    return string(a) // convert byte slice to string and return
}

func main() {
    // Test example
    fmt.Println(mergeAlternately("abc", "pqr")) // Output: "apbqcr"
}