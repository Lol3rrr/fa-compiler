package parser

import (
  "fmt"

  "compiler/frontend/tokenizer"
)

func Parse(tokens []tokenizer.Token) (*PT_Node, bool) {
  root := &PT_Node{}

  hasEnd := false
  lastEnd := 0
  for i, tmp := range tokens {
    hasEnd = false

    if tmp.Type == "SEMICOLON" {
      hasEnd = true
      line := tokens[lastEnd:i + 1]
      worked := parseDeclaration(line, root)
      worked += parseAssignment(line, root)
      worked += parseFunctionCall(line, root)

      if worked < 1 {
        lineStr := ""
        for _, linePart := range line {
          lineStr += linePart.Content + " "
        }
        fmt.Printf("[Error] in line %d: '%s' %d \n", tmp.Line, lineStr)
        return nil, false
      }

      lastEnd = i + 1;
    }
  }

  if !hasEnd {
    fmt.Printf("[Error] Missing semicolon at the end of line %d \n", tokens[len(tokens) - 1].Line)
  }

  return root, true
}
