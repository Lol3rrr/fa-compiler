package parser

import (
  "compiler/frontend/tokenizer"
)

func parseFunctionCall(line []tokenizer.Token, rootElement *PT_Node) (int) {
  nRoot := &PT_Node{
    Type: "CALL",
    Parent: rootElement,
    Children: make([]*PT_Node, 0),
  }

  if line[0].Type != "ID" {
    return 0
  }
  nRoot.addChild("ID", line[0].Content)

  if line[1].Type != "LPAREN" {
    return 0
  }
  if len(line) <= 4 && line[2].Type != "RPAREN" {
    return 0
  }
  if len(line) <= 4 && line[3].Type != "SEMICOLON" {
    return 0
  }
  if len(line) <= 4 {
    rootElement.Children = append(rootElement.Children, nRoot)
    return 1
  }

  if line[2].Type != "ID" && line[2].Type != "INT" {
    return 0
  }
  nRoot.addChild("VALUE", "")
  valueRoot := nRoot.Children[1]
  valueRoot.addChild(line[2].Type, line[2].Content)

  if line[3].Type != "RPAREN" {
    return 0
  }
  if line[4].Type != "SEMICOLON" {
    return 0
  }

  rootElement.Children = append(rootElement.Children, nRoot)

  return 1
}
