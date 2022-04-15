package FinanransdonScriptLib

import (
	"fmt"
	"strings"
	"unicode"
)

func pendChar(c rune) int {
	if c == '{' {
		return 1
	}
	if c == '*' {
		return 2
	}
	if c >= '0' && c <= '9' {
		// Number
		return 3
	}
	if unicode.ToUpper(c) >= 'A' && unicode.ToUpper(c) <= 'F' {
		// Hexadecimal Digits
		return 4
	}
	if c == '(' {
		return 5
	}
	if c == ')' {
		return 6
	}
	if c == ' ' || c == '\t' || c == '　' {
		return 7
	}
	if c == '\n' {
		return 8
	}
	return -1
}

func TokenizeCode(code string) ([]Token, bool) {
	result := []Token{}
	errors := false
	stateCode := 0
	clipBuffer := ""
	for i := 0; i < len(code); i++ {
		currentChar := code[i]
		switch stateCode {
		case 0:
			// Normal mode
			switch pendChar(rune(currentChar)) {
			case 1:
				// Comment start
				stateCode = 1
			case 2:
				// The repeat sign
				result = append(result, Token{
					clip:      "*",
					tokenType: TokenTypeOpRepeat,
				})
			case 3:
				// Number digits 0-9
				clipBuffer += string(currentChar)
				if len(result) == 0 || result[len(result)-1].tokenType == TokenTypeOpRepeat {
					// Ordinary decimal number
					stateCode = 2
				} else {
					// Hexdecimal digits 0-9
					stateCode = 3
				}
			case 4:
				// Hexdecimal digits A-F
				clipBuffer += string(currentChar)
				stateCode = 3
			case 5:
				result = append(result, Token{
					clip:      "(",
					tokenType: TokenLeftBracket,
				})
			case 6:
				result = append(result, Token{
					clip:      ")",
					tokenType: TokenRightBracket,
				})
			case 7, 8:
				break
			default:
				_ = fmt.Errorf("Error: Unexpected token %s. \nError: Inopinatum signum %s.\n", string(currentChar), string(currentChar))
				errors = true
			}
		case 1:
			// Comment parsing
			if currentChar == '}' {
				result = append(result, Token{
					clip:      clipBuffer,
					tokenType: TokenTypeComment,
				})
				// Empty the buffer
				clipBuffer = ""
				stateCode = 0
			} else {
				// Add the clip buffer
				clipBuffer += string(currentChar)
			}
		case 2:
			// Decimal digit
			if currentChar >= '0' && currentChar <= '9' {
				// Still a valid digit
				clipBuffer += string(currentChar)
			} else {
				result = append(result, Token{
					clip:      clipBuffer,
					tokenType: TokenTypeNumberLiteral,
				})
				clipBuffer = ""
				stateCode = 0
				i--
			}
		case 3:
			// Hexadecimal digit
			if pendChar(rune(currentChar)) == 4 || pendChar(rune(currentChar)) == 3 {
				// A valid (next) digit
				clipBuffer += string(currentChar)
				clipBuffer = strings.ToUpper(clipBuffer)
				result = append(result, Token{
					clip:      clipBuffer,
					tokenType: TokenTypeHexLiteral,
				})
				clipBuffer = ""
				stateCode = 0
			} else {
				_ = fmt.Errorf("Error: Hexadecimal must be grouped in two digits. \nError: Hexadecimal debet esse binos et binos.\n")
				stateCode = 0
				errors = true
			}
		}
	}
	return result, errors
}
