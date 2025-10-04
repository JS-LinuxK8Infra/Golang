# Go Language Systematic Learning & Development

**Supporting Repository for Resume Entry: Platform Engineering Lab (GCP/Kubernetes R&D)**

## 🚀 Project Overview

This repository documents a dedicated, systematic learning initiative for Go, with a focus on core concepts essential for modern Platform Engineering and **Infrastructure as Code (IaC). The goal is to build the knowledge necessary for developing high-performance, concurrent, and scalable systems tools (CLI utilities, API services, Kubernetes controllers).

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
| **`OPERATIONAL_LOG.md`** | **Full, chronological, dated log of all tasks, outputs, and troubleshooting steps.** |

---

## 📚 Core Learning Concepts by Category

The following sections organize the learning artifacts by logical concept, demonstrating progression from basic syntax to data structure manipulation.

### 1. Variables, Types, and I/O Fundamentals

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-04 | `good-evening.go`, `greeting.go` | Basic package declaration and I/O using `fmt.Print/Println`. Includes initial troubleshooting/debugging on comment syntax (`#` vs `//`). |
| 2025-09-06 | `name-location-opinion*.go` | Demonstrates variable assignment, string interpolation, and I/O formatting (newline characters and `fmt.Println`). |
| 2025-09-07 | `shorthand.go` | Usage of shorthand declaration (`:=`) for implicit type inference. |
| 2025-09-11 | `outer_inner_blocks.go` | Exploration of variable scope and shadowing in outer and inner code blocks. |
| 2025-09-11 | `zero_values.go` | Demonstrates default zero values for various data types (string, int, bool). |
| 2025-09-14 | `variable_types_typeof.go` | Explicitly verifies variable types using the `reflect.TypeOf` function. |
| 2025-09-14 | `int_to_float64.go`, `float64_to_int.go` | Demonstrates explicit type casting between numeric types. |
| 2025-09-14 | `str_to_int_noerr.go`, `str_to_int_err.go` | Uses `strconv.Atoi` for string-to-integer conversion, including structured error handling and intentionally triggering errors to demonstrate error flow. |
| 2025-09-14 | `constants.go`, `constants_error.go` | Defines typed and untyped constants, including demonstrating compile-time errors with mismatched constant types. |

### 2. Operators and Control Flow

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-17 | `comparison_operators.go` | Usage of all six comparison operators (==, !=, >, <, >=, <=) to return boolean values. |
| 2025-09-20 | `arithmatic_operators.go` | Execution of basic mathematical operations (+, -, *, /, %). |
| 2025-09-26 | `logical_operators.go` | Demonstrates logical operators (AND `&&`, OR `||`, NOT `!`) in conditional contexts. |
| 2025-09-26 | `assignment_operators.go`, `bitwise_operators.go` | Demonstrates all five assignment and bitwise operators. |
| 2025-09-28 | `if_else_statements.go` | Implements conditional logic using `if-else` blocks. **Includes troubleshooting on incorrect shorthand usage.** |
| 2025-09-28 | `switch_case.go` | Advanced usage of `switch` statements, including the use of `fallthrough` to execute subsequent case blocks. |
| 2025-09-30 | `infinite_loop.go` | Basic structure of an infinite loop (which will be leveraged later for listeners/polling in system utilities). |

### 3. Data Structures

| Date | File | Concept & Documentation |
| :--- | :--- | :--- |
| 2025-09-30 | `arrays.go` | Defines and manipulates single- and multi-dimensional (2D) arrays, including iteration using `for` loops and element assignment/replacement. |

---
