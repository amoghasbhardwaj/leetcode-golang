package main

import "fmt"

// gcdOfStrings finds the greatest common divisor string of two input strings
func gcdOfStrings(str1 string, str2 string) string {
    // If concatenating in different orders gives different results,
    // then there's no common base string that can form both.
    if str1+str2 != str2+str1 {
        return ""
    }

    // Find the GCD of the lengths of the two strings
    lengthGCD := gcd(len(str1), len(str2))

    // The common base string is the prefix of str1 up to lengthGCD
    return str1[:lengthGCD]
}

// gcd calculates the Greatest Common Divisor using the Euclidean algorithm
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    // Example test case
    fmt.Println(gcdOfStrings("ABCABC", "ABC")) // Output: "ABC"
}