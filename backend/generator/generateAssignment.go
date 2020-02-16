package generator

import (
  "compiler/frontend/parser"
)

func generateAssignment(root *parser.PT_Node) (string)  {
  result := ""
  result += root.Children[0].Content + " = "

  result += getValue(root.Children[1]) + ";\n"

  return result
}
