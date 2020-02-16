package parser

import (
  "compiler/frontend/tokenizer"
)

func parseAssignment(line []tokenizer.Token, rootElement *PT_Node) (int) {
  // Creates a new Temporary Root Element
  nRoot := &PT_Node{
    Type: "ASSIGNMENT",
    Parent: rootElement,
    Children: make([]*PT_Node, 0),
  }

  // Checks if the first Token in the Line is an ID for a Variable, throws error if it ISNT
  if line[0].Type != "ID" {
    return 0
  }
  // Adds the Value of the first Token to the new Root element as the ID of the new Variable
  nRoot.addChild("ID", line[0].Content)

  // Checks if the next Token is an Equals Symbol, throws error if it ISNT
  if line[1].Type != "EQUALS" {
    return 0
  }
  // Checks if the Line ends with a Semicolon, throws error if it DOESNT
  if line[len(line) - 1].Type != "SEMICOLON" {
    return 0
  }

  // Add a new Child to the new Root element, with the Type "VALUE", this will have the
  // actual Value of it as Child elements
  nRoot.addChild("VALUE", "")
  // Get the Value for the Value object, throw error if something went wrong
  if parseValue(line[2:], nRoot) != 1 {
    return 0
  }

  // Add the new Root element to the existing root element
  rootElement.Children = append(rootElement.Children, nRoot)

  // Return 1 so signal, that everything worked
  return 1
}
