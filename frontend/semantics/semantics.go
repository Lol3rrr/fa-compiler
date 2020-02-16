package semantics

import (
  "compiler/frontend/parser"
)

var variables map[string]string

func CheckSemantics(root *parser.PT_Node) bool {
  variables = make(map[string]string)

  // Declare global variables
  variables["print"] = "function"

  for _, child := range root.Children {
    if !check(child) {
      return false
    }
  }

  return true
}
