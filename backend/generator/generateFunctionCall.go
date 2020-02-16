package generator

import (
  "compiler/frontend/parser"
)

func generateFunctionCall(root *parser.PT_Node) (string)  {
  result := ""

  if root.Children[0].Content == "print" {
    return "printf(\"%d \\n \", " + root.Children[1].Children[0].Content + ");\n"
  }

  return result
}
