package backend

import (
  "fmt"

  "compiler/frontend/parser"

  "compiler/backend/optimizer"
  "compiler/backend/generator"
)

func OptimizeAndGenerateCode(root *parser.PT_Node) (string) {
  nRoot := optimizer.Optimize(root)

  fmt.Printf("\n\n")
  fmt.Printf("*** After Optimization: *** \n")
  nRoot.Print()
  fmt.Printf("\n\n")

  return generator.GenerateCode(nRoot)
}
