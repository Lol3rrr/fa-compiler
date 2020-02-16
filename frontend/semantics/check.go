package semantics

import (
  "fmt"

  "compiler/frontend/parser"
)

func check(tmpRoot *parser.PT_Node) bool {
  if tmpRoot.Type == "DECLARATION" {
    dataType := tmpRoot.Children[0].Content
    id := tmpRoot.Children[1].Content

    if doesVariableExist(id) {
      fmt.Printf("[Error] %s is already defined \n", id)
      return false
    }

    variables[id] = dataType
  } else if tmpRoot.Type == "ASSIGNMENT" {
    id := tmpRoot.Children[0].Content

    if !doesVariableExist(id) {
      fmt.Printf("[Error] %s is not defined \n", id)
      return false
    }
  } else if tmpRoot.Type == "ID" {
    id := tmpRoot.Content

    if !doesVariableExist(id) {
      fmt.Printf("[Error] %s is not defined \n", id)
      return false
    }
  }

  for _, child := range tmpRoot.Children {
    if !check(child) {
      return false
    }
  }

  return true
}
