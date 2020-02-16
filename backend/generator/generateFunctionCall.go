package generator

import (
  "compiler/frontend/parser"
)

func generateFunctionCall(root *parser.PT_Node) (string)  {
  result := ""

  if root.Children[0].Content == "print" {
    value := root.Children[1].Children[0].Content
    if root.Children[1].Children[0].Type == "OPERATION" {
      value = getValue(root.Children[1].Children[0])
    }

    return "printf(\"%d \\n \", " + value + ");\n"
  }

  return result
}
