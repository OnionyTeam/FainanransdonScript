package FinanransdonScriptLib

type Node struct {
	children  []Node
	operation int
	data      int
	isRoot    bool
}

const (
	OperationNone   = 0
	OperationLoop   = 1
	OperationOutput = 2
)
