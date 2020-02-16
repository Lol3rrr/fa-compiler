package generator

import (
  "compiler/frontend/parser"
)


func getValue(root *parser.PT_Node) (string)  {
  if root.Type == "INT" || root.Type == "ID" {
    return root.Content
  }

  if root.Type == "OPERATION" {
    return getValue(root.Children[0]) + " " + root.Content + " " + getValue(root.Children[1])
  }

  return getValue(root.Children[0])
}
