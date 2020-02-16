package optimizer

import (
  "strconv"

  "compiler/frontend/parser"
)

func optimizeChild(root *parser.PT_Node) {
  // Calls the Optimization function for all its Child elements
  for _, child := range root.Children {
    optimizeChild(child)
  }

  // if the current Element is of Type OPERATION, try to perform some optimizations
  if root.Type == "OPERATION" {
    // if both of its Child elements are already know variables, add them together and replace the Operation
    // with the constant value of this operation
    if root.Children[0].Type == "INT" && root.Children[1].Type == "INT" {
      // Convert the String representation of the Values to actual ints
      childOneValue, _ := strconv.Atoi(root.Children[0].Content)
      childTwoValue, _ := strconv.Atoi(root.Children[1].Content)

      // Calculate the Value based on the Operation
      nValue := 0
      if root.Content == "+" {
        nValue = childOneValue + childTwoValue
      } else if root.Content == "-" {
        nValue = childOneValue - childTwoValue
      }

      // Set the Type of the Root elemt to INT, its value to the previously calculated
      // one and clear its Child elements
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
