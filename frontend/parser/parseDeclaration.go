package parser

import (
  "compiler/frontend/tokenizer"
)

func parseDeclaration(line []tokenizer.Token, rootElement *PT_Node) (int) {
  // Create new Root element for this scope
  nRoot := &PT_Node{
    Type: "DECLARATION",
    Parent: rootElement,
    Children: make([]*PT_Node, 0),
  }

  // Checks if the Line starts with a Datatype, throws error if it DOESNT
  if line[0].Type != "DATATYPE" {
    return 0
  }
  // Adds the Value of the Datatype as the Datatype for the new Variable
  nRoot.addChild("DATATYPE", line[0].Content)

  // Checks if the next Token in the Line is an ID, throws error if it ISNT
  if line[1].Type != "ID" {
    return 0
  }
  // Adds the Value of the ID as the ID for the new Variable
  nRoot.addChild("ID", line[1].Content)

  // Checks if the next Token is an Equals Symbol, throws error if it ISNT
  if line[2].Type != "EQUALS" {
    return 0
  }
  // Checks if the Line ends with a Semicolon, throws error if it DOESNT
  if line[len(line) - 1].Type != "SEMICOLON" {
    return 0
  }

  nRoot.addChild("VALUE", "")
  // Parse the Value for the new Variable
  if parseValue(line[3:], nRoot) != 1 {
    return 0
  }

  // Add the new Root to the given Root element
  rootElement.Children = append(rootElement.Children, nRoot)

  // Return 1 to indicate that everything worked
  return 1
}
