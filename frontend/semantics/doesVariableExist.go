package semantics

func doesVariableExist(id string) bool {
  _, found := variables[id]

  return found
}
