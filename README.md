## 💻 Go Proficiency Enhancement

Structured Go proficiency enhancement to include CLI tooling and concurrency debugging.

Built over 45 exercises with full operational logs, screenshots, and debugging documentation.

🧠 Focus: Go for IaC & platform automation · 🪵 40+ logged troubleshooting events · ⚙️ Verified reproducible execution

## 🛠 Project Structure

The repository is organized to provide maximum transparency between code, execution, and troubleshooting.

```
├── Code
│   └── Go code and screenshots ranging from variables to structs; includes troubleshooting.
├── Concurrency
│   └── Production safety and troubleshooting.
├── Miscellaneous 
│   └── Compiler install and version validation.
├── Tooling 
│   └── Dual-purpose CLI tools implementing closures, higher-order functions, multiple return values, pointers, structs, deferred execution, and file logging.
├── OPS_LOG.md
├── README.md
```

---

## 🚀 Project Overview

This repository documents systematic Go proficiency development from fundamentals through concurrency patterns and CLI tooling. Every exercise includes executable code, output screenshots, and troubleshooting documentation.

**Built:** 28 code exercises, 13 debugging scenarios, 3 concurrency patterns, 1 production CLI tool  
**Documented:** 45 execution events with full screenshots and root cause analysis  
**Focus:** Go for Infrastructure as Code and platform automation

| Environment | Detail |
| :--- | :--- |
| **Operating System** | MacOS 26.2 (Tahoe) |
| **Golang Version** | go1.25.4 darwin/arm64 |

---

## 📚 Core Learning Concepts by Category

The following sections organize the learning artifacts by logical concept, demonstrating progression from basic syntax to concurrency patterns and CLI tooling.

### 1. Variables, Types, and I/O Fundamentals

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-04 | `goodEvening.go` | Basic package declaration and I/O using `fmt.Print/Println`. Includes initial troubleshooting/debugging on comment syntax (`#` vs `//`). |
| 2025-09-06 | `author.go`, `nameLocationOpinion.go`, `months.go` | Variable assignment, string interpolation, and I/O formatting with `fmt.Printf`. |
| 2025-09-07 | `shorthand.go` | Shorthand declaration (`:=`) for implicit type inference. |
| 2025-09-11 | `outerInnerBlocks.go` | Variable scope and shadowing in outer and inner code blocks. |
| 2025-09-11 | `zeroValues.go` | Default zero values for various data types (string, int, bool). |
| 2025-09-14 | `intToFloat64.go`, `float64ToInt.go`, `intToStr.go` | Explicit type casting between numeric and string types. |

### 2. Operators and Control Flow

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-20 | `arithmeticOperators.go` | Basic mathematical operations (+, -, *, /, %). |
| 2025-09-26 | `logicalOperators.go`, `assignmentOperators.go`, `bitwiseOperators.go` | Logical, assignment, and bitwise operators. |
| 2025-09-28 | `ifElseStatements.go` | Conditional logic using `if-else` blocks. **Includes troubleshooting on incorrect shorthand usage.** |
| 2025-09-28 | `switchCase.go` | Advanced `switch` statements with `fallthrough` for executing subsequent case blocks. |
| 2025-09-30 | `infiniteLoop.go` | Infinite loop structure for system utilities and event listeners. |

### 3. Loops and Data Structures

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-30 | `arrays.go` | Single- and multi-dimensional (2D) arrays, iteration with `for` loops, element assignment/replacement, `len()` usage. |
| 2025-10-03 | `forLoop.go` | For loop iteration patterns and break statements. |
| 2025-10-10 | (Advanced slices) | Slice creation, capacity management, append, element removal, copy, and iteration. **Includes troubleshooting on subslice indexing.** |
| 2025-10-10 | (Maps) | Complete map CRUD operations: create, read, update, delete, truncation, comma–ok idiom for existence checks. |

### 4. Functions and Advanced Concepts

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-11-19 | `multipleValuesReturned.go` | Functions returning multiple values. |
| 2025-11-19 | `multipleVariadicFunctions.go` | Variadic functions accepting variable number of arguments. |
| 2025-11-19 | `blankIdentifier.go` | Using `_` to ignore unwanted return values. |
| 2025-11-19 | `factorialRecursive.go` | Recursive function implementation. |
| 2025-11-19 | `anonymousFunctions.go` | Anonymous functions and closures. |
| 2025-11-27 | `addressOperators.go` | Declaring and initializing pointers, dereferencing, memory addresses. |

### 5. Structs and Methods

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-12-08 | `smithStruct.go` | Struct definition, instantiation methods, struct comparison with `==`/`!=`, slice of structs, iteration with range, Go-syntax formatting with `%#v`. **Includes troubleshooting on struct type mismatches.** |

### 6. Concurrency Patterns

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-12-22 | `closedChannels.go` | Buffered channels and closed channel detection with boolean checks. |
| 2025-12-22 | `timeOutCode.go` | Channel timeout patterns using `select` and `time.After`. |
| 2025-12-22 | `unbufferedChannelForLoop.go` | Unbuffered channel communication with for-range loops. |

### 7. CLI Tooling

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-12-07 | `tipTaxCalculator.go` | Production CLI tool implementing closures, higher-order functions, multiple return values, struct field access from slices. Real-world application: calculates sales tax and tips across 9 service types for Seattle metro area. **Includes troubleshooting on function signature matching and return value handling.** |

---

## 📊 Repository Metrics

| Category | Count | Documentation |
|----------|-------|---------------|
| **Code Exercises** | 28 | Complete execution screenshots (cde1-cde28) |
| **Debugging Scenarios** | 13 | Root cause analysis with fixes (db1-db13) |
| **Concurrency Patterns** | 3 | Production safety patterns (con1-con3) |
| **CLI Tools** | 1 | Full-featured calculator with closures (tool1) |
| **Total Logged Events** | 45 | OPS_LOG.md contains all screenshots and troubleshooting |

🧠 **Focus:** Go for IaC & platform automation  
🪵 **45 logged execution and troubleshooting events**  
⚙️ **Verified reproducible execution with screenshots**
---
