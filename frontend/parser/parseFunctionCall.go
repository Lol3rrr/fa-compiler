package parser

import (
  "compiler/frontend/tokenizer"
)

func parseFunctionCall(line []tokenizer.Token, rootElement *PT_Node) (int) {
  // Create a new Root element, only acts as Root in this scope
  nRoot := &PT_Node{
    Type: "CALL",
    Parent: rootElement,
    Children: make([]*PT_Node, 0),
  }

  // Check if the Line starts with an ID, throw error if it DOESNT
  if line[0].Type != "ID" {
    return 0
  }
  // Use the content of the first Token as the ID of the function
  nRoot.addChild("ID", line[0].Content)

  // Check if the next Token is an opening normal Bracket, throw error if it ISNT
  if line[1].Type != "LPAREN" {
    return 0
  }
  // Check if the Line has a closing normal Bracket before the Semicolon, throw error if it DOESNT
  if line[len(line) - 2].Type != "RPAREN" {
    return 0
  }
  // Check if the Line has a Semicolon at the end, throw error if it DOESNT
  if line[len(line) - 1].Type != "SEMICOLON" {
    return 0
  }
  // Already return if the Function doesnt take any parameter
  if len(line) <= 4 {
    // Adds the new Root element to the given Root Element
    rootElement.Children = append(rootElement.Children, nRoot)
    // Returns 1 to indicate that everything worked
    return 1
  }

  nRoot.addChild("VALUE", "")
  // Parses the value given as the Parameter for the Function, throws error if something DIDNT work
  if parseValue(line[2:len(line) - 2], nRoot) != 1 {
    return 0
  }

  // Add the new Root element to the given Root element
  rootElement.Children = append(rootElement.Children, nRoot)

  // Return 1 to indicate that everything worked
  return 1
}
