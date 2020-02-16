package parser

import (
  "fmt"
)

func (root *PT_Node) Print() {
  // Used to start the whole process with zero indentation
  for _, child := range root.Children {
    child.printIndentation("")
  }
}

func (node *PT_Node) printIndentation(indent string) {
  // Prints the indentation + the Type of the Node + the Content of the Node
  fmt.Printf("%s->[%s]%s\n", indent, node.Type, node.Content)

  // Prints every Child element, but with more indentation
  for _, child := range node.Children {
    child.printIndentation(indent + "--")
  }
}
