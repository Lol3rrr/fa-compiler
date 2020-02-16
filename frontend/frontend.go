package frontend

import (
  "fmt"

  "compiler/frontend/parser"
  "compiler/frontend/semantics"
  "compiler/frontend/tokenizer"
)

func GenerateSytnaxTree(data []byte) (*parser.PT_Node, bool) {
  // Tokenizes the input Code
  tokens := tokenizer.Tokenize(data)

  fmt.Printf("Tokens: '%+v' \n", tokens)

  // Uses that slice of Tokens to create a simple AST, this also checks Syntax while parsing
  parseTree, parseWorked := parser.Parse(tokens)
  if !parseWorked {
    return nil, false
  }

  // Outputs the whole tree
  parseTree.Print()

  // Takes the AST and checks if all the Semantics are correct (mostly if variables were declared before being used)
  semanticWorked := semantics.CheckSemantics(parseTree)
  if !semanticWorked {
    return nil, false
  }

  // Returns the AST and true to indicate that everything worked
  return parseTree, true
}
