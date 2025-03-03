`testify/assert` is a powerful assertion library in Go that simplifies writing tests. It provides a variety of assertion functions (operators) to compare expected and actual values.

---

## **📌 Categories of Operators in `testify/assert`**
1. **Basic Assertions**  
2. **Boolean Assertions**  
3. **Numeric Assertions**  
4. **String Assertions**  
5. **Error Assertions**  
6. **Slice & Map Assertions**  
7. **Time Assertions**  
8. **Type Assertions**  
9. **Comparison Assertions**  
10. **Panic & Recover Assertions**  

---

## **1️⃣ Basic Assertions**
| Operator                     | Description | Example |
|------------------------------|------------|---------|
| `assert.True(t, cond)`       | Checks if `cond` is `true` | `assert.True(t, 5 > 3, "Expected true")` |
| `assert.False(t, cond)`      | Checks if `cond` is `false` | `assert.False(t, 5 < 3, "Expected false")` |
| `assert.Equal(t, expected, actual)` | Checks if two values are equal | `assert.Equal(t, "hello", "hello")` |
| `assert.NotEqual(t, expected, actual)` | Checks if two values are **not** equal | `assert.NotEqual(t, 10, 20)` |

---

## **2️⃣ Boolean Assertions**
| Operator                     | Description | Example |
|------------------------------|------------|---------|
| `assert.True(t, condition)`  | Asserts a boolean value is `true` | `assert.True(t, 1 == 1)` |
| `assert.False(t, condition)` | Asserts a boolean value is `false` | `assert.False(t, 1 > 2)` |

---

## **3️⃣ Numeric Assertions**
| Operator                        | Description | Example |
|---------------------------------|------------|---------|
| `assert.Greater(t, a, b)`      | Checks if `a > b` | `assert.Greater(t, 10, 5)` |
| `assert.GreaterOrEqual(t, a, b)` | Checks if `a >= b` | `assert.GreaterOrEqual(t, 10, 10)` |
| `assert.Less(t, a, b)`         | Checks if `a < b` | `assert.Less(t, 5, 10)` |
| `assert.LessOrEqual(t, a, b)`  | Checks if `a <= b` | `assert.LessOrEqual(t, 10, 10)` |

---

## **4️⃣ String Assertions**
| Operator                       | Description | Example |
|--------------------------------|------------|---------|
| `assert.Contains(t, str, substr)` | Checks if `substr` exists in `str` | `assert.Contains(t, "hello world", "world")` |
| `assert.NotContains(t, str, substr)` | Checks if `substr` **does not** exist in `str` | `assert.NotContains(t, "hello", "bye")` |
| `assert.Equal(t, expected, actual)` | Compares two strings | `assert.Equal(t, "hello", "hello")` |

---

## **5️⃣ Error Assertions**
| Operator                   | Description | Example |
|----------------------------|------------|---------|
| `assert.Error(t, err)`     | Asserts `err` is **not nil** | `assert.Error(t, someFuncThatReturnsError())` |
| `assert.NoError(t, err)`   | Asserts `err` is `nil` | `assert.NoError(t, someFuncThatReturnsNil())` |
| `assert.EqualError(t, err, expected)` | Checks if `err.Error()` matches `expected` | `assert.EqualError(t, err, "file not found")` |

---

## **6️⃣ Slice & Map Assertions**
| Operator                       | Description | Example |
|--------------------------------|------------|---------|
| `assert.Len(t, obj, length)`  | Checks if length matches | `assert.Len(t, []int{1, 2, 3}, 3)` |
| `assert.ElementsMatch(t, slice1, slice2)` | Checks if slices contain the same elements (ignoring order) | `assert.ElementsMatch(t, []int{1,2,3}, []int{3,2,1})` |
| `assert.Subset(t, superset, subset)` | Checks if `subset` is in `superset` | `assert.Subset(t, []int{1,2,3,4}, []int{2,4})` |
| `assert.Contains(t, map, key)` | Checks if a map contains a key | `assert.Contains(t, map[string]int{"a":1}, "a")` |

---

## **7️⃣ Time Assertions**
| Operator                          | Description | Example |
|-----------------------------------|------------|---------|
| `assert.WithinDuration(t, time1, time2, delta)` | Checks if two times are within `delta` duration | `assert.WithinDuration(t, time1, time2, time.Second)` |
| `assert.WithinRange(t, time, start, end)` | Checks if a time is within a range | `assert.WithinRange(t, time.Now(), startTime, endTime)` |

---

## **8️⃣ Type Assertions**
| Operator                    | Description | Example |
|-----------------------------|------------|---------|
| `assert.IsType(t, expected, actual)` | Checks if `actual` is of `expected` type | `assert.IsType(t, int(0), 123)` |

---

## **9️⃣ Comparison Assertions**
| Operator                   | Description | Example |
|----------------------------|------------|---------|
| `assert.Same(t, ptr1, ptr2)` | Checks if two pointers reference the same object | `assert.Same(t, &a, &a)` |
| `assert.NotSame(t, ptr1, ptr2)` | Checks if two pointers reference different objects | `assert.NotSame(t, &a, &b)` |

---

## **🔟 Panic & Recover Assertions**
| Operator                     | Description | Example |
|------------------------------|------------|---------|
| `assert.Panics(t, func())`  | Asserts that a function **panics** | `assert.Panics(t, func() { panic("error") })` |
| `assert.PanicsWithValue(t, expected, func())` | Asserts that the panic message matches `expected` | `assert.PanicsWithValue(t, "error", func() { panic("error") })` |
| `assert.NotPanics(t, func())` | Asserts that a function **does not panic** | `assert.NotPanics(t, func() { fmt.Println("safe") })` |

---
