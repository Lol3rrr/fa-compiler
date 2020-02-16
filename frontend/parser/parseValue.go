package parser

import (
  "compiler/frontend/tokenizer"
)


func parseValue(restLine []tokenizer.Token, rootElement *PT_Node) int {
  if restLine[0].Type != "INT" && restLine[0].Type != "ID" {
    return 0
  }
  valueRoot := rootElement.Children[len(rootElement.Children) - 1]
  if restLine[1].Type == "ADDITION" || restLine[1].Type == "SUBTRACTION" {
    valueRoot.addChild("OPERATION", restLine[1].Content)

    opRoot := valueRoot.Children[len(valueRoot.Children) - 1]
    opRoot.addChild(restLine[0].Type, restLine[0].Content)

    return parseValue(restLine[2:], valueRoot)
  }

  valueRoot.addChild(restLine[0].Type, restLine[0].Content)

  return 1
}
