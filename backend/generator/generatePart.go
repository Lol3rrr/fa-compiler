package generator

import (
  "compiler/frontend/parser"
)

func generatePart(root *parser.PT_Node) (string) {
  if root.Type == "DECLARATION" {
    return generateDeclaration(root)
  }
  if root.Type == "ASSIGNMENT" {
    return generateAssignment(root)
  }
  if root.Type == "CALL" {
    return generateFunctionCall(root)
  }

  return ""
}
