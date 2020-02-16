package semantics

func doesVariableExist(id string) bool {
  // Checks if the given ID is found in the Map of Variables
  _, found := variables[id]

  return found
}
