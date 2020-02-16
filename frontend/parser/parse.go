package parser

import (
  "fmt"

  "compiler/frontend/tokenizer"
)

func Parse(tokens []tokenizer.Token) (*PT_Node, bool) {
  // Create the Root element of the AST
  root := &PT_Node{}

  // hasEnd is used to know if there were any statements left "open"/not closed
  hasEnd := false
  lastEnd := 0
  // goes through all the tokens
  for i, tmp := range tokens {
    // sets this to false
    hasEnd = false

    // Checks if it has reached the end of a statement
    if tmp.Type == "SEMICOLON" {
      // Sets this back to true, because at that moment there are no statements left open
      hasEnd = true
      // Gets the current Line in question
      line := tokens[lastEnd:i + 1]
      // Runs it through every parsing function
      worked := parseDeclaration(line, root)
      worked += parseAssignment(line, root)
      worked += parseFunctionCall(line, root)

      // Check if any of them managed to parse the line, throws error if they DIDNT
      if worked < 1 {
        lineStr := ""
        for _, linePart := range line {
          lineStr += linePart.Content + " "
        }
        fmt.Printf("[Error] in line %d: '%s' %d \n", tmp.Line, lineStr)
        return nil, false
      }

      // Sets the last end, which acts as a starting point for the next line, to the next element
      lastEnd = i + 1;
    }
  }

  // Throws error if there were statements left open
  if !hasEnd {
    fmt.Printf("[Error] Missing semicolon at the end of line %d \n", tokens[len(tokens) - 1].Line)
    return root, false
  }

  // Return the root element + true to indicate that everything worked
  return root, true
}
