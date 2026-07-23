# Golang Tutorial

A collection of Go tutorials and Persian/Iranian utility packages for learning Go programming.

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
└── README.md
```

## Tutorials

### 01 - Hello World

A simple Go program demonstrating:
- Basic program structure
- The `init()` function
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

## Banking Package

Functions for working with Iranian banking systems:

```go
package banking

// GetBankName returns the bank name for a given card number
func GetBankName(cardNumber string) string

// GetShebaName returns the SHEBA bank name for a given number
func GetShebaName(cardNumber string) string
```

## Persian Date Package

Functions for converting between Gregorian and Persian (Jalali) calendars:

```go
package persiandate

// MiladiToShamsi converts Gregorian date to Shamsi (Persian) date
func MiladiToShamsi()

// ShamsiTOMiladi converts Shamsi (Persian) date to Gregorian date
func ShamsiTOMiladi()
```

## Requirements

- Go 1.26.5 or higher

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/your-username/Golang_tutorial.git
cd Golang_tutorial
```

2. Run any tutorial:
```bash
cd 01-HelloWorld
go run main.go
```

## License

This project is for educational purposes.
