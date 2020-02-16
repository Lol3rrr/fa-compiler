package generator

import (
  "compiler/frontend/parser"
)

func generatePart(root *parser.PT_Node) (string) {
  // Checks for the different Operations that are usually performed and then
  // generates the Code accordingly
  if root.Type == "DECLARATION" {
    return generateDeclaration(root)
  }
  if root.Type == "ASSIGNMENT" {
    return generateAssignment(root)
  }
  if root.Type == "CALL" {
    return generateFunctionCall(root)
  }

  // If it wasnt a known Operation, defaults to returning an empty string
  return ""
}
