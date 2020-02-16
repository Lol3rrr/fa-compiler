package frontend

import (
  "fmt"

  "compiler/frontend/parser"
  "compiler/frontend/semantics"
  "compiler/frontend/tokenizer"
)

func GenerateSytnaxTree(data []byte) (*parser.PT_Node, bool) {
  tokens := tokenizer.Tokenize(data)

  fmt.Printf("Tokens: '%+v' \n", tokens)

  parseTree, parseWorked := parser.Parse(tokens)
  if !parseWorked {
    return nil, false
  }
  parseTree.Print()

  semanticWorked := semantics.CheckSemantics(parseTree)
  if !semanticWorked {
    return nil, false
  }

  return parseTree, true
}
