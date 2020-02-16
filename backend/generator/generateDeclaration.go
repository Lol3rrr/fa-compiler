package generator

import (
  "compiler/frontend/parser"
)

func generateDeclaration(root *parser.PT_Node) (string) {
  result := ""
  result += root.Children[0].Content + " "
  result += root.Children[1].Content + " = "

  result += getValue(root.Children[2]) + ";\n"

  return result
}
