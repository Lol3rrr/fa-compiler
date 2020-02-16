package optimizer

import (
  "strconv"

  "compiler/frontend/parser"
)

func optimizeChild(root *parser.PT_Node) {
  for _, child := range root.Children {
    optimizeChild(child)
  }

  if root.Type == "OPERATION" {
    if root.Children[0].Type == "INT" && root.Children[1].Type == "INT" {
      childOneValue, _ := strconv.Atoi(root.Children[0].Content)
      childTwoValue, _ := strconv.Atoi(root.Children[1].Content)

      nValue := 0
      if root.Content == "+" {
        nValue = childOneValue + childTwoValue
      } else if root.Content == "-" {
        nValue = childOneValue - childTwoValue
      }

      root.Type = "INT"
      root.Content = strconv.Itoa(nValue)
      root.Children = make([]*parser.PT_Node, 0)

      return
    }
  }
}

func Optimize(root *parser.PT_Node) (*parser.PT_Node) {
  for _, child := range root.Children {
    optimizeChild(child)
  }

  return root
}
