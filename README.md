# Golang Tutorial

A comprehensive collection of Go tutorials covering basic concepts, data types, variables, constants, and functions. Includes Persian/Iranian utility packages for practical applications.

## Project Structure

```
Golang_tutorial/
├── 01-HelloWorld/          # Basic Go program with init() function
│   └── main.go
├── 02-moduleTest/          # Go module with Persian utility packages
│   ├── go.mod
│   ├── Banking/            # Iranian banking utilities
│   │   ├── card.go         # Card number processing
│   │   └── sheba.go        # SHEBA number processing
│   └── persianDate/        # Persian calendar utilities
│       └── dateConversion.go  # Miladi ↔ Shamsi conversion
├── 03-DataTypes/           # Go data types and their usage
│   ├── 01-basic/           # Basic integer and float types
│   │   └── integer.go
│   ├── 02-Enum/            # Enum-like patterns with constants
│   │   └── Enum.go
│   └── 03-Rune/            # Rune and string handling with Unicode
│       └── Example.go
├── 04-VariablesAndConstants/  # Variable and constant declarations
│   ├── Consts.go          # Constants in Go
│   └── Variable.go        # Different ways to declare variables
├── 05-Functions/           # Go functions and string operations
│   ├── Funcs.go           # String manipulation functions
│   └── print.go           # Different print methods
└── README.md
```

## Tutorials

### 01 - Hello World

A simple Go program demonstrating:
- Basic program structure
- The `init()` function (runs before main)
- Package and import declarations

**Run it:**
```bash
cd 01-HelloWorld
go run main.go
```

### 02 - Module Test (PersianUtil)

A Go module demonstrating package creation and organization with Persian/Iranian utilities.

**Packages included:**

| Package | Description |
|---------|-------------|
| `Banking` | Iranian banking utilities (card numbers, SHEBA numbers) |
| `persianDate` | Persian calendar conversion (Miladi ↔ Shamsi) |

**Run it:**
```bash
cd 02-moduleTest
go run .
```

### 03 - Data Types

Learn about different data types in Go:

#### 03.1 - Basic Integer and Float Types

Demonstrates various integer and float types:
- `int`, `int8`, `int16`, `int32`, `int64`
- `float32`, `float64`
- Using `unsafe.Sizeof()` to check memory size
- `math/bits.UintSize` for platform-dependent integer size

**Run it:**
```bash
cd 03-DataTypes/01-basic
go run integer.go
```

#### 03.2 - Enum-like Patterns

Shows how to create enums using constants and custom types:
- Defining constants for status values
- Creating custom types based on integers
- Using structs with enum-like fields
- JSON marshaling of custom types

**Run it:**
```bash
cd 03-DataTypes/02-Enum
go run Enum.go
```

#### 03.3 - Rune and String Handling

Learn about Unicode support in Go:
- `rune` type for Unicode code points
- String vs rune slice for proper Unicode handling
- Working with emojis and multi-byte characters
- `unsafe.Sizeof()` for string and rune slice

**Run it:**
```bash
cd 03-DataTypes/03-Rune
go run Example.go
```

### 04 - Variables and Constants

Master variable and constant declarations in Go:

#### 04.1 - Constants

Learn about constants:
- Declaring constants with `const`
- Grouping constants in parentheses
- Individual constant declarations
- Use cases: URLs, API endpoints, fixed values

**Run it:**
```bash
cd 04-VariablesAndConstants
go run Consts.go
```

#### 04.2 - Variables

Explore different ways to declare variables:
- Method 1: Explicit type declaration (`var name type = value`)
- Method 2: Declaration without initialization
- Method 3: Type inference
- Method 4: Short variable declaration (`:=`)
- Method 5: Group variables with multiple assignments
- Method 6: Group variables using parentheses

**Run it:**
```bash
cd 04-VariablesAndConstants
go run Variable.go
```

### 05 - Functions

Learn about functions and string operations:

#### 05.1 - String Manipulation

Comprehensive string functions from the `strings` package:
- `Contains`, `ContainsAny` - substring checks
- `Count` - counting occurrences
- `Cut` - splitting at first occurrence
- `Split`, `Join` - splitting and joining strings
- `Replace`, `ReplaceAll` - string replacement
- `Compare`, `EqualFold` - string comparison
- `HasPrefix`, `HasSuffix` - prefix/suffix checks
- `Index` - finding substring index
- `ToLower`, `ToUpper`, `Title` - case conversion
- `Trim`, `TrimLeft`, `TrimRight` - trimming characters

**Run it:**
```bash
cd 05-Functions
go run Funcs.go
```

#### 05.2 - Print Methods

Different ways to print output:
- `print()` - prints without newline
- `println()` - prints with newline
- `fmt.Printf()` - formatted printing with specifiers
- Format specifiers: `%s`, `%d`, `%f`, `%T`, `%b`

**Run it:**
```bash
cd 05-Functions
go run print.go
```

## Requirements

- Go 1.26.5 or higher

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/soheilsleep/Golang_tutorial.git
cd Golang_tutorial
```

2. Run any tutorial:
```bash
cd 01-HelloWorld
go run main.go
```

## License

This project is for educational purposes.