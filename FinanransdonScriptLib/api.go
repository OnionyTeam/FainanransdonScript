package FinanransdonScriptLib

import "fmt"

func RunScript(script string, verbose ...bool) []byte {
	isVerbose := false
	if len(verbose) > 1 {
		panic("Error: Only allow at most 2 arguments for FinanransdonScriptLib.RunScript()")
	}
	if len(verbose) == 1 {
		isVerbose = verbose[0]
	}
	tokenList, hasErr := TokenizeCode(script)
	if isVerbose {
		fmt.Printf("Album de nota generatae: %#v\n", tokenList)
	}
	if hasErr {
		return nil
	}
	ast, hasErr := CreateASTFromTokenList(tokenList)
	if isVerbose {
		fmt.Printf("Arbor de AST generatae: %#v\n", ast)
	}
	if hasErr {
		return nil
	}
	result := GenerateByteSliceFromAST(ast)
	if isVerbose {
		fmt.Printf("Binarii content generatae: %#v\n", result)
	}
	return result
}

func DecompileFromData(data []byte, verbose ...bool) string {
	isVerbose := false
	if len(verbose) > 1 {
		panic("Error: Only allow at most 2 arguments for FinanransdonScriptLib.RunScript()")
	}
	if len(verbose) == 1 {
		isVerbose = verbose[0]
	}
	if isVerbose {
		fmt.Printf("Satus decompiling notitia %#v....\n", data)
	}
	result := ""
	isZeroMode := false
	zeroCount := 0
	for index, b := range data {
		if !isZeroMode {
			if b != 0 {
				result += fmt.Sprintf("%02X ", b)
			} else {
				// Zero opitimiziation
				isZeroMode = true
				zeroCount = 1
			}
			if index%13 == 12 {
				result += "\n"
			}
		} else {
			if b == 0 {
				zeroCount++
			} else {
				// Exit zero opitimiziation
				if zeroCount <= 3 {
					for i := 0; i < zeroCount; i++ {
						result += "00 "
					}
				} else {
					result += fmt.Sprintf("*%d(00)", zeroCount)
				}
				result += fmt.Sprintf("%02X ", b)
				isZeroMode = false
			}
		}
	}
	return result
}
