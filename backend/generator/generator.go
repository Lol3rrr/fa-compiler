package generator

import (
  "compiler/frontend/parser"
)

func GenerateCode(root *parser.PT_Node) (string) {
  result := "#include <stdio.h> \nint main() { \n"

  for _, child := range root.Children {
    result += generatePart(child)
  }

  result += "return 0; \n}"

  return result
}
