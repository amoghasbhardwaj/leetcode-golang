# GCD of Strings

## Problem Statement

For two strings `str1` and `str2`, a string `x` is said to be a common divisor of `str1` and `str2` if:

- `str1 = x + x + ...` (x repeated some number of times)
- `str2 = x + x + ...` (x repeated some number of times)

Find the largest such string `x` (i.e., the greatest common divisor string). If no such string exists, return an empty string.

### Example:
Input: str1 = “ABCABC”, str2 = “ABC”
Output: “ABC”

Input: str1 = “ABABAB”, str2 = “ABAB”
Output: “AB”

Input: str1 = “LEET”, str2 = “CODE”
Output: “”

---

## ✅ Approach

1. **String Compatibility Check**:
   - If a valid GCD string `x` exists, then `str1 + str2 == str2 + str1` must hold.
   - This ensures both strings are made up of repeated copies of the same substring.

2. **GCD Length**:
   - Compute the greatest common divisor (GCD) of the lengths of the two strings using the **Euclidean Algorithm**.
   - The GCD of lengths corresponds to the length of the largest substring that can divide both strings.

3. **Return the Prefix**:
   - The potential common base string is simply the prefix of `str1` up to the GCD length.

---

## 🧠 Why GCD Works Here

Let’s say:
- `str1 = x * m`
- `str2 = x * n`

Where `x` is the GCD string and `m, n` are repeat counts. Then the length of `x` is the GCD of `len(str1)` and `len(str2)`, so `x = str1[:gcd(len1, len2)]`.

---

## ⏱️ Time and Space Complexity

- **Time Complexity:**  
  - `O(m + n)` for the string concatenation check (`str1+str2 != str2+str1`)  
  - `O(log(min(m, n)))` for computing the GCD  
  - `O(g)` to slice the result (where `g = GCD(m, n)`)

- **Space Complexity:** `O(1)` (excluding result string)

---

## 🔧 Test Case
```go
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
---