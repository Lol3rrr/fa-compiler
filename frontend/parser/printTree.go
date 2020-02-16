package parser

import (
  "fmt"
)

func (root *PT_Node) Print() {
  for _, child := range root.Children {
    child.printIndentation("")
  }
}

func (node *PT_Node) printIndentation(indent string) {
  fmt.Printf("%s->[%s]%s\n", indent, node.Type, node.Content)

  for _, child := range node.Children {
    child.printIndentation(indent + "--")
  }
}
