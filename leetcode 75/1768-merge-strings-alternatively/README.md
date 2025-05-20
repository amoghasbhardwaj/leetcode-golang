---

## ✅ Approach

- Initialize two pointers `i` and `j` to iterate over `word1` and `word2` respectively.
- Use a loop that runs while either pointer is within bounds.
- Alternately append characters from each string if the current pointer is still within the string.
- When one string ends, the loop continues to append characters from the remaining string.
- Use a `[]byte` slice for efficient string building, then convert it back to a string.

---

## ⏱️ Time and Space Complexity

- **Time Complexity:** `O(m + n)` where `m` is the length of `word1` and `n` is the length of `word2`.  
  Each character is processed exactly once.

- **Space Complexity:** `O(m + n)` for storing the resulting merged string.