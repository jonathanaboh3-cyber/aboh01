package main

import "os"

// parseNumber converts a string to int64, handling signs and checking for validity
func parseNumber(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	// Handle sign
	neg := false
	start := 0
	if s[0] == '-' {
		neg = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start == len(s) {
		return 0, false
	}

	// Parse digits
	result := int64(0)
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int64(s[i]-'0')
		// Rough overflow check during parsing
		if result > 9223372036854775807 {
			return 0, false
		}
	}

	if neg {
		result = -result
	}
	return result, true
}

// int64ToString converts an int64 to a string without using fmt or strconv
func int64ToString(n int64) string {
	if n == 0 {
		return "0"
	}

	// Handle negative numbers
	neg := n < 0
	if neg {
		n = -n
	}

	// Convert digits to string (in reverse order)
	digits := make([]byte, 0, 20)
	for n > 0 {
		digit := byte(n%10) + '0'
		digits = append(digits, digit)
		n /= 10
	}

	// Reverse the digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// Add negative sign if needed
	if neg {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}

func main() {
	// Check for exactly 3 arguments (plus program name)
	if len(os.Args) != 4 {
		return
	}

	// Parse first value
	num1, ok1 := parseNumber(os.Args[1])
	if !ok1 {
		return
	}

	// Validate operator
	op := os.Args[2]
	if len(op) != 1 || (op != "+" && op != "-" && op != "*" && op != "/" && op != "%") {
		return
	}

	// Parse second value
	num2, ok2 := parseNumber(os.Args[3])
	if !ok2 {
		return
	}

	// Constants for int64 bounds
	const maxInt64 = 9223372036854775807
	const minInt64 = -maxInt64 - 1

	// Buffer for output
	var output []byte

	switch op {
	case "+":
		// Check overflow/underflow
		if num1 > 0 && num2 > maxInt64-num1 {
			return
		}
		if num1 < 0 && num2 < minInt64-num1 {
			return
		}
		output = []byte(int64ToString(num1 + num2))
	case "-":
		// Check overflow/underflow
		if num1 > 0 && num2 < -(maxInt64-num1) {
			return
		}
		if num1 < 0 && num2 > maxInt64+num1 {
			return
		}
		output = []byte(int64ToString(num1 - num2))
	case "*":
		// Handle zero case
		if num1 == 0 || num2 == 0 {
			output = []byte("0")
		} else {
			// Check overflow for positive/negative combinations
			if num1 > 0 && num2 > 0 && num1 > maxInt64/num2 {
				return
			}
			if num1 < 0 && num2 < 0 && num1 < maxInt64/num2 {
				return
			}
			if num1 > 0 && num2 < 0 && num2 < minInt64/num1 {
				return
			}
			if num1 < 0 && num2 > 0 && num1 < minInt64/num2 {
				return
			}
			output = []byte(int64ToString(num1 * num2))
		}
	case "/":
		if num2 == 0 {
			output = []byte("No division by 0")
		} else {
			output = []byte(int64ToString(num1 / num2))
		}
	case "%":
		if num2 == 0 {
			output = []byte("No modulo by 0")
		} else {
			output = []byte(int64ToString(num1 % num2))
		}
	}

	// Append newline and write to stdout
	output = append(output, '\n')
	os.Stdout.Write(output)
}
