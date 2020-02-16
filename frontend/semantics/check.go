package semantics

import (
  "fmt"

  "compiler/frontend/parser"
)

func check(tmpRoot *parser.PT_Node) bool {
  if tmpRoot.Type == "DECLARATION" {
    // Checks if the "Parameters" of this node match the ones expected for a declaration
    dataType := tmpRoot.Children[0].Content
    id := tmpRoot.Children[1].Content

    // Checks if the Variable already exists, throws an error if it DOES
    if doesVariableExist(id) {
      fmt.Printf("[Error] %s is already defined \n", id)
      return false
    }

    variables[id] = dataType
  } else if tmpRoot.Type == "ASSIGNMENT" {
    id := tmpRoot.Children[0].Content

    // Checks if the Variable you want to assign exists, throws error if it DOESNT
    if !doesVariableExist(id) {
      fmt.Printf("[Error] %s is not defined \n", id)
      return false
    }
  } else if tmpRoot.Type == "ID" {
    id := tmpRoot.Content

    // Checks if the Variable you want to use exists, throws error if it DOESNT
    if !doesVariableExist(id) {
      fmt.Printf("[Error] %s is not defined \n", id)
      return false
    }
  }

  // Also checks the Semantics for all its Child Elements
  for _, child := range tmpRoot.Children {
    if !check(child) {
      return false
    }
  }

  // Returns true to indicate that everything worked
  return true
}
