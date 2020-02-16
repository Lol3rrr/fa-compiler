package generator

import (
  "compiler/frontend/parser"
)


func getValue(root *parser.PT_Node) (string)  {
  // Checks if the Value is just a Simple constant, returns that if it IS
  if root.Type == "INT" || root.Type == "ID" {
    return root.Content
  }

  // Checks if the current Element is an Operation and returns the Value of the
  // first Child Element + the Type of Operation + the Value of the second Child
  // Element
  if root.Type == "OPERATION" {
    return getValue(root.Children[0]) + " " + root.Content + " " + getValue(root.Children[1])
  }

  return getValue(root.Children[0])
}
