package FinanransdonScriptLib

func generate(children []Node) (result []byte) {
	for _, child := range children {
		if child.operation == OperationOutput {
			result = append(result, byte(child.data))
		}
		if child.operation == OperationLoop {
			for i := 0; i < child.data; i++ {
				result = append(result, generate(child.children)...)
			}
		}
	}
	return
}

func GenerateByteSliceFromAST(ast Node) []byte {
	return generate(ast.children)
}
