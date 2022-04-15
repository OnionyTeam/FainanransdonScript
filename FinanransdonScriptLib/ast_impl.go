package FinanransdonScriptLib

import (
	"fmt"
	"strconv"
)

func findTheIndexOfTheMatchingRightBracketWithGiveIndexOfTheLeftBracket(tokenList []Token, indexOfTheLeftBracket int) (result int) {
	stack := []interface{}{}
	for i := indexOfTheLeftBracket; i < len(tokenList); i++ {
		if tokenList[i].tokenType == TokenRightBracket {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				result = i
				return
			}
		}
		if tokenList[i].tokenType == TokenLeftBracket {
			stack = append(stack, nil)
		}
	}
	return
}

func parseASTLists(tokenList []Token) (result []Node, hasError bool) {
	stateCode := 0
	bufferNode := Node{
		isRoot: false,
	}
	for k := 0; k < len(tokenList); k++ {
		v := tokenList[k]
		switch stateCode {
		case 0:
			switch v.tokenType {
			case TokenTypeHexLiteral:
				// Hex literal
				data, err := strconv.ParseInt(v.clip, 16, 32)
				if err != nil {
					_ = fmt.Errorf("Error: Golang internal error occurred: %v\n", err)
					return
				}
				result = append(result, Node{
					children:  nil,
					data:      int(data),
					operation: OperationOutput,
					isRoot:    false,
				})
			case TokenTypeOpRepeat:
				// Repeat operation
				bufferNode.operation = OperationLoop
				stateCode = 1
			}
		case 1:
			if v.tokenType != TokenTypeNumberLiteral {
				_ = fmt.Errorf("Error: Unexpected token after repeat operator \"*\". \nError: Inopinatus signum post repetere operator \"*\".\n")
			}
			data, err := strconv.ParseInt(v.clip, 10, 32)
			if err != nil {
				_ = fmt.Errorf("Error: Golang internal error occurred: %v\n", err)
				return
			}
			bufferNode.data = int(data)
			stateCode = 0
			endIndex := findTheIndexOfTheMatchingRightBracketWithGiveIndexOfTheLeftBracket(tokenList, k)
			bufferNode.children, hasError = parseASTLists(tokenList[k+1 : endIndex])
			k = endIndex
			result = append(result, bufferNode)
			bufferNode = Node{}
		}
	}
	return
}

func CreateASTFromTokenList(tokenList []Token) (result Node, hasError bool) {
	children, hasError := parseASTLists(tokenList)
	result = Node{
		children:  children,
		data:      -1,
		operation: OperationNone,
		isRoot:    true,
	}
	return
}
