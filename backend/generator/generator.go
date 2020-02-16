package generator

import (
  "compiler/frontend/parser"
)

func GenerateCode(root *parser.PT_Node) (string) {
  // Uses this as the Default Start for every program
  result := "#include <stdio.h> \nint main() { \n"

  // Adds the Generated code of each part to it
  for _, child := range root.Children {
    result += generatePart(child)
  }

  // Adds this to make sure it follows the C-Standard
  result += "return 0; \n}"

  return result
}
