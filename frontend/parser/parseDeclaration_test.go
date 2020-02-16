package parser

import (
  "testing"

  "compiler/frontend/tokenizer"
)

func TestParseDeclaration(t *testing.T) {
  tables := []struct{
    InputLine []tokenizer.Token
    InputRoot *PT_Node
    Result int
  }{
    {
      InputLine: []tokenizer.Token{
        {
          Type: "DATATYPE",
          Content: "int",
        },
        {
          Type: "ID",
          Content: "test",
        },
        {
          Type: "EQUALS",
          Content: "=",
        },
        {
          Type: "INT",
          Content: "6",
        },
        {
          Type: "SEMICOLON",
          Content: ";",
        },
      },
      InputRoot: &PT_Node{},
      Result: 1,
    },
  }

  for _, table := range tables {
    inLine := table.InputLine
    inRoot := table.InputRoot
    result := table.Result

    output := parseDeclaration(inLine, inRoot)

    if output != result {
      t.Errorf("Failed to parse it correctly")
      continue
    }
  }
}
