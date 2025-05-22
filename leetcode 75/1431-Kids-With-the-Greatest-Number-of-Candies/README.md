# Kids With the Greatest Number of Candies

## 🧩 Problem Statement

You are given an array `candies` where `candies[i]` represents the number of candies the `i-th` kid has, and an integer `extraCandies`.

Return a boolean array `result` where `result[i]` is `true` if, after giving all the `extraCandies` to the `i-th` kid, they will have the greatest number of candies among all the kids. Otherwise, return `false`.

---

### 💡 Example:
Input: candies = [2, 3, 5, 1, 3], extraCandies = 3
Output: [true, true, true, false, true]

Explanation:
Kid 0: 2 + 3 = 5 → matches current max (5) → true
Kid 1: 3 + 3 = 6 → exceeds current max → true
Kid 2: 5 + 3 = 8 → exceeds → true
Kid 3: 1 + 3 = 4 → less than max → false
Kid 4: 3 + 3 = 6 → exceeds → true

---

## ✅ Approach

1. Find the current **maximum** number of candies any kid has.
2. For each kid, add `extraCandies` to their current candies.
3. If the result is **greater than or equal to the max**, mark `true`; else `false`.
4. Return the result as a boolean slice.

---

## 🧠 Why This Works

We only need to compare each kid’s candies + `extraCandies` with the **original max**, since we don’t actually mutate the array.  
No sorting or extra space needed.

---

## ⏱️ Time and Space Complexity

- **Time Complexity:** `O(n)`  
  - One pass to find the max  
  - One pass to compare each value

- **Space Complexity:** `O(n)`  
  - To store the boolean result array

---

## 🧪 Test Case

```go
Input: candies = [2, 3, 5, 1, 3], extraCandies = 3
Output: [true, true, true, false, true]