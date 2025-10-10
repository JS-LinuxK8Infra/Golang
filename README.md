# Go Language Systematic Learning & Development

**Supporting Repository for Resume Entry: Platform Engineering Lab (GCP/Kubernetes R&D)**

## 🚀 Project Overview

This repository documents a dedicated, systematic learning initiative for Go, with a focus on core concepts essential for modern Platform Engineering and Infrastructure as Code (IaC). The goal is to build the knowledge necessary for developing high-performance, concurrent, and scalable systems tools (CLI utilities, API services, Kubernetes controllers).

| Environment | Detail |
| :--- | :--- |
| **Operating System** | Ubuntu 24.04 LTS |
| **Golang Version** | 1.22.2 |

---

## 🛠 Project Structure

The repository is organized to provide maximum transparency between code, execution, and troubleshooting.

| Folder | Description |
| :--- | :--- |
| `Code/` | Contains foundational code and structured Go files, including data structures and type conversions. |
| `Scripts/` | Contains various scripts focusing on I/O, variables, and control flow mechanics. |
| `Screenshots/` | Visual evidence of all code execution and troubleshooting steps (referenced in the `OPERATIONAL_LOG.md`). |
| `OPERATIONS_LOG.md` | Full, chronological, dated log of all tasks, outputs, and troubleshooting steps. |

---

## 📚 Core Learning Concepts by Category

The following sections organize the learning artifacts by logical concept, demonstrating progression from basic syntax to data structure manipulation.

### 1. Variables, Types, and I/O Fundamentals

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-04 | `goodEvening.go`, `greeting.go` | Basic package declaration and I/O using `fmt.Print/Println`. Includes initial troubleshooting/debugging on comment syntax (`#` vs `//`). |
| 2025-09-06 | `nameLocationOpinion*.go` | Demonstrates variable assignment, string interpolation, and I/O formatting (newline characters and `fmt.Println`). |
| 2025-09-07 | `shorthand.go` | Usage of shorthand declaration (`:=`) for implicit type inference. |
| 2025-09-11 | `outerInnerBlocks.go` | Exploration of variable scope and shadowing in outer and inner code blocks. |
| 2025-09-11 | `zeroValues.go` | Demonstrates default zero values for various data types (string, int, bool). |
| 2025-09-14 | `dataTypes.go` | Explicitly verifies variable types using the `reflect.TypeOf` function. |
| 2025-09-14 | `intToFloat64.go`, `float64ToInt.go` | Demonstrates explicit type casting between numeric types. |
| 2025-09-14 | `strIntAtoi.go`, `strToIntErr.go` | Uses `strconv.Atoi` for string-to-integer conversion, including structured error handling and intentionally triggering errors to demonstrate error flow. |
| 2025-09-14 | `constants.go`, `constantsError.go` | Defines typed and untyped constants, including demonstrating compile-time errors with mismatched constant types. |

### 2. Operators and Control Flow

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-17 | `comparisonOperators.go` | Usage of all six comparison operators (==, !=, >, <, >=, <=) to return boolean values. |
| 2025-09-20 | `arithmeticOperators.go` | Execution of basic mathematical operations (+, -, *, /, %). |
| 2025-09-26 | `logicalOperators.go` | Demonstrates logical operators (AND `&&`, OR `||`, NOT `!`) in conditional contexts. |
| 2025-09-26 | `assignmentOperators.go`, `bitwise_operators.go` | Demonstrates all five assignment and bitwise operators. |
| 2025-09-28 | `ifElseStatements.go` | Implements conditional logic using `if-else` blocks. **Includes troubleshooting on incorrect shorthand usage.** |
| 2025-09-28 | `switchCase.go` | Advanced usage of `switch` statements, including the use of `fallthrough` to execute subsequent case blocks. |
| 2025-09-30 | `infiniteLoop.go` | Basic structure of an infinite loop (which will be leveraged later for listeners/polling in system utilities). |

### 3. Data Structures
| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-30 | `arrays.go` | Defines and manipulates single- and multi-dimensional (2D) arrays, including iteration using `for` loops and element assignment/replacement. Demonstrates `len()` usage, index element swapping, and nested (2D) array access. |
| 2025-10-10 | `slices.go`, `maps.go` | Advanced data structure handling — demonstrates slice creation, append/capacity behavior, and complete map CRUD functionality (create, read, update, delete, iteration, truncation). |

---
