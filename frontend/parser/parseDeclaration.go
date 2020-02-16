package parser

import (
  "compiler/frontend/tokenizer"
)

func parseDeclaration(line []tokenizer.Token, rootElement *PT_Node) (int) {
  nRoot := &PT_Node{
    Type: "DECLARATION",
    Parent: rootElement,
    Children: make([]*PT_Node, 0),
  }

  if line[0].Type != "DATATYPE" {
    return 0
  }
  nRoot.addChild("DATATYPE", line[0].Content)

  if line[1].Type != "ID" {
    return 0
  }
  nRoot.addChild("ID", line[1].Content)

  if line[2].Type != "EQUALS" {
    return 0
  }
  if line[len(line) - 1].Type != "SEMICOLON" {
    return 0
  }

  nRoot.addChild("VALUE", "")
  if parseValue(line[3:], nRoot) != 1 {
    return 0
  }

  rootElement.Children = append(rootElement.Children, nRoot)

  return 1
}
