package backend

import (
  "fmt"

  "compiler/frontend/parser"

  "compiler/backend/optimizer"
  "compiler/backend/generator"
)

func OptimizeAndGenerateCode(root *parser.PT_Node) (string) {
  // Optimizes the given AST
  nRoot := optimizer.Optimize(root)

  // Outputs the new optimized AST
  fmt.Printf("\n\n")
  fmt.Printf("*** After Optimization: *** \n")
  nRoot.Print()
  fmt.Printf("\n\n")

  // Directly returns the Generated Code
  return generator.GenerateCode(nRoot)
}
