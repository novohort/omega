package evaluator

import (
	"interpreter/ast"
	"interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
