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
