package parser

import (
  "compiler/frontend/tokenizer"
)


func parseValue(restLine []tokenizer.Token, rootElement *PT_Node) int {
  // Checks if the first Element ISNOT a Value or ID of a Variable and returns that it DIDNT work
  if restLine[0].Type != "INT" && restLine[0].Type != "ID" {
    return 0
  }
  valueRoot := rootElement.Children[len(rootElement.Children) - 1]
  // Checks if the next Token in the line is an Operation (like + or -)
  if len(restLine) > 1 && (restLine[1].Type == "ADDITION" || restLine[1].Type == "SUBTRACTION") {
    valueRoot.addChild("OPERATION", restLine[1].Content)

    opRoot := valueRoot.Children[len(valueRoot.Children) - 1]
    opRoot.addChild(restLine[0].Type, restLine[0].Content)

    return parseValue(restLine[2:], valueRoot)
  }

  valueRoot.addChild(restLine[0].Type, restLine[0].Content)

  return 1
}
