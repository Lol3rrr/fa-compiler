package parser

// Parse Tree Node
type PT_Node struct {
  Type string
  Content string
  Parent *PT_Node
  Children []*PT_Node
}
